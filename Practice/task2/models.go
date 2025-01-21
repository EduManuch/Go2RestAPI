package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

type Task struct {
	ID   int       `json:"id"`
	Text string    `json:"text"`
	Tags []string  `json:"tags"`
	Due  time.Time `json:"due"`
}

type TaskSqlReceiver struct {
	db *sql.DB
}

func InitDB() *TaskSqlReceiver {
	db, err := sql.Open("sqlite", "./tasks.db")
	CheckError(err, 0, nil)
	CreateAndFillTable(db)
	ts := TaskSqlReceiver{db: db}
	log.Println("Task SQL Receiver initialized")
	return &ts
}

func CreateAndFillTable(db *sql.DB) {
	stmt := `CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	text TEXT,
	tags TEXT,
	Due INTEGER);`

	_, err := db.Exec(stmt)
	CheckError(err, 0, nil)
	for i := 1; i < 4; i++ {
		task := fmt.Sprintf("Task#%d", i)
		tags := fmt.Sprintf("tag%d,tag1%d", i, i)
		_, err = db.Exec("INSERT INTO tasks(text, tags, due) VALUES(?, ?, ?)", task, tags, time.Now().AddDate(0, 0, i).Unix())
		CheckError(err, 0, nil)
	}
	log.Println("Database created")
}

func (ts *TaskSqlReceiver) CreateTaskDB(task *Task) (int, error) {
	var tags string
	for _, tg := range task.Tags {
		tags += tg + ","
	}
	tags, _ = strings.CutSuffix(tags, ",")

	result, err := ts.db.Exec("INSERT INTO tasks(text,tags,due) VALUES(?, ?, ?)", task.Text, tags, task.Due.Unix())

	newID, _ := result.LastInsertId()
	return int(newID), err
}

func (ts *TaskSqlReceiver) GetTaskByIdDB(id int) (*Task, error) {
	t := &Task{}
	var tags string
	var due int64

	row := ts.db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	err := row.Scan(&t.ID, &t.Text, &tags, &due)
	if err == nil {
		t.Tags = strings.Split(tags, ",")
		t.Due = time.Unix(due, 0)
	}
	return t, err
}

func (ts *TaskSqlReceiver) GetAllTasksDB() ([]Task, error) {
	rows, err := ts.db.Query("SELECT * FROM tasks;")
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		t := Task{}
		var tags string
		var due int64

		rows.Scan(&t.ID, &t.Text, &tags, &due)

		t.Tags = strings.Split(tags, ",")
		t.Due = time.Unix(due, 0)
		tasks = append(tasks, t)
	}

	return tasks, err
}

func (ts *TaskSqlReceiver) DeleteTaskByIdDB(id int) error {
	tx, err := ts.db.Begin()

	t := Task{}
	var tags, due string

	if err == nil {
		row := tx.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
		err := row.Scan(&t.ID, &t.Text, &tags, &due)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.Exec("DELETE FROM tasks WHERE id = ?", id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return err
}

func (ts *TaskSqlReceiver) DeleteAllTasksDB() error {
	_, err := ts.db.Exec("DELETE FROM tasks;")
	return err
}

func (ts *TaskSqlReceiver) GetTaskByTagDB(tag string) ([]Task, error) {
	stmt := fmt.Sprintf("SELECT * FROM tasks WHERE tags LIKE '%%%s%%';", tag)
	rows, err := ts.db.Query(stmt)
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		t := Task{}
		var tags string
		var due int64

		rows.Scan(&t.ID, &t.Text, &tags, &due)

		t.Tags = strings.Split(tags, ",")
		t.Due = time.Unix(due, 0)
		tasks = append(tasks, t)
	}

	return tasks, err
}

func (ts *TaskSqlReceiver) GetTaskByDateDB(y, m, d int) ([]Task, error) {
	uDate := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local).Unix()

	rows, err := ts.db.Query("SELECT * FROM tasks WHERE (due - ?) BETWEEN 0 AND 86400;", uDate)
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		t := Task{}
		var tags string
		var due int64

		rows.Scan(&t.ID, &t.Text, &tags, &due)

		t.Tags = strings.Split(tags, ",")
		t.Due = time.Unix(due, 0)
		tasks = append(tasks, t)
	}

	return tasks, err
}
