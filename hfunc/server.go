package forum

import "net/http"

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
}

func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!="/"{
		err
	}
}
