package main

import (
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	html := "<h1>{{.Title}}<h1/>"
	t, err := template.New("index").Parse(html)
	if err != nil {
		log.Printf("failed to parse index template, error: %v", err)
		return
	}
	data := map[string]string{
		"Title": "Store App! :)",
	}
	if err := t.Execute(w, data); err != nil {
		log.Printf("failed to execute index template, error: %v", err)
		return
	}
}

func main() {
	log.Println("server running on :8080")
	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
