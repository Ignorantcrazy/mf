package main

import (
	"encoding/json"
	"net/http"
)

func ResponseWithJson(w http.ResponseWriter, v interface{}, status int) {
	j, _ := json.Marshal(v)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}
