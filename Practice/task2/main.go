package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	log.Println("Starting API server...")
	os.Remove("./tasks.db")
	ts := InitDB()
	defer ts.db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/tasks/", ts.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", ts.GetTaskByID).Methods("GET")
	router.HandleFunc("/tasks/", ts.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", ts.DeleteTaskByID).Methods("DELETE")
	router.HandleFunc("/tasks/", ts.DeleteAllTasks).Methods("DELETE")
	router.HandleFunc("/tags/{tagname}", ts.GetTaskByTag).Methods("GET")
	router.HandleFunc("/due/{yy}/{mm}/{dd}", ts.GetTaskByDate).Methods("GET")

	log.Printf("Server successfully started on port%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
