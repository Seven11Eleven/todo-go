package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func ConnectDB() {
	var err error
	Conn, err = pgx.Connect(context.Background(), "user=postgres dbname=todolist password=todo123 port=1337 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

}

// func createTableIfNotExist(){
	
// 	createToDoList := `
// 		CREATE TABLE IF NOT EXISTS todo_lists (
// 			id SERIAL PRIMARY KEY,
// 			name TEXT
// 		);
// 	`
// 	createTasks :=`
// 	CREATE TABLE IF NOT EXISTS tasks(
// 		id SERIAL PRIMARY KEY,
// 		todo_id INT REFERENCES todo_lists(id),
// 		name TEXT,
// 		completed BOOL
// 	);
// 	`

// 	_, err = Сonn.Exec(context.Background(), createToDoList)
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	_, err = Сonn.Exec(context.Background(), createTasks)
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// }
