package main

import (
	"fmt"
	"net/http"
)

var tasks = []string{"task1", "task2", "task3"}

func main() {
	http.HandleFunc("/", rootPath)
	http.HandleFunc("/home", homePath)
	// first server will be started with port 8090, second does not start
	http.HandleFunc("/tasks", taskList)
	http.ListenAndServe(":8090", nil)
	http.ListenAndServe(":8080", nil) // starting server in port 8080
}

func taskList(writer http.ResponseWriter, request *http.Request) {
	for _, task := range tasks {
		fmt.Fprintln(writer, task)
	}
}

func homePath(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Home path"))
}

func rootPath(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Root path")
}
