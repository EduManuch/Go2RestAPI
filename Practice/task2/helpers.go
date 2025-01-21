package main

import (
	"log"
	"net/http"
)

func CheckError(err error, status int, w http.ResponseWriter) {
	if err != nil {
		if status > 0 {
			w.WriteHeader(status)
			log.Println(err)
		} else {
			log.Fatal(err)
		}
	}
}
