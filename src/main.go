package main

import "net/http"

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World\n"))
	})
	http.ListenAndServe(":7000", http.DefaultServeMux)
}
