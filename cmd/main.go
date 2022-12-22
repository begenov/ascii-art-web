package main

import (
	"fmt"
	"net/http"
	"web/pkg"
)

// func main() {
// 	mux := http.NewServeMux()
// 	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
// 	mux.HandleFunc("/", pkg.Home)
// 	mux.HandleFunc("/ascii", pkg.Ascii)
// 	http.ListenAndServe(":8080", mux)
// }
func main() {
	fmt.Println("Web server: http://localhost:8080/")
	http.HandleFunc("/", pkg.Home)
	http.ListenAndServe(":8080", nil)
}
