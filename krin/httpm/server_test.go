package httpm

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestRunServer(t *testing.T) {

	fn := func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("hello"))
		if err != nil {
			log.Fatal(err)
		}
	}
	h := http.HandlerFunc(fn)
	//s := Server{8087, http.FileServer(http.Dir("."))} // serve files
	s := Server{8087, h}
	err := s.serve()
	if err != nil {
		fmt.Print(err)
	}
}
