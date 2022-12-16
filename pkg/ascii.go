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
	input := r.FormValue("input")        // create a variable
	inputBanner := r.FormValue("banner") // create a variable
	d.Input = ""
	arr := Startascii(input, inputBanner)
	d.Input += arr
	d.Banner = inputBanner
	// fmt.Println(d.Input)
	fmt.Println(d)

	res.ExecuteTemplate(w, "ascii", d)
}
