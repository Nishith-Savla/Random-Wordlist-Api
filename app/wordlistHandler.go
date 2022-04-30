package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Nishith-Savla/Random-Wordlist-Api/service"
	"github.com/gorilla/mux"
)

type WordlistHandler struct {
	service service.WordlistService
}

func (h WordlistHandler) getWords(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var limit int
	var err error
	if limit, err = strconv.Atoi(vars["limit"]); err != nil {
		writeJSONResponse(w, 400, err)
		return
	}

	words := h.service.GetWords(limit)

	writeJSONResponse(w, 200, words)
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Panic(err.Error())
	}
}
