package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"flag"
)

//templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

//ServeHTTP handles the http Request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application")
	flag.Parse() // Parge them flags

	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	//Get the room join

	go r.run()
	log.Println("Starting web server on", *addr)
	//Start the web server
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
