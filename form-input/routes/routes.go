package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func testString(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func getCategories() []string {
	return []string{"Apple", "Orange"}
}

func Serve() {
	mux := mux.NewRouter()

	mux.HandleFunc("/", testString).Methods("GET")
	mux.HandleFunc("/test", testString).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", mux))
}
