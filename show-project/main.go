package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	api := "https://jsonplaceholder.typicode.com/todos"

	response, _ := http.Get(api)

	responseData, _ := ioutil.ReadAll(response.Body)

	var todos []Todo

	json.Unmarshal(responseData, &todos)

	csvFile, _ := os.Create("./test.csv")
	writer := csv.NewWriter(csvFile)

	for _, todo := range todos {
		var row []string
		row = append(row, strconv.Itoa(todo.UserID))
		row = append(row, strconv.Itoa(todo.ID))
		row = append(row, todo.Title)

		writer.Write(row)
	}

	writer.Flush()
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
