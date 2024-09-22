package webserver

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler struct {
	handlerFunc http.HandlerFunc
	path        string
	method      string
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

func (s *WebServer) AddHandler(path string, handleFunc http.HandlerFunc, method string) {
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
		s.Router.Method(handler.method, handler.path, handler.handlerFunc)
	}
	http.ListenAndServe(":"+s.WebServerPort, s.Router)
}
