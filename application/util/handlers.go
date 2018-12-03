package util

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithXML(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := xml.Marshal(payload)

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithoutError(w http.ResponseWriter, code int, payload interface{}, contentType string) {
	if contentType != "application/xml" {
		respondWithJSON(w, code, payload)
	} else {
		respondWithXML(w, code, payload)
	}
}
