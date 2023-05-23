package main

import "main/configstore"

// swagger:parameters deleteConfigById
type DeleteConfigByIdRequest struct {
	// Config ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters deleteConfigByIdAndVersion
type DeleteConfigByIdAndVersionRequest struct {
	// Config ID
	// in: path
	Id string `json:"id"`

	// Config version
	// in: path
	Version string `json:"version"`
}

// swagger:parameters getConfigById
type GetConfigByIdRequest struct {
	// Config ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters getConfigByIdAndVersion
type GetConfigByIdAndVersionRequest struct {
	// Config ID
	// in: path
	Id string `json:"id"`

	// Config version
	// in: path
	Version string `json:"version"`
}

// swagger:parameters getCfGroupById
type GetCfGroupByIdRequest struct {
	// Group ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters getCfGroupByIdAndVersion
type GetCfGroupByIdAndVersionRequest struct {
	// Group ID
	// in: path
	Id string `json:"id"`

	// Group version
	// in: path
	Version string `json:"version"`
}

// swagger:parameters deleteGroupById
type DeleteCfGroupByIdRequest struct {
	// Group ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters deleteGroupByIdAndVersion
type DeleteCfGroupByIdAndVersionRequest struct {
	// Group ID
	// in: path
	Id string `json:"id"`

	// Group version
	// in: path
	Version string `json:"version"`
}

// swagger:parameters getGroupConfigByLabel
type GetGroupConfigByLabelRequest struct {
	// Group groupId
	// in: path
	GroupId string `json:"groupId"`

	// Group version
	// in: path
	Version string `json:"version"`

	// Config label
	// in: path
	Labels string `json:"label"`
}

// swagger:parameters getGroupConfigByIdAndLabel
type GetGroupConfigByIdAndLabelRequest struct {
	// Group groupId
	// in: path
	GroupId string `json:"groupId"`

	// Group version
	// in: path
	Version string `json:"version"`

	// Config label
	// in: path
	Labels string `json:"label"`

	// Config configId
	// in: path
	ConfigId string `json:"configId"`
}

// swagger:parameters deleteGroupConfigByLabelAndId
type DeleteGroupConfigByLabelAndId struct {
	// Group groupId
	// in: path
	GroupId string `json:"groupId"`

	// Group version
	// in: path
	Version string `json:"version"`

	// Config label
	// in: path
	Labels string `json:"label"`

	// Config configId
	// in: path
	ConfigId string `json:"configId"`
}

// swagger:parameters config createConfig
type RequestConfigBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/Config"
	//  required: true
	Body configstore.Config `json:"body"`
}

// swagger:parameters cfgroup createConfigGroup
type RequestCfGroupBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/CfGroup"
	//  required: true
	Body configstore.CfGroup `json:"body"`
}

// swagger:parameters config putConfig
type RequestPutConfigBody struct {
	// Config ID
	// in: path
	Id string `json:"id"`

	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/Config"
	//  required: true
	Body configstore.Config `json:"body"`
}

// swagger:parameters config expandConfigGroup
type RequestExpandConfigGroupBody struct {
	// CfGroup ID
	// in: path
	Id string `json:"id"`

	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/Config"
	//  required: true
	Body configstore.Config `json:"body"`
}
