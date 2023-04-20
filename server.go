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

func (ts *postServer) createPostHandler(w http.ResponseWriter, req *http.Request) {
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

	id := createId()
	rt.Id = id
	ts.data[id] = rt
	renderJSON(w, rt)
}

func (ts *postServer) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allTasks := []*RequestPost{}
	for _, v := range ts.data {
		allTasks = append(allTasks, v)
	}

	renderJSON(w, allTasks)
}

func (ts *postServer) getPostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := ts.data[id]
	if !ok {
		err := errors.New("key not found")
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
