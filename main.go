package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	models "github.com/evileric-com/events-svc/models"
	services "github.com/evileric-com/events-svc/services"
	"github.com/gorilla/mux"
)

func main() {
	repo := services.NewEventRepository(services.ConnectionSettings{
		Host:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASS"),
	})

	baseRouter := mux.NewRouter().StrictSlash(false)
	baseRouter.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		services.RespondWithJSON(w, 200, "pong")
	})

	baseRoute := baseRouter.PathPrefix("/api/v1/")

	subRoutes := baseRoute.Subrouter()
	subRoutes.HandleFunc("/events/{id}", PostEvent(repo)).Methods("POST")
	subRoutes.HandleFunc("/events/{id}", GetEvent(repo)).Methods("GET")

	http.ListenAndServe(":3001", baseRouter)
}

func PostEvent(repo services.EventRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Event
		services.ReadPayload(r, &e)

		a, b := json.Marshal(e)
		if b != nil {
			fmt.Println(string(a))
		} else {
			fmt.Println(b)
		}

		services.SaveEvent(repo, &e)
	}
}

func GetEvent(repo services.EventRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		event := services.GetEvent(repo, id)
		services.RespondWithJSON(w, 200, event)
	}
}
