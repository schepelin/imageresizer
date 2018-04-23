package http

import (
	"io"
	"log"
	"net/http"
)

func createImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {

	}
	io.WriteString(w, "Hello World")
}

type Server struct {
	Addr   string
	Logger *log.Logger
}

func (s *Server) Start() {
	http.HandleFunc("/images", createImage)
	s.Logger.Printf("Server started at %s", s.Addr)

	s.Logger.Fatal(http.ListenAndServe(s.Addr, nil))
}

func NewServer(addr string, logger *log.Logger) *Server {
	return &Server{
		Addr:   addr,
		Logger: logger,
	}
}
