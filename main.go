package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@host:hads1408@tcp(127.0.0.1:8080)/campusfora")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", formHandler)
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	var name string = ""

	if r.Method == "POST" {
		// Handle form submission
		name = r.FormValue("name")
		id := r.FormValue("input")
	}

	stmt, err := db.Prepare("INSERT INTO data(id,input) VALUES(?,?)", id, name)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var id2 string = ""
	if r.Method == "DELETE" {
		id2 := r.FormValue("id")
	}

	stm, err := db.Prepare("DELETE FROM data WHERE id=?", id2)
	if err != nil {
		log.Fatal(err)
	}
	defer stm.Close()

	tmpl := template.Must(template.ParseFiles("C:/Users/hardi/OneDrive/Desktop/campus fora/backend/index.html"))
	tmpl.Execute(w, nil)

}
