package art

import (
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("Метод запрещен!"))

		return
	}
	res, err := template.ParseFiles("templates/index.html", "templates/header.html") // parses HTML files
	Check(err)                                                                       // check for error
	res.ExecuteTemplate(w, "index", nil)                                             // run index.html
}
