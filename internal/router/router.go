package router

import (
	"net/http"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)

	return mux
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is just an health check endpoint {Status : OK}"))
}
