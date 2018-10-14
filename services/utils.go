package services

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	models "github.com/evileric-com/events-svc/models"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ReadPayload(r *http.Request, payload *models.Event) *models.Event {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err == nil {
		fmt.Println("no err")
		e := json.Unmarshal(body, payload)
		if e != nil {
			fmt.Println("Could not read event json", e)
		}

		return payload
	} else {
		fmt.Println("Could not parse event", err)
	}

	return nil
}
