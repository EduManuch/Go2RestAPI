package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()
	fmt.Println("File descriptor successfully created!")

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)
	fmt.Println(result["users"])
}
