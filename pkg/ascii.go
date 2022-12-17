package art

import (
	"fmt"
	"net/http"
	"text/template"
)

func Ascii(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost { // check method
	// http.Redirect(w, r, "/", http.StatusSeeOther)
	// return

	//}
	if r.Method != http.MethodPost {
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)

		// Вызываем метод w.WriteHeader() для возвращения статус-кода 405
		// и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".
		w.WriteHeader(405)
		w.Write([]byte("GET-Метод запрещен!"))

		// Затем мы завершаем работу функции вызвав "return", чтобы
		// последующий код не выполнялся.
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
