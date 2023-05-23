package main

import (
	"encoding/json"
	"io"
	"main/configstore"
	"net/http"

	"github.com/google/uuid"
)

func decodeConfigGroupBody(r io.Reader) (*configstore.CfGroup, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var cfgroup *configstore.CfGroup
	if err := dec.Decode(&cfgroup); err != nil {
		return nil, err
	}
	return cfgroup, nil
}

func decodeConfigBody(r io.Reader) (*configstore.Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var config *configstore.Config
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
