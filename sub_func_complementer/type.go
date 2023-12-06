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
	Header                         *Header                           `json:"Header"`
	HeaderDoc                      *[]HeaderDoc                      `json:"HeaderDoc"`
	Item                           *[]Item                           `json:"Item"`
	ItemComponent                  *[]ItemComponent                  `json:"ItemComponent"`
	ItemComponentStockConfirmation *[]ItemComponentStockConfirmation `json:"ItemComponentStockConfirmation"`
	ItemComponentCosting           *[]ItemComponentCosting           `json:"ItemComponentCosting"`
	ItemOperations                 *[]ItemOperations                 `json:"ItemOperations"`
}

type Header struct {
	ProductionOrder                     int      `json:"ProductionOrder"`
	ProductionOrderType                 *string  `json:"ProductionOrderType"`
	CreationDate                        string   `json:"CreationDate"`
	LastChangeDate                      string   `json:"LastChangeDate"`
	HeaderIsReleased                    *bool    `json:"HeaderIsReleased"`
	HeaderIsPartiallyConfirmed          *bool    `json:"HeaderIsPartiallyConfirmed"`
	HeaderIsConfirmed                   *bool    `json:"HeaderIsConfirmed"`
	HeaderIsLocked                      *bool    `json:"HeaderIsLocked"`
	HeaderIsMarkedForDeletion           *bool    `json:"HeaderIsMarkedForDeletion"`
	Product                             *string  `json:"Product"`
	OwnerProductionPlant                string   `json:"OwnerProductionPlant"`
	OwnerProductionPlantBusinessPartner int      `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantStorageLocation *string  `json:"OwnerProductionPlantStorageLocation"`
	MRPArea                             *string  `json:"MRPArea"`
	MRPController                       *string  `json:"MRPController"`
	ProductionSupervisor                *string  `json:"ProductionSupervisor"`
	ProductionVersion                   *int     `json:"ProductionVersion"`
	PlannedOrder                        *int     `json:"PlannedOrder"`
	OrderID                             *int     `json:"OrderID"`
	OrderItem                           *int     `json:"OrderItem"`
	ProductionOrderPlannedStartDate     *string  `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime     *string  `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate       *string  `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime       *string  `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate    *string  `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime    *string  `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate      *string  `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime      *string  `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate        *string  `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime        *string  `json:"ProductionOrderActualEndTime"`
	ProductionUnit                      *string  `json:"ProductionUnit"`
	TotalQuantity                       float32  `json:"TotalQuantity"`
	PlannedScrapQuantity                *float32 `json:"PlannedScrapQuantity"`
	ConfirmedYieldQuantity              *float32 `json:"ConfirmedYieldQuantity"`
	ProductionOrderHeaderText           *string  `json:"ProductionOrderHeaderText"`
}

type HeaderDoc struct {
	ProductionOrder          int     `json:"ProductionOrder"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type Item struct {
	ProductionOrder                       int      `json:"ProductionOrder"`
	ProductionOrderItem                   int      `json:"ProductionOrderItem"`
	CreationDate                          string   `json:"CreationDate"`
	LastChangeDate                        string   `json:"LastChangeDate"`
	ItemIsReleased                        *bool    `json:"ItemIsReleased"`
	ItemIsPartiallyConfirmed              *bool    `json:"ItemIsPartiallyConfirmed"`
	ItemIsConfirmed                       *bool    `json:"ItemIsConfirmed"`
	ItemIsLocked                          *bool    `json:"ItemIsLocked"`
	ItemIsMarkedForDeletion               *bool    `json:"ItemIsMarkedForDeletion"`
	ProductionOrderHasGeneratedOperations *bool    `json:"ProductionOrderHasGeneratedOperations"`
	ProductAvailabilityIsNotChecked       *bool    `json:"ProductAvailabilityIsNotChecked"`
	PrecedingItem                         *int     `json:"PrecedingItem"`
	FollowingItem                         *int     `json:"FollowingItem"`
	Product                               *string  `json:"Product"`
	ProductionPlant                       string   `json:"ProductionPlant"`
	ProductionPlantBusinessPartner        int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlantStorageLocation        *string  `json:"ProductionPlantStorageLocation"`
	MRPArea                               *string  `json:"MRPArea"`
	MRPController                         *string  `json:"MRPController"`
	ProductionSupervisor                  *string  `json:"ProductionSupervisor"`
	ProductionVersion                     *int     `json:"ProductionVersion"`
	PlannedOrder                          *int     `json:"PlannedOrder"`
	OrderID                               *int     `json:"OrderID"`
	OrderItem                             *int     `json:"OrderItem"`
	MinimumLotSizeQuantity                *float32 `json:"MinimumLotSizeQuantity"`
	StandardLotSizeQuantity               *float32 `json:"StandardLotSizeQuantity"`
	LotSizeRoundingQuantity               *float32 `json:"LotSizeRoundingQuantity"`
	MaximumLotSizeQuantity                *float32 `json:"MaximumLotSizeQuantity"`
	LotSizeIsFixed                        *bool    `json:"LotSizeIsFixed"`
	ProductionOrderPlannedStartDate       *string  `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime       *string  `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate         *string  `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime         *string  `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate      *string  `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime      *string  `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate        *string  `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime        *string  `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate          *string  `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime          *string  `json:"ProductionOrderActualEndTime"`
	ProductionUnit                        *string  `json:"ProductionUnit"`
	TotalQuantity                         float32  `json:"TotalQuantity"`
	PlannedScrapQuantity                  *float32 `json:"PlannedScrapQuantity"`
	ConfirmedYieldQuantity                *float32 `json:"ConfirmedYieldQuantity"`
	ProductionOrderItemText               *string  `json:"ProductionOrderItemText"`
}

type ItemComponent struct {
	ProductionOrder                      int      `json:"ProductionOrder"`
	ProductionOrderItem                  int      `json:"ProductionOrderItem"`
	Operations                           int      `json:"Operations"`
	OperationsItem                       int      `json:"OperationsItem"`
	BillOfMaterial                       int      `json:"BillOfMaterial"`
	BillOfMaterialItem                   int      `json:"BillOfMaterialItem"`
	Reservation                          *int     `json:"Reservation"`
	ReservationItem                      *int     `json:"ReservationItem"`
	ComponentProduct                     *string  `json:"ComponentProduct"`
	ComponentProductRequirementDate      *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime      *string  `json:"ComponentProductRequirementTime"`
	ComponentProductIsMarkedForBackflush *bool    `json:"ComponentProductIsMarkedForBackflush"`
	ComponentProductBusinessPartner      *int     `json:"ComponentProductBusinessPartner"`
	StockConfirmationPlant               *string  `json:"StockConfirmationPlant"`
	PlannedOrder                         *int     `json:"PlannedOrder"`
	OrderID                              *int     `json:"OrderID"`
	OrderItem                            *int     `json:"OrderItem"`
	SortField                            *string  `json:"SortField"`
	BOMItemDescription                   *string  `json:"BOMItemDescription"`
	StorageLocation                      *string  `json:"StorageLocation"`
	Batch                                *string  `json:"Batch"`
	GoodsRecipientName                   *string  `json:"GoodsRecipientName"`
	UnloadingPointName                   *string  `json:"UnloadingPointName"`
	ProductCompIsAlternativeItem         *bool    `json:"ProductCompIsAlternativeItem"`
	CostingPolicy                        *string  `json:"CostingPolicy"`
	PriceUnitQty                         *string  `json:"PriceUnitQty"`
	StandardPrice                        *float32 `json:"StandardPrice"`
	MovingAveragePrice                   *float32 `json:"MovingAveragePrice"`
	ComponentScrapInPercent              *float32 `json:"ComponentScrapInPercent"`
	OperationScrapInPercent              *float32 `json:"OperationScrapInPercent"`
	BaseUnit                             *string  `json:"BaseUnit"`
	RequiredQuantity                     *float32 `json:"RequiredQuantity"`
	WithdrawnQuantity                    *float32 `json:"WithdrawnQuantity"`
	ConfirmedAvailableQuantity           *float32 `json:"ConfirmedAvailableQuantity"`
	ProductCompOriginalQuantity          *float32 `json:"ProductCompOriginalQuantity"`
	IsMarkedForDeletion                  *bool    `json:"IsMarkedForDeletion"`
}

type ItemComponentStockConfirmation struct {
	ProductionOrder           int     `json:"ProductionOrder"`
	ProductionOrderItem       int     `json:"ProductionOrderItem"`
	Operations                int     `json:"Operations"`
	OperationsItem            int     `json:"OperationsItem"`
	BillOfMaterial            int     `json:"BillOfMaterial"`
	BillOfMaterialItem        int     `json:"BillOfMaterialItem"`
	InventoryStockType        *string `json:"InventoryStockType"`
	InventorySpecialStockType *string `json:"InventorySpecialStockType"`
	AvailableProductStock     float32 `json:"AvailableProductStock"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}

type ItemComponentCosting struct {
	ProductionOrder     int      `json:"ProductionOrder"`
	ProductionOrderItem int      `json:"ProductionOrderItem"`
	Operations          int      `json:"Operations"`
	OperationsItem      int      `json:"OperationsItem"`
	BillOfMaterial      int      `json:"BillOfMaterial"`
	BillOfMaterialItem  int      `json:"BillOfMaterialItem"`
	ComponentProduct    string   `json:"ComponentProduct"`
	Currency            *string  `json:"Currency"`
	CostingAmount       *float32 `json:"CostingAmount"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type ItemDoc struct {
	ProductionOrder          int     `json:"ProductionOrder"`
	ProductionOrderItem      int     `json:"ProductionOrderItem"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type ItemOperations struct {
	ProductionOrder                      int      `json:"ProductionOrder"`
	ProductionOrderItem                  int      `json:"ProductionOrderItem"`
	Operations                           int      `json:"Operations"`
	OperationsItem                       int      `json:"OperationsItem"`
	Sequence                             int      `json:"Sequence"`
	OperationsText                       *string  `json:"OperationsText"`
	SequenceText                         *string  `json:"SequenceText"`
	OperationIsReleased                  *bool    `json:"OperationIsReleased"`
	OperationIsPartiallyConfirmed        *bool    `json:"OperationIsPartiallyConfirmed"`
	OperationIsConfirmed                 *bool    `json:"OperationIsConfirmed"`
	OperationIsClosed                    *bool    `json:"OperationIsClosed"`
	OperationIsMarkedForDeletion         *bool    `json:"OperationIsMarkedForDeletion"`
	ProductionPlant                      *string  `json:"ProductionPlant"`
	WorkCenter                           *int     `json:"WorkCenter"`
	OperationErlstSchedldExecStrtDte     *string  `json:"OperationErlstSchedldExecStrtDte"`
	OperationErlstSchedldExecStrtTme     *string  `json:"OperationErlstSchedldExecStrtTme"`
	OperationErlstSchedldExecEndDate     *string  `json:"OperationErlstSchedldExecEndDate"`
	OperationErlstSchedldExecEndTme      *string  `json:"OperationErlstSchedldExecEndTme"`
	OperationActualExecutionStartDate    *string  `json:"OperationActualExecutionStartDate"`
	OperationActualExecutionStartTime    *string  `json:"OperationActualExecutionStartTime"`
	OperationActualExecutionEndDate      *string  `json:"OperationActualExecutionEndDate"`
	OperationActualExecutionEndTime      *string  `json:"OperationActualExecutionEndTime"`
	ErlstSchedldExecDurnInWorkdays       *string  `json:"ErlstSchedldExecDurnInWorkdays"`
	OperationActualExecutionDays         *string  `json:"OperationActualExecutionDays"`
	OperationUnit                        *string  `json:"OperationUnit"`
	OperationPlannedTotalQuantity        *float32 `json:"OperationPlannedTotalQuantity"`
	OperationTotalConfirmedYieldQuantity *float32 `json:"OperationTotalConfirmedYieldQuantity"`
}

type NumberRange struct {
	NumberRangeID            string `json:"NumberRangeID"`
	ServiceLabel             string `json:"ServiceLabel"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	LatestNumber             int    `json:"LatestNumber"`
}
