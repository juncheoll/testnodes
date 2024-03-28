package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	http.Handler
}

func CreateHandler() *Handler {
	mux := mux.NewRouter()
	handler := &Handler{
		Handler: mux,
	}

	mux.HandleFunc("/provider/{provider:[a-z-_]+}/model/{model:[a-z-_]+}/{version:[0-9]+}/infer/{tn:[0-9-_]+}", handler.testInferV2Handler).Methods("POST") // Inference version 2.0

	return handler
}
