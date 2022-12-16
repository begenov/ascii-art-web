package art

import (
	"fmt"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	res, err := template.ParseFiles("templates/index.html", "templates/header.html")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(1)
	res.ExecuteTemplate(w, "index", nil)
}
