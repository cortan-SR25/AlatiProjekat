package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"
)

func decodeConfigGroupBody(r io.Reader) ([]*Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var configs []*Config
	if err := dec.Decode(&configs); err != nil {
		return nil, err
	}
	return configs, nil
}

func decodeConfigBody(r io.Reader) (*Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var config *Config
	if err := dec.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
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
