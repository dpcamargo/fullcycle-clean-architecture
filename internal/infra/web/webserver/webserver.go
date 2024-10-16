package webserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Method string

const (
	POST Method = "POST"
	GET  Method = "GET"
)

type Handler struct {
	handlerFunc http.HandlerFunc
	path        string
	method      Method
}

type WebServer struct {
	Router        chi.Router
	Handlers      []Handler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make([]Handler, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handleFunc http.HandlerFunc, method Method) {
	handler := Handler{
		handlerFunc: handleFunc,
		path:        path,
		method:      method,
	}
	s.Handlers = append(s.Handlers, handler)
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		s.Router.Method(string(handler.method), handler.path, handler.handlerFunc)
	}
	err := http.ListenAndServe(s.WebServerPort, s.Router)
	if err != nil {
		panic(err)
	}
}
