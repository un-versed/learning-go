package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("./public/*.html"))

	http.HandleFunc("/hello-world", func(w http.ResponseWriter, r *http.Request) {
		user := user{Name: "Unversed", Email: "gu.bat98@gmail.com"}
		templates.ExecuteTemplate(w, "index.html", user)
	})

	log.Fatal(http.ListenAndServe(":3333", nil))
}
