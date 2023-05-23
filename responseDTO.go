package main

import "main/configstore"

// swagger:response ResponseConfig
type ResponseConfig struct {
	// Id of the config
	// in: string
	Id string `json:"Id"`

	// Entries of the config
	// in: string
	Entries map[string]string `json:"Entries"`

	// Labels of the config
	// in: string
	Labels string `json:"labels"`

	// Version of the config
	// in: string
	Version string `json:"version"`
}

//swagger:response ResponseCfGroup
type CfGroup struct {
	// Id of the cfgroup
	// in: string
	Id string `json:"Id"`

	// Configurations of the cfgroup
	// in: []*Config
	Configurations []*configstore.Config `json:"Configurations"`

	// Version of the cfgroup
	// in: string
	Version string `json:"version"`
}

// swagger:response ErrorResponse
type ErrorResponse struct {
	// Error status code
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

// swagger:response NoContentResponse
type NoContentResponse struct{}
