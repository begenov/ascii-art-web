package pkg

import (
	"html/template"
	"net/http"
)

type Errors struct {
	Num      int    `json:"num"`
	ResError string `json:"reserror"`
}

func Error(num int, w http.ResponseWriter, r *http.Request) {
	var reserror string
	if num == 405 {
		reserror = "Method Not Allowed"
	} else if num == 500 {
		reserror = "Internal Server Error"
	} else if num == 400 {
		reserror = "Bad Request"
	} else if num == 404 {
		reserror = "Not Found"
	}
	w.WriteHeader(num)
	tem, err := template.ParseFiles("templates/error.html")
	if err != nil {
		Error(500, w, r)
		return
	}
	var Result Errors
	Result.Num = num
	Result.ResError = reserror
	tem.ExecuteTemplate(w, "error.html", Result)
}
