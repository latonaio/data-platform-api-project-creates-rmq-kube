package sub_func_complementer

type SDC struct {
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	APIType             string   `json:"api_type"`
	Message             Message  `json:"message"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	SQLUpdateResult     *bool    `json:"sql_update_result"`
	SQLUpdateError      string   `json:"sql_update_error"`
	SubfuncResult       *bool    `json:"subfunc_result"`
	SubfuncError        string   `json:"subfunc_error"`
	ExconfResult        *bool    `json:"exconf_result"`
	ExconfError         string   `json:"exconf_error"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type Message struct {
	Project    *Project      `json:"Project"`
	Network    *[]Network    `json:"Network"`
	WBSElement *[]WBSElement `json:"WBSElement"`
}

type Project struct {
	Project               int     `json:"Project"`
	ProjectCode           string  `json:"ProjectCode"`
	ProjectDescription    string  `json:"ProjectDescription"`
	OwnerBusinessPartner  int     `json:"OwnerBusinessPartner"`
	OwnerPlant            string  `json:"OwnerPlant"`
	ProjectProfile        *string `json:"ProjectProfile"`
	ResponsiblePerson     *int    `json:"ResponsiblePerson"`
	ResponsiblePersonName *string `json:"ResponsiblePersonName"`
	PlannedStartDate      *string `json:"PlannedStartDate"`
	PlannedEndDate        *string `json:"PlannedEndDate"`
	ActualStartDate       *string `json:"ActualStartDate"`
	ActualEndDate         *string `json:"ActualEndDate"`
	CreationDate          string  `json:"CreationDate"`
	LastChangeDate        string  `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}

type Network struct {
	Project               int     `json:"Project"`
	WBSElement            int     `json:"WBSElement"`
	Network               int     `json:"Network"`
	NetworkDescription    *string `json:"NetworkDescription"`
	BusinessPartner       int     `json:"BusinessPartner"`
	Plant                 string  `json:"Plant"`
	ResponsiblePerson     *int    `json:"ResponsiblePerson"`
	ResponsiblePersonName *string `json:"ResponsiblePersonName"`
	PlannedStartDate      *string `json:"PlannedStartDate"`
	PlannedEndDate        *string `json:"PlannedEndDate"`
	ActualStartDate       *string `json:"ActualStartDate"`
	ActualEndDate         *string `json:"ActualEndDate"`
	CreationDate          string  `json:"CreationDate"`
	CreationTime          string  `json:"CreationTime"`
	LastChangeDate        string  `json:"LastChangeDate"`
	LastChangeTime        string  `json:"LastChangeTime"`
	IsReleased            *bool   `json:"IsReleased"`
	IsLocked              *bool   `json:"IsLocked"`
	IsCancelled           *bool   `json:"IsCancelled"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}

type WBSElement struct {
	Project               int     `json:"Project"`
	WBSElement            int     `json:"WBSElement"`
	WBSElementCode        string  `json:"WBSElementCode"`
	WBSElementDescription string  `json:"WBSElementDescription"`
	BusinessPartner       int     `json:"BusinessPartner"`
	Plant                 string  `json:"Plant"`
	ResponsiblePerson     *int    `json:"ResponsiblePerson"`
	ResponsiblePersonName *string `json:"ResponsiblePersonName"`
	PlannedStartDate      *string `json:"PlannedStartDate"`
	PlannedEndDate        *string `json:"PlannedEndDate"`
	ActualStartDate       *string `json:"ActualStartDate"`
	ActualEndDate         *string `json:"ActualEndDate"`
	CreationDate          string  `json:"CreationDate"`
	LastChangeDate        string  `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
}

type NumberRange struct {
	NumberRangeID            string `json:"NumberRangeID"`
	ServiceLabel             string `json:"ServiceLabel"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	LatestNumber             int    `json:"LatestNumber"`
}
