package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

type todo struct {
	id      int
	task    string
	owner   string
	checked int
}

func main() {
	os.Remove("./todo.db")
	db, err := sql.Open("sqlite", "./todo.db")
	if err != nil {
		log.Fatal((err))
	}
	defer db.Close()

	{
		sqlStmt := `
		CREATE TABLE IF NOT EXISTS task (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT,
		owner TEXT,
		checked INTEGER);
		`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			log.Printf("%q: %s\n", err, sqlStmt)
		}
	}

	{
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}
		stmt, err := tx.Prepare("INSERT INTO task(id, task, owner, checked) VALUES(?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		tasks := []*todo{}
		tasks = append(tasks, &todo{id: 1, task: "Learn REST API", owner: "teacher", checked: 0})
		tasks = append(tasks, &todo{id: 2, task: "Make practice", owner: "students", checked: 0})

		for i := range tasks {
			_, err := stmt.Exec(tasks[i].id, tasks[i].task, tasks[i].owner, tasks[i].checked)
			if err != nil {
				log.Fatal(err)
			}
		}
		if err := tx.Commit(); err != nil {
			log.Fatal(err)
		}
	}

	{
		stmt, err := db.Prepare("SELECT id, task, owner FROM task WHERE checked = ?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		rows, err := stmt.Query(0)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var task string
			var owner string

			err = rows.Scan(&id, &task, &owner)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, task, owner)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}
