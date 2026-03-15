package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHttpServer(httpHandler *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {

	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandlerCreateTask)
	

	if err:= http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err,http.ErrServerClosed) {
			return nil
		} else {
			return err
		}
	}

	return nil
}
