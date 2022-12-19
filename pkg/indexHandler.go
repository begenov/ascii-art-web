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
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)

		// Вызываем метод w.WriteHeader() для возвращения статус-кода 405
		// и вызывается метод w.Write() для возвращения тела-ответа с текстом "Метод запрещен".
		w.WriteHeader(405)
		w.Write([]byte("Метод запрещен!"))

		// Затем мы завершаем работу функции вызвав "return", чтобы
		// последующий код не выполнялся.
		return
	}
	res, err := template.ParseFiles("templates/index.html", "templates/header.html") // parses HTML files
	Check(err)                                                                       // check for error
	res.ExecuteTemplate(w, "index", nil)                                             // run index.html
}
