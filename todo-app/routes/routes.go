package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"todo-app/model"

	"github.com/gorilla/mux"
)

func sendTodos(w http.ResponseWriter) {
	todos, err := model.GetAllTodos()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("static/index.html"))

	err = tmpl.ExecuteTemplate(w, "Todos", todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
	}
}

func loadAlltodos(w http.ResponseWriter, r *http.Request) {
	todos, err := model.GetAllTodos()
	if err != nil {
		log.Fatal("Could not get all todos from db. Error: ", err)
	}

	tmpl := template.Must(template.ParseFiles("static/index.html"))

	err = tmpl.Execute(w, todos)
	if err != nil {
		log.Fatal("Could not execute template. Error: ", err)
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form", err)
	}

	err = model.CreateTodo(r.FormValue("todo"))
	if err != nil {
		log.Fatal("Could not create todo item. Error: ", err)
	}

	sendTodos(w)
}

func markTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Println("Could not parse id", err)
	}

	err = model.ChangeDoneStatus(id)
	if err != nil {
		fmt.Println("Could not update todo", err)
	}

	sendTodos(w)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		fmt.Println("Could not parse id", err)
	}

	err = model.Delete(id)
	if err != nil {
		fmt.Println("Could not delete", err)
	}

	sendTodos(w)
}

func SetupAndRun() {
	mux := mux.NewRouter()

	// create handlers/routes
	mux.HandleFunc("/", loadAlltodos)
	mux.HandleFunc("/create", createTodo).Methods("POST")
	mux.HandleFunc("/todo/{id}", markTodo).Methods("PUT")
	mux.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")

	// load static files
	mux.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// start server
	log.Fatal(http.ListenAndServe(":8000", mux))
}
