package main

import (
	"errors"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

type service struct {
	Data map[string][]*Config `json:"Configuration groups"`
}

func (srvc *service) createConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	configId := createId()
	if len(srvc.Data) == 0 {
		srvc.Data = map[string][]*Config{}
	}
	var cfgs []*Config
	for i := 0; i < len(rt); i++ {
		var cfg Config
		cfg.Entries = make(map[string]string)
		cfg.Entries = rt[i]
		cfgs = append(cfgs, &cfg)
	}
	srvc.Data[configId] = cfgs
	renderJSON(w, rt)
}

func (srvc *service) createConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, ok := srvc.Data[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if len(task) < 2 {
		err := errors.New("can't add a configuration to another one that's outside of a group")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var cfgs []*Config

	for i := 0; i < len(rt); i++ {
		var cfg Config
		cfg.Entries = make(map[string]string)
		cfg.Entries = rt[i]
		cfgs = append(cfgs, &cfg)
	}
	for i := 0; i < len(cfgs); i++ {
		srvc.Data[id] = append(srvc.Data[id], cfgs[i])
	}
	renderJSON(w, srvc.Data[id])
}

func (srvc *service) getConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := srvc.Data[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if len(task) < 2 {
		err := errors.New("This is a single configuration not a group")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

func (srvc *service) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := srvc.Data[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if len(task) > 1 {
		err := errors.New("This is group of configurations not a single one")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

func (srvc *service) delConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if v, ok := srvc.Data[id]; ok {
		if len(v) < 2 {
			err := errors.New("can't delete one configuration, must be a group")
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		delete(srvc.Data, id)
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (srvc *service) delConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if v, ok := srvc.Data[id]; ok {
		if len(v) > 1 {
			err := errors.New("can't delete a group, must be a single configuration")
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		delete(srvc.Data, id)
		renderJSON(w, v)
		return
	}
	err := errors.New("key not found")
	http.Error(w, err.Error(), http.StatusNotFound)
}
