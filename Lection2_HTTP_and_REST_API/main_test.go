package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetGreet(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.01:8080/", nil)
	if err != nil {
		log.Fatal(err)
	}

	res := httptest.NewRecorder()
	GetGreet(res, req)

	exp := "<h1>Hello World!</h1>"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s, go %s", exp, act)
	}

	resCode := res.Result().StatusCode
	if resCode != 200 {
		t.Fatalf("Expected 200, go %d", resCode)
	}
}
