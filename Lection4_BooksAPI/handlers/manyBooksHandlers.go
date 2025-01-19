package handlers

import (
	"BooksAPI/models"
	"encoding/json"
	"log"
	"net/http"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllBooks(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Get info about all books in database")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(models.DB)
}
