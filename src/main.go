package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// TODO: Create a way to parse the port and ip flags
// TODO: Separate code from here to a server file

func render(w http.ResponseWriter, tmpl string, p *Page) {

	t, err := template.ParseFiles(tmpl)

	if err != nil {
		log.Fatal("Something went won't while trying to parse the page")
	}

	t.Execute(w, p)
}

func root(w http.ResponseWriter, r *http.Request) {
	title := "./html/index.html"

	page, err := Load(title)

	if err != nil {
		fmt.Printf("Could not load the page %s\n", title)
	} else {
		render(w, title, page)
	}
}

func main() {

	http.HandleFunc("/", root)
	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Printf("Server running on %s:%s\n", "127.0.0.1", "8080")
}
