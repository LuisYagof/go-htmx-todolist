package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id   int
	Text string
}

func main() {
	data := map[string][]Todo{
		"TodoList": {
			Todo{Id: 1, Text: "HTTP Server with Golang & HTMX"},
		},
	}

	fmt.Println(data)

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))

		templ.Execute(w, data)
	}

	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		newItem := r.PostFormValue("new-item")
		templ := template.Must(template.ParseFiles(("index.html")))
		todo := Todo{Id: len(data["TodoList"]) + 1, Text: newItem}
		data["TodoList"] = append(data["TodoList"], todo)

		fmt.Println(data)

		templ.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo", addTodoHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))

}
