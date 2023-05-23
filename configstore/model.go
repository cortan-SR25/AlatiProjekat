package configstore

//swagger:model Config
type Config struct {
	// Id of the config
	// in: string
	Id string `json:"Id"`

	// Entries of the config
	// in: map[string]string
	Entries map[string]string `json:"Entries"`

	// Labels of the config
	// in: string
	Labels string `json:"labels"`

	// Version of the config
	// in: string
	Version string `json:"version"`
}

//swagger:model CfGroup
type CfGroup struct {
	// Id of the cfgroup
	// in: string
	Id string `json:"Id"`

	// Configurations of the cfgroup
	// in: []*Config
	Configurations []*Config `json:"Configurations"`

	// Version of the cfgroup
	// in: string
	Version string `json:"version"`
}
