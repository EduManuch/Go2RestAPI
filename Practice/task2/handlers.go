package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ts *TaskSqlReceiver) CreateTask(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var task Task
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&task)
	CheckError(err, http.StatusBadRequest, writer)

	if err == nil {
		resp := make(map[string]int)
		resp["id"], err = ts.CreateTaskDB(&task)
		CheckError(err, http.StatusInternalServerError, writer)
		if err == nil {
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(resp)
		}
	}
}

func (ts *TaskSqlReceiver) GetTaskByID(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	CheckError(err, http.StatusBadRequest, writer)

	task, err := ts.GetTaskByIdDB(id)
	CheckError(err, http.StatusNotFound, writer)
	if err == nil {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(task)
	}
}

func (ts *TaskSqlReceiver) GetAllTasks(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	tasks, err := ts.GetAllTasksDB()
	CheckError(err, http.StatusNotFound, writer)

	if err == nil {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(tasks)
	}
}

func (ts *TaskSqlReceiver) DeleteTaskByID(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	CheckError(err, http.StatusBadRequest, writer)

	if err == nil {
		err = ts.DeleteTaskByIdDB(id)
		CheckError(err, http.StatusNotFound, writer)
		if err == nil {
			m := make(map[string]string)
			m["id"] = strconv.Itoa(id)
			m["deleted"] = "ok"
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(m)
		}
	}
}

func (ts *TaskSqlReceiver) DeleteAllTasks(writer http.ResponseWriter, request *http.Request) {
	err := ts.DeleteAllTasksDB()
	CheckError(err, http.StatusInternalServerError, writer)
	if err == nil {
		writer.WriteHeader(http.StatusOK)
	}
}

func (ts *TaskSqlReceiver) GetTaskByTag(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	tag := mux.Vars(request)["tagname"]
	tasks, err := ts.GetTaskByTagDB(tag)
	CheckError(err, http.StatusBadRequest, writer)

	if err == nil {
		if len(tasks) == 0 {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(tasks)
		}
	}
}

func (ts *TaskSqlReceiver) GetTaskByDate(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	y, errY := strconv.Atoi(mux.Vars(request)["yy"])
	m, errM := strconv.Atoi(mux.Vars(request)["mm"])
	d, errD := strconv.Atoi(mux.Vars(request)["dd"])
	if errY != nil || errM != nil || errD != nil {
		CheckError(errors.New("incorrect date"), http.StatusBadRequest, writer)
	}

	y += 2000
	tasks, err := ts.GetTaskByDateDB(y, m, d)
	CheckError(err, http.StatusInternalServerError, writer)

	if err == nil {
		if len(tasks) == 0 {
			writer.WriteHeader(http.StatusNotFound)
		} else {
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(tasks)
		}
	}
}
