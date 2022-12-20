package art

import (
	"fmt"
	"net/http"
	"text/template"
)

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodPost {
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)

		// Вызываем метод w.WriteHeader() для возвращения статус-кода 405
		// и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".
		w.WriteHeader(405)
		fmt.Fprint(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))

		// Затем мы завершаем работу функции вызвав "return", чтобы
		// последующий код не выполнялся.
		return
	}

	res, err := template.ParseFiles("templates/ascii.html", "templates/header.html") // parses HTML files
	Check(err)
	input := r.FormValue("input")        // create a variable for input
	inputBanner := r.FormValue("banner") // create a variable for Banner
	count := 2
	for i := range r.Form {
		if i != "banner" {
			count--
		} else if i != "input" {
			count--
		} else if count != 2 {
			fmt.Println("ASd")
			fmt.Fprint(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
			return
		} else {
			continue
		}
	}

	ASCII.Input = ""
	arr, err := Startascii(input, inputBanner) // start ascii art
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// fmt.Print(arr)
	if arr == "Ошибка" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	ASCII.Input += arr
	res.ExecuteTemplate(w, "ascii", ASCII) //  run ascii.html
}
