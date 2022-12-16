package art

import (
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	res, err := template.ParseFiles("templates/index.html", "templates/header.html") // parses HTML files
	Check(err)                                                                       // check for error
	res.ExecuteTemplate(w, "index", nil)                                             // run index.html
}
