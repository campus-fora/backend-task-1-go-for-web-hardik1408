package main

import (
	"html/template"
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", formHandler)
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	var name string = "" // declare and initialize name variable

	if r.Method == "POST" {
		// Handle form submission
		name = r.FormValue("name")
		email := r.FormValue("input")
		// Do something with name and email
		fmt.Println(name, email)
	}

	// Generate HTML code for the form
	tmpl := template.Must(template.ParseFiles("C:/Users/hardi/OneDrive/Desktop/campus fora/backend/index.html"))
	tmpl.Execute(w, nil)

}
