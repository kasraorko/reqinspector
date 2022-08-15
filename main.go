package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Request struct {
	Method         string
	URL            *url.URL
	Submissions    map[string][]string
	Header         http.Header
	Host           string
	RemoteAddr     string
	RequestURI     string
	ResponseHdears http.Header
}

func requestInspector(w http.ResponseWriter, req *http.Request) {

	data := Request{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.RemoteAddr,
		req.RequestURI,
		w.Header(),
	}

	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	if req.FormValue("output") == "json" {
		w.Header().Set("Content-Type", "application/json")
		payload, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = w.Write(payload)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		err := tpl.ExecuteTemplate(w, "index.html", data)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", requestInspector)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
