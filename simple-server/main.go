package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

const base = "https://jsonplaceholder.typicode.com/"

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "Title: %s" .Title}}</div>
<div>{{printf "Body: %s" .Body}}</div>`

func handler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(base + r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	var p post
	
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	tmpl := template.New("mine")
	tmpl.Parse(form)
	tmpl.Execute(w, p)

}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}