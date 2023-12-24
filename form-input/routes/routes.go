package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	// w.Write([]string{"Apple", "Orange"})
}

func Serve() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", helloWorld).Methods("GET")
	mux.HandleFunc("/test", helloWorld).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", mux))
}
