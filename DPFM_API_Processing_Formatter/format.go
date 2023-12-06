package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-project-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToProjectUpdates(header dpfm_api_input_reader.Project) *ProjectUpdates {
	data := project

	return &ProjectUpdates{
		ProjectCode: data.ProjectCode,
	}
}

func ConvertToNetworkUpdates(networkUpdates dpfm_api_input_reader.Network) *NetworkUpdates {
	data := networkUpdates

	return &NetworkUpdates{
		project: data.Project,
	}
}

func ConvertToWBSElementUpdates(wBSElementUpdates dpfm_api_input_reader.WBSElement) *WBSElementUpdates {
	data := wBSElementUpdates

	return &WBSElementUpdates{
		Project: data.Project,
	}
}
