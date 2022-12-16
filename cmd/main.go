package main

import (
	"fmt"
	"net/http"

	art "art/pkg"
)

func main() {
	fmt.Println("Запуск веб-приложения http://localhost:8080/")

	// in-app routing systems "/"
	http.HandleFunc("/", art.IndexHandler)

	// in-app routing systems "/ascii"
	http.HandleFunc("/ascii", art.Ascii)

	// file request handler static functions
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// launch web application
	http.ListenAndServe(":8080", nil)
}
