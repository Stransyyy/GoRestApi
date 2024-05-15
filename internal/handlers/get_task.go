package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Stransyyy/GoRestApi/api"
	"github.com/Stransyyy/GoRestApi/internal/tools"
	esq "github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	var params = api.TaskParams{}
	var decoder *esq.Decoder = esq.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.TaskDetails
	tokenDetails = (*database).GetTaskDetails(params.Username)
	if tokenDetails == nil {
		log.Error("No Task Details Found")
		api.InternalErrorHandler(w)
		return
	}

	var reponse = api.TaskResponse{
		Code:   http.StatusOK,
		Result: "Number of tasks: " + string(tokenDetails.Tasks),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(reponse)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
