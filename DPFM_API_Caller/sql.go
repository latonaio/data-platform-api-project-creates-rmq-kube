package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-project-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-project-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-project-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var project *dpfm_api_output_formatter.Project
	var network *[]dpfm_api_output_formatter.Network
	var wBSElement *[]dpfm_api_output_formatter.WBSElement
	for _, fn := range accepter {
		switch fn {
		case "Project":
			project = c.projectCreateSql(nil, mtx, input, output, errs, log)
		case "Network":
			network = c.networkCreateSql(nil, mtx, input, output, errs, log)
		case "WBSElement":
			wBSElement = c.wBSElementCreateSql(nil, mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Project:    project,
		Network:    network,
		WBSElement: wBSElement,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var project *dpfm_api_output_formatter.Project
	var network *[]dpfm_api_output_formatter.Network
	var wBSElement *[]dpfm_api_output_formatter.WBSElement
	for _, fn := range accepter {
		switch fn {
		case "Project":
			project = c.projectUpdateSql(mtx, input, output, errs, log)
		case "Network":
			network = c.networkUpdateSql(mtx, input, output, errs, log)
		case "WBSElement":
			wBSElement = c.wBSElementUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		Project:    project,
		Network:    network,
		WBSElement: wBSElement,
	}

	return data
}

func (c *DPFMAPICaller) projectCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Project {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_business_partner_general_dataの更新
	projectData := input.Project
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": projectData, "function": "ProjectProject", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Project Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToProjectCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) networkCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Network {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_business_partner_fin_inst_dataの更新
	for i := range input.Project.Network {
		input.Header.Project.Network[i].Project = input.Project.Network
		networkData := input.Project.Network[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": networkData, "function": "ProjectNetwork", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Network Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToNetworkCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) wBSElementCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.WBSElement {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	// data_platform_business_partner_role_dataの更新
	for i := range input.Project.WBSElement {
		input.Project.WBSElement[i].Project = input.Project.WBSElement
		wBSElementData := input.Project.WBSElement[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": wBSElementData, "function": "ProjectWBSElement", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "WBSElement Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToWBSElementCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) projectUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Project {
	project := input.Project
	projectData := dpfm_api_processing_formatter.ConvertToProjectUpdates(project)

	sessionID := input.RuntimeSessionID
	if projectIsUpdate(projectData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": projectData, "function": "ProjectProject", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Project Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToProjectUpdates(project)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) networkUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Network {
	req := make([]dpfm_api_processing_formatter.NetworkUpdates, 0)
	sessionID := input.RuntimeSessionID

	project := input.Project
	for _, network := range project.Network {
		networkData := *dpfm_api_processing_formatter.ConvertToNetworkUpdates(project, network)

		if networkIsUpdate(&networkData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": networkData, "function": "ProjectNetwork", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Network Data cannot update"
				return nil
			}
		}
		req = append(req, networkData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToNetworkUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) wBSElementUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.WBSElement {
	req := make([]dpfm_api_processing_formatter.WBSElementUpdates, 0)
	sessionID := input.RuntimeSessionID

	project := input.Project
	for _, wBSElement := range project.WBSElement {
		wBSElementData := *dpfm_api_processing_formatter.ConvertToWBSElementUpdates(project, wBSElement)

		if wBSElementIsUpdate(&wBSElementData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": wBSElementData, "function": "ProjectWBSElement", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "WBSElement Data cannot update"
				return nil
			}
		}
		req = append(req, wBSElement)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToWBSElementUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func projectIsUpdate(project *dpfm_api_processing_formatter.ProjectUpdates) bool {
	projectCode := project.ProjectCode

	return !(projectCode == 0)
}

func networkIsUpdate(network *dpfm_api_processing_formatter.NetworkUpdates) bool {
	project := network.Project

	return !(project == 0)
}

func wBSElementIsUpdate(wBSElement *dpfm_api_processing_formatter.WBSElementUpdates) bool {
	businessPartner := wBSElement.BusinessPartner
	businessPartnerRole := wBSElement.BusinessPartnerRole
	validityEndDate := wBSElement.ValidityEndDate
	validityStartDate := wBSElement.ValidityStartDate

	return !(businessPartner == 0 || businessPartnerRole == "" || validityEndDate == "" || validityStartDate == "")
}
