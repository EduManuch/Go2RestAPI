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
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"

	"github.com/gorilla/mux"
)

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

func GetFirst(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>first num is: %d</h1>", rand.IntN(100))

}

func GetSecond(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>second num is: %d</h1>", rand.IntN(100))

}

func GetAdd(w http.ResponseWriter, r *http.Request) {
	res := rand.IntN(100) + rand.IntN(100)
	fmt.Fprintf(w, "<h1>Add result: %d</h1>", res)

}

func GetSub(w http.ResponseWriter, r *http.Request) {
	res := rand.IntN(100) - rand.IntN(100)
	fmt.Fprintf(w, "<h1>Sub result: %d</h1>", res)

}

func GetMul(w http.ResponseWriter, r *http.Request) {
	res := rand.IntN(100) * rand.IntN(100)
	fmt.Fprintf(w, "<h1>Mul result: %d</h1>", res)

}

func GetDiv(w http.ResponseWriter, r *http.Request) {
	res := rand.IntN(100) / rand.IntN(100)
	fmt.Fprintf(w, "<h1>Div result: %d</h1>", res)

}

func main() {
	fmt.Println("Starting REST API...")
	router := mux.NewRouter()
	router.HandleFunc("/info", GetInfo)
	router.HandleFunc("/first", GetFirst)
	router.HandleFunc("/second", GetSecond)
	router.HandleFunc("/add", GetAdd)
	router.HandleFunc("/sub", GetSub)
	router.HandleFunc("/mul", GetMul)
	router.HandleFunc("/div", GetDiv)
	log.Println("Router configured successfully! Let's go!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
