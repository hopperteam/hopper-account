package web

import (
	"encoding/json"
	"net/http"
)

type apiResultElement struct {
	Status string `json:"status"`
	Result interface{} `json:"result"`
}

type apiErrorElement struct {
	Status string `json:"status"`
	Reason interface{} `json:"reason"`
}

func apiResult(w http.ResponseWriter, result interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&apiResultElement{
		Status: "success",
		Result: result,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func apiError(w http.ResponseWriter, error string, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(&apiErrorElement{
		Status: "error",
		Reason: error,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

