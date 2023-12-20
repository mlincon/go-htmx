package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	fmt.Println("hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))

		// films is a key-value pair. The values are a list/array/slice of Film object
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		log.Print("HTMX request received")
		log.Print(r.Header.Get("HX-Request")) // if true it indicates that the request is comming from htmx

		// simulate some latency/external dependency like database to show the loading icon
		// on submit click
		time.Sleep(1 * time.Second)

		title := r.PostFormValue("title")       // get the data from the index.html input with the name "title"
		director := r.PostFormValue("director") // get the data from the index.html input with the name "director"

		/* This shows one option to render the html template, where we create a simple
		string and pass it onto the template

		htmlStr := fmt.Sprintf(
			"<li class='list-group-item bg-primary text-width'>"+
				"%s - %s"+
				"</li>", title, director,
		)

		// now create a new Template object. The name of the object does not matter
		// here it is called "t"
		// parse the formatted string to the parser
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
		*/

		/* Alternatively use the ExecuteTemplate to target the block in the html
		as shown below. This way we do not have to hard-code any html within the code
		*/

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "film-list-element", Film{
			Title:    title,
			Director: director,
		})
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
