package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func decodeBody(r io.Reader) ([]map[string]string, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var configs []map[string]string
	if err := dec.Decode(&configs); err != nil {
		return nil, err
	}
	return configs, nil
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createId() string {
	return uuid.New().String()
}
