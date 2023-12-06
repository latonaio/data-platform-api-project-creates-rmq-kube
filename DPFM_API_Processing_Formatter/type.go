package dpfm_api_processing_formatter

type ProjectUpdates struct {
	ProjectCode string `json:"ProjectCode"`
}

type NetworkUpdates struct {
	Project int `json:"Project"`
}

type WBSElementUpdates struct {
	Project int `json:"Project"`
}
