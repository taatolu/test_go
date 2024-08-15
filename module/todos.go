package module

import(
	"log"
	"time"
)

type Todo struct{
	ID int
	Title	string
	Userid	int
	Create_at	time.Time
}


func (u *User) CreateTodo(title string)(err error){
	cmd:= `INSERT INTO todos(
	title,
	userid,
	create_at) values ($1,$2,$3)`

	_, err = Db.Exec(cmd, title, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTodo (todoid int)(todo Todo, err error){
	todo = Todo{}

	cmd := `SELECT * FROM todos WHERE id=$1`
	err = Db.QueryRow(cmd, todoid).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Userid,
		&todo.Create_at)
	
	return todo, err
}

func (u *User)GetTodosByUserID()(todos []Todo, err error){
	cmd := `SELECT * FROM todos WHERE userid=$1`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil{
		log.Fatalln(err)
	}

	for rows.Next() {
		todo := Todo{}
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Userid,
			&todo.Create_at)
		if err != nil {
			log.Fatalln(err)
		}

		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func DeleteTodo(todoid int)(err error){
	cmd := `DELETE FROM todos WHERE id = $1`
	_, err = Db.Exec(cmd, todoid)
	if err != nil{
		log.Fatalln(err)
	}
	return err
}