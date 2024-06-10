package module

import(
	"time"
	"log"
)


type User struct{
	ID			int
	UUID		string
	Name		string
	Email		string
	Password	string
	CreateAt	time.Time
}

func (u *User) CreateUser()(err error){
	cmdCreateUser := `insert into users(
		uuid,
		name,
		email,
		password,
		create_at) values ($1,$2,$3,$4,$5)`

	_, err = Db.Exec(cmdCreateUser,
		CreateUUID(),
		u.Name,
		u.Email,
		Encrypt(u.Password),
		time.Now())
	
	if err != nil{
		log.Fatalln(err)
	}
	return err
}