package main

import (
	"errors"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

type service struct {
	Data           map[string][]*Config `json:"Configuration groups"`
	Configurations []*Config            `json:"Configurations"`
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

	rt, err := decodeConfigGroupBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	configGroupId := createId()
	if len(srvc.Data) < 1 {
		srvc.Data = map[string][]*Config{}
	}
	var cfgs []*Config
	for i := 0; i < len(rt); i++ {
		configId := createId()
		var cfg Config
		cfg.Entries = make(map[string]string)
		cfg.Entries = rt[i].Entries
		cfg.Id = configId
		cfgs = append(cfgs, &cfg)
	}
	srvc.Data[configGroupId] = cfgs
	displayData := map[string][]*Config{}
	displayData[configGroupId] = cfgs
	renderJSON(w, displayData)
}

func (srvc *service) createConfigHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Contesnt-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeConfigBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cfg Config
	cfg.Entries = rt.Entries
	cfg.Id = createId()

	if len(srvc.Configurations) < 1 {
		srvc.Configurations = []*Config{}
	}
	srvc.Configurations = append(srvc.Configurations, &cfg)
	renderJSON(w, &cfg)
}

func (srvc *service) expandConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Contesnt-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeConfigBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, ok := srvc.Data[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	var cfg Config
	cfg.Id = createId()
	cfg.Entries = make(map[string]string)
	cfg.Entries = rt.Entries
	var cfgs []*Config
	cfgs = srvc.Data[id]
	cfgs = append(cfgs, &cfg)
	srvc.Data[id] = cfgs
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
	renderJSON(w, task)
}

func (srvc *service) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	for _, config := range srvc.Configurations {
		if config.Id == id {
			renderJSON(w, config)
			return
		}
	}
	for _, cfGroup := range srvc.Data {
		for _, v := range cfGroup {
			if v.Id == id {
				renderJSON(w, v)
				return
			}
		}
	}
	err := errors.New("key not found")
	http.Error(w, err.Error(), http.StatusNotFound)
}

func (srvc *service) delConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if _, ok := srvc.Data[id]; ok {
		delete(srvc.Data, id)
		renderJSON(w, srvc.Data[id])
		return
	}
	err := errors.New("key not found")
	http.Error(w, err.Error(), http.StatusNotFound)
}

func (srvc *service) delConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	for i := 0; i < len(srvc.Configurations); i++ {
		if srvc.Configurations[i].Id == id {
			displayConfig := srvc.Configurations[i]
			srvc.Configurations = remove(srvc.Configurations, i)
			renderJSON(w, displayConfig)
			return
		}
	}
	for cfGroup := range srvc.Data {
		for i := 0; i < len(srvc.Data[cfGroup]); i++ {
			if srvc.Data[cfGroup][i].Id == id {
				displayConfig := srvc.Data[cfGroup][i]
				srvc.Data[cfGroup] = remove(srvc.Data[cfGroup], i)
				renderJSON(w, displayConfig)
				return
			}
		}
	}
	err := errors.New("key not found")
	http.Error(w, err.Error(), http.StatusNotFound)
}

func (srvc *service) getAllHandler(w http.ResponseWriter, req *http.Request) {
	renderJSON(w, srvc)
}

func remove(slice []*Config, i int) []*Config {
	return append(slice[:i], slice[i+1:]...)
}
