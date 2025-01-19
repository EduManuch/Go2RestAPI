package handlers

import (
	"BooksAPI/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error occurs while parsing id field:", err)
		writer.WriteHeader(http.StatusBadRequest)
		message := models.Message{Message: "don't use parametr ID as uncasted to int."}
		json.NewEncoder(writer).Encode(message)
		return
	}
	book, ok := models.FindBookByID(id)
	log.Println("get book with id:", id)
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		message := models.Message{Message: "book with that ID does not exist in database."}
		json.NewEncoder(writer).Encode(message)
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(book)
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Creating new book ...")
	var book models.Book

	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&book)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		message := models.Message{Message: "provided json is invalid"}
		json.NewEncoder(writer).Encode(message)
		return
	}

	var NewBookId int = len(models.DB) + 1
	book.ID = NewBookId
	models.DB = append(models.DB, book)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(book)
}

func UpdateBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Updating book ...")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error occurs while parsing id field:", err)
		writer.WriteHeader(http.StatusBadRequest)
		message := models.Message{Message: "don't use ID parametr as uncasted to int."}
		json.NewEncoder(writer).Encode(message)
	}

	_, ok := models.FindBookByID(id)
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		message := models.Message{Message: "don't use ID parametr as uncasted to int."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	var newBook models.Book

	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&newBook)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		message := models.Message{Message: "provided json file is invalid."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	res := models.UpdateBookById(id, newBook)
	if !res {
		writer.WriteHeader(http.StatusNoContent)
		return
	}

	writer.WriteHeader(http.StatusOK)
	message := models.Message{Message: "book has successfully changed."}
	json.NewEncoder(writer).Encode(message)
}

func DeleteBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error occurs while parsing id field:", err)
		writer.WriteHeader(http.StatusBadRequest)
		message := models.Message{Message: "don't use ID parametr as uncasted to int."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	_, ok := models.FindBookByID(id)
	if !ok {
		log.Println("book with id =", id, "not found.")
		writer.WriteHeader(http.StatusNotFound)
		message := models.Message{Message: "book with that ID does not exist in database."}
		json.NewEncoder(writer).Encode(message)
		return
	}

	models.DeleteBookById(id)
	message := models.Message{Message: "book has successfully deleted from database."}
	json.NewEncoder(writer).Encode(message)
}
