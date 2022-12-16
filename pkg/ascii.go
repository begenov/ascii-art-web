package art

import (
	"fmt"
	"net/http"
	"text/template"
)

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { // check method
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	res, err := template.ParseFiles("templates/ascii.html", "templates/header.html") // parses HTML files
	Check(err)
	input := r.FormValue("input")        // create a variable for input
	inputBanner := r.FormValue("banner") // create a variable for Banner
	ASCII.Input = ""
	arr := Startascii(input, inputBanner) // start ascii art
	fmt.Print(arr)
	if arr == "Ошибка" {
		fmt.Fprint(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	ASCII.Input += arr
	res.ExecuteTemplate(w, "ascii", ASCII) //  run ascii.html
}
