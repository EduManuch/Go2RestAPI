/*
## Задача № 1 (Calculator API)
Написать API для указанных маршрутов(endpoints)
"/info"   // Информация об API
"/first"  // Случайное число
"/second" // Случайное число
"/add"    // Сумма двух случайных чисел
"/sub"    // Разность
"/mul"    // Произведение
"/div"    // Деление

*результат вернуть в виде JSON
import "math/rand"
number := rand.Intn(100)

// Queries
GET http://127.0.0.1:1234/info
GET http://127.0.0.1:1234/first
GET http://127.0.0.1:1234/second
GET http://127.0.0.1:1234/add
GET http://127.0.0.1:1234/sub
GET http://127.0.0.1:1234/mul
GET http://127.0.0.1:1234/div
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"

	"github.com/gorilla/mux"
)

type Nums struct {
	First      int    `json:"first"`
	Second     int    `json:"second"`
	Result     int    `json:"result"`
	ErrMessage string `json:"error"`
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	txt := `
	<h1>Calculator API</h1>
	<h2>Methods:</h2>
	<ul>
		<li>/first</li>
		<li>/second</li>
		<li>/add</li>
		<li>/sub</li>
		<li>/mul</li>
		<li>/div</li>
	</ul>`
	fmt.Fprint(w, txt)
}

func (n *Nums) GetFirst(w http.ResponseWriter, r *http.Request) {
	n.ErrMessage = ""
	n.First = rand.IntN(100)
	js, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "<h1>FIRST: %d</h1><h2>%s</h2>", n.First, js)

}

func (n *Nums) GetSecond(w http.ResponseWriter, r *http.Request) {
	n.ErrMessage = ""
	n.Second = rand.IntN(100)
	js, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "<h1>SECOND: %d</h1><h2>%s</h2>", n.Second, js)

}

func (n *Nums) GetAdd(w http.ResponseWriter, r *http.Request) {
	n.ErrMessage = ""
	n.Result = n.First + n.Second
	js, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "<h1>ADD RESULT: %d</h1><h2>%s</h2>", n.Result, js)
}

func (n *Nums) GetSub(w http.ResponseWriter, r *http.Request) {
	n.ErrMessage = ""
	n.Result = n.First - n.Second
	js, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "<h1>SUB RESULT: %d</h1><h2>%s</h2>", n.Result, js)
}

func (n *Nums) GetMul(w http.ResponseWriter, r *http.Request) {
	n.ErrMessage = ""
	n.Result = n.First * n.Second
	js, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "<h1>MUL RESULT: %d</h1><h2>%s</h2>", n.Result, js)
}

func (n *Nums) GetDiv(w http.ResponseWriter, r *http.Request) {
	var message string
	n.ErrMessage = ""
	if n.Second == 0 {
		n.ErrMessage = "Деление на ноль запрещено!"
		message = fmt.Sprintf("<h1>DIV RESULT: %s</h1>", n.ErrMessage)
	} else {
		n.Result = n.First / n.Second
		message = fmt.Sprintf("<h1>DIV RESULT: %d</h1>", n.Result)
	}
	js, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, message+"<h2>%s</h2>", js)
}

func main() {
	fmt.Println("Starting REST API...")
	nums := Nums{}
	router := mux.NewRouter()
	router.HandleFunc("/info", GetInfo)
	router.HandleFunc("/first", nums.GetFirst)
	router.HandleFunc("/second", nums.GetSecond)
	router.HandleFunc("/add", nums.GetAdd)
	router.HandleFunc("/sub", nums.GetSub)
	router.HandleFunc("/mul", nums.GetMul)
	router.HandleFunc("/div", nums.GetDiv)
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
