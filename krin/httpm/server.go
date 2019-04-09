package httpm

import (
	"fmt"
	"net/http"
)

type Server struct {
	port    int
	handler http.Handler
}

func (s *Server) serve() error {

	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.port), s.handler)
	if err != nil {
		return err
	}
	return nil
}
