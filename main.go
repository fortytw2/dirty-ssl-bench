package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte("Hello World!"))
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
	// http.ListenAndServeTLS(":8080", "localhost.crt", "localhost.key", nil)
}
