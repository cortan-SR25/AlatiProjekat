package main

import (
	"errors"
	"main/configstore"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

type configServer struct {
	store *configstore.ConfigStore
}

// swagger:route POST /cfgroup/ cfgroup createConfigGroup
// Add new configuration group
//
// responses:
//  415: ErrorResponse
//  400: ErrorResponse
//  200: ResponseCfGroup
func (cs *configServer) createConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
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

	cfGroup, err := cs.store.PostCfGroup(rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, cfGroup)
}

// swagger:route POST /config/ config createConfig
// Add new configuration
//
// responses:
//  415: ErrorResponse
//  400: ErrorResponse
//  200: ResponseConfig
func (cs *configServer) createConfigHandler(w http.ResponseWriter, req *http.Request) {
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

	config, err := cs.store.PostConfig(rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, config)
}

// swagger:route GET /configs/ config getAllConfigs
// Get all configurations
//
// responses:
//  200: []ResponseConfig
func (cs *configServer) getAllConfigsHandler(w http.ResponseWriter, req *http.Request) {
	allTasks, err := cs.store.GetAllConfigs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, allTasks)
}

// swagger:route GET /cfgroups/ cfgroup getAllCfGroups
// Get all configuration groups
//
// responses:
//  200: []ResponseCfGroup
func (cs *configServer) getAllCfGroupsHandler(w http.ResponseWriter, req *http.Request) {
	allTasks, err := cs.store.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, allTasks)
}

// swagger:route PUT /cfgroup/{id}/config/ cfgroup expandConfigGroup
// Add new configuration to group by ID
//
// responses:
//  415: ErrorResponse
//  400: ErrorResponse
//  200: ResponseCfGroup
func (cs *configServer) expandConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
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

	cfGroup, err := cs.store.PutGroupConfigByGroupId(rt, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, cfGroup)
}

// swagger:route GET /config/{id}/{version}/ config getConfigByIdAndVersion
// Get configuration by ID and version
//
// responses:
//  404: ErrorResponse
//  200: []ResponseConfig
func (cs *configServer) getConfigByIdAndVersionHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	task, err := cs.store.GetConfigByIdAndVersion(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route GET /config/{id}/ config getConfigById
// Get configuration by ID
//
// responses:
//  404: ErrorResponse
//  200: []ResponseConfig
func (cs *configServer) getConfigByIdHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, err := cs.store.GetConfigById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route DELETE /config/{id}/ config deleteConfigById
// Delete configuration by ID
//
// responses:
//  404: ErrorResponse
//  204: NoContentResponse
func (cs *configServer) deleteConfigByIdHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, err := cs.store.DeleteConfigById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route DELETE /config/{id}/{version} config deleteConfigByIdAndVersion
// Delete configuration by ID and version
//
// responses:
//  404: ErrorResponse
//  204: NoContentResponse
func (cs *configServer) deleteConfigByIdAndVersionHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	task, err := cs.store.DeleteConfigByIdAndVersion(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route GET /cfgroup/{id}/ cfgroup getCfGroupById
// Get configuration group by ID
//
// responses:
//  415: ErrorResponse
//  400: ErrorResponse
//  200: ResponseCfGroup
func (cs *configServer) getCfGroupByIdHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, err := cs.store.GetCfGroupById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route GET /cfgroup/{id}/{version}/ cfgroup getCfGroupByIdAndVersion
// Get configuration group by ID and version
//
// responses:
//  415: ErrorResponse
//  400: ErrorResponse
//  200: ResponseCfGroup
func (cs *configServer) getCfGroupByIdAndVersionHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	task, err := cs.store.GetCfGroupByIdAndVersion(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route DELETE /cfgroup/{id}/ cfgroup deleteGroupById
// Delete configuration group by ID
//
// responses:
//  404: ErrorResponse
//  204: NoContentResponse
func (cs *configServer) deleteGroupByIdHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, err := cs.store.DeleteCfGroupById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route DELETE /cfgroup/{id}/{version}/ cfgroup deleteGroupByIdAndVersion
// Delete configuration group by ID and version
//
// responses:
//  404: ErrorResponse
//  204: NoContentResponse
func (cs *configServer) deleteGroupByIdAndVersionHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["id"]
	task, err := cs.store.DeleteCfGroupByIdAndVersion(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route DELETE /cfgroup/{groupId}/{version}/config/{label}/{configId}/ cfgroup deleteGroupConfigByLabelAndId
// Delete configuration by configId label version and groupId
//
// responses:
//  404: ErrorResponse
//  204: NoContentResponse
func (cs *configServer) deleteGroupConfigByLabelAndIdHandler(w http.ResponseWriter, req *http.Request) {
	groupId := mux.Vars(req)["groupId"]
	version := mux.Vars(req)["version"]
	configId := mux.Vars(req)["configId"]
	labels := mux.Vars(req)["label"]
	task, err := cs.store.DeleteGroupConfigByLabelAndId(groupId, version, labels, configId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route GET /cfgroup/{groupId}/{version}/config/{label}/{configId}/ config getGroupConfigByIdAndLabel
// Get configuration by configId label version and groupId
//
// responses:
//  415: ErrorResponse
//  400: ErrorResponse
//  200: []ResponseConfig
func (cs *configServer) getGroupConfigByIdAndLabelHandler(w http.ResponseWriter, req *http.Request) {
	groupId := mux.Vars(req)["groupId"]
	version := mux.Vars(req)["version"]
	configId := mux.Vars(req)["configId"]
	label := mux.Vars(req)["label"]
	task, err := cs.store.GetGroupConfigByIdAndLabel(groupId, version, configId, label)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route GET /cfgroup/{groupId}/{version}/config/{label}/ config getGroupConfigByLabel
// Get configuration by label version and groupId
//
// responses:
//  415: ErrorResponse
//  400: ErrorResponse
//  200: []ResponseConfig
func (cs *configServer) getGroupConfigByLabelHandler(w http.ResponseWriter, req *http.Request) {
	groupId := mux.Vars(req)["groupId"]
	version := mux.Vars(req)["version"]
	label := mux.Vars(req)["label"]
	task, err := cs.store.GetGroupConfigByLabel(groupId, version, label)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route PUT /config/{id}/ config putConfig
// Put configuration by ID
//
// responses:
//  415: ErrorResponse
//  400: ErrorResponse
//  200: ResponseConfig
func (cs *configServer) putConfigHandler(w http.ResponseWriter, req *http.Request) {
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

	config, err := cs.store.PutConfig(rt, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, config)
}

func (cs *configServer) swaggerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./swagger.yaml")
}
