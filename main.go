package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const tpl = `
	RequestURI: %v
	Host:       %v
	Form:       %v
	Some:       %v
`

func serveForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, fmt.Sprintf(tpl, 
		r.RequestURI, 
		r.Host, 
		r.Form,
		r.Form.Get("Some"),
	))
}

// sends files to the browser
func serveFile(w http.ResponseWriter, r *http.Request) {
	var err error
	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}
	http.ServeFile(w, r, filepath.Join(wd, r.URL.Path))
}

func main() {
	http.HandleFunc("/form", serveForm)
	http.HandleFunc("/", serveFile)
	log.Fatal(http.ListenAndServe(":8080", nil))
}