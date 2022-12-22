package pkg

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		if r.Method == http.MethodGet {
			Get(w, r)
		} else {
			Error(405, w, r)
			return
		}
	} else if r.URL.Path == "/ascii-art" {
		if r.Method == http.MethodPost {
			Post(w, r)
		} else {
			Error(405, w, r)
			return
		}
	} else {
		Error(404, w, r)
		return
	}
}
