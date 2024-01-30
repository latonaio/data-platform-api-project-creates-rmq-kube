package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-project-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-project-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"encoding/json"
	"time"

	"golang.org/x/xerrors"
)

func ConvertToProjectCreates(sdc *dpfm_api_input_reader.SDC) (*Project, error) {
	data := sdc.Project

	project, err := TypeConverter[*Project](data)
	if err != nil {
		return nil, err
	}
	// general.CreationDate = *getSystemDatePtr()
	// general.CreationTime = *getSystemTimePtr()
	// general.LastChangeDate = getSystemDatePtr()
	// general.LastChangeTime = getSystemTimePtr()

	return project, nil
}

func ConvertToNetworkCreates(sdc *dpfm_api_input_reader.SDC) (*[]Network, error) {
	network := make([]Network, 0)

	for _, data := range sdc.Project.Network {
		network, err := TypeConverter[*Network](data)
		if err != nil {
			return nil, err
		}

		networks = append(networks, *network)
	}

	return &networks, nil
}

func ConvertToWBSElementCreates(sdc *dpfm_api_input_reader.SDC) (*[]WBSElement, error) {
	wBSElement := make([]WBSElement, 0)

	for _, data := range sdc.Project.WBSElement {
		wBSElement, err := TypeConverter[*WBSElement](data)
		if err != nil {
			return nil, err
		}

		wBSElements = append(wBSElements, *wBSElement)
	}

	return &wBSElements, nil
}

func ConvertToProjectUpdates(projectData dpfm_api_input_reader.Project) (*Project, error) {
	data := projectData

	project, err := TypeConverter[*Project](data)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func ConvertToNetworkUpdates(networkUpdates *[]dpfm_api_processing_formatter.NetworkUpdates) (*[]Network, error) {
	networks := make([]Network, 0)

	for _, data := range *networkUpdates {
		network, err := TypeConverter[*Network](data)
		if err != nil {
			return nil, err
		}

		networks = append(networks, *network)
	}

	return &networks, nil
}

func ConvertToWBSElementUpdates(wBSElementUpdates *[]dpfm_api_processing_formatter.WBSElement) (*[]WBSElement, error) {
	wBSElements := make([]WBSElement, 0)

	for _, data := range *wBSElementUpdates {
		wBSElement, err := TypeConverter[*WBSElement](data)
		if err != nil {
			return nil, err
		}

		wBSElements = append(wBSElements, *wBSElement)
	}

	return &wBSElements, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}

func getSystemDatePtr() *string {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// day := time.Now().In(jst)

	day := time.Now()
	res := day.Format("2006-01-02")
	return &res
}

func getSystemTimePtr() *string {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// day := time.Now().In(jst)

	day := time.Now()
	res := day.Format("15:04:05")
	return &res
}
