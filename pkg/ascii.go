package art

import (
	"fmt"
	"net/http"
	"text/template"
)

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	res, err := template.ParseFiles("templates/ascii.html", "templates/header.html")
	if err != nil {
		return
	}
	input := r.FormValue("input")
	inputBanner := r.FormValue("banner")
	// fmt.Println(inputBanner)
	d.Input = ""
	arr := Startascii(input, inputBanner)
	d.Input += arr
	d.Banner = inputBanner
	// fmt.Println(d.Input)
	fmt.Println(d)

	res.ExecuteTemplate(w, "ascii", d)
}
