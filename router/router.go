package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DemoController - Interface for demo requests
type DemoController interface {
	Demo(w http.ResponseWriter, r *http.Request)
}

// Setup returns router instance which is used in main package to register handlers.
func Setup(
	demoController DemoController,
) http.Handler {
	r := mux.NewRouter()

	// Demos endpoints
	r.HandleFunc(
		"/",
		demoController.Demo).Methods("GET").Name("demo")
	return r
}
