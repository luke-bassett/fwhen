package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Title: "FormulaWhen"}
	t, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
