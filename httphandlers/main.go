package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	to := r.URL.Path[1:]
	_, err := fmt.Fprintf(w, "hello %s", to)
	//_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	_, err := fmt.Fprintf(w, "article : %s", title)
	if err != nil {
		log.Fatal(err)
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/view/", viewHandler)

	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		log.Fatal(err)
	}
}
