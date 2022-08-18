package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) Run(port string) {
	err := http.ListenAndServe(port, a.Router)
	if err != nil {
		return
	}
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/health", func(w http.ResponseWriter, request *http.Request) {
		response, _ := json.Marshal(struct {
			UP string `json:"up"`
		}{UP: "UP"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(response)
		if err != nil {
			return
		}
	}).Methods(http.MethodGet)
}
