package module

import(
	"fmt"
	"log"
	"database/sql"

    _ "github.com/lib/pq"
    "github.com/google/uuid"
	"crypto/sha1"

	"main/config"
)

var Db *sql.DB
var err error

const(
	tableNameUser = "users"
)

func init() {
	host := config.DbConfig.Host
    port := config.DbConfig.Port
    user := config.DbConfig.UserName
    password := config.DbConfig.Password
    dbname := config.DbConfig.DbName

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	Db, err = sql.Open("postgres",connStr)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
        id			SERIAL,
		uuid		VARCHAR,
        name		VARCHAR,
		email		VARCHAR,
		password	VARCHAR,
		create_at	TIMESTAMP
		)`,tableNameUser)

	_, err = Db.Exec(cmdU)
	if err != nil{
		log.Fatalln(err)
	}

}

func CreateUUID()(uuidobj uuid.UUID){
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string)(cryptext string){
	cryptext = fmt.Sprintf("%x",sha1.Sum([]byte(plaintext)))
	return cryptext
}