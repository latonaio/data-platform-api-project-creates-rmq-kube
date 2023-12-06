package requests

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
