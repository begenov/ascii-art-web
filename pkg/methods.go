package pkg

import (
	"fmt"
	"html/template"
	"net/http"
	ascii "web/ascii"
)

func Get(w http.ResponseWriter, r *http.Request) {
	tp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Error(500, w, r)
		return
	}
	tp.ExecuteTemplate(w, "index.html", nil)
}

func Post(w http.ResponseWriter, r *http.Request) {
	tp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		Error(500, w, r)
		return
	}
	intput := r.FormValue("input")
	fmt.Println(intput)
	bannerinput := r.FormValue("banner")
	result, err := ascii.Startascii(intput, bannerinput)
	if err != nil {
		Error(400, w, r)
		return
	}
	tp.ExecuteTemplate(w, "index.html", result)
}
