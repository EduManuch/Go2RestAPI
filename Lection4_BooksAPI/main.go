package main

import (
	"BooksAPI/utils"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const (
	apiPrfix string = "/api/v1"
)

var (
	port                    string
	bookResourcePrefix      string = apiPrfix + "/book"
	manyBooksResourcePrefix string = apiPrfix + "/books"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Coluld not found .env file", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManyBooksResource(router, manyBooksResourcePrefix)

	log.Println("Router initializing successfully. Ready to go.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
