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

func GetUser (userId int) (user User, err error) {
	user = User{}
	cmd:= `SELECT * FROM users WHERE ID = $1`
	
	err = Db.QueryRow(cmd,userId).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreateAt,
	)

	return user, err
}


func (u *User)UpdateUserPassword()(err error){
	cmd := `UPDATE users SET password=$2 WHERE ID=$1`
	newpss := Encrypt(u.Password)
	_, err = Db.Exec(cmd, u.ID, newpss)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}


func DeleteUser(userid int)(err error){
	cmd := `DELETE FROM users WHERE id=$1`
	_, err = Db.Exec(cmd,userid)
	if err != nil{
		log.Fatalln(err)
	}
	return err
}