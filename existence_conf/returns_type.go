package existence_conf

type Returns struct {
	ConnectionKey         string                `json:"connection_key"`
	Result                bool                  `json:"result"`
	RedisKey              string                `json:"redis_key"`
	RuntimeSessionID      string                `json:"runtime_session_id"`
	BusinessPartnerID     *int                  `json:"business_partner"`
	Filepath              string                `json:"filepath"`
	ServiceLabel          string                `json:"service_label"`
	ProductMasterReturn   ProductMasterReturn   `json:"ProductMaster"`
	PlannedOrderReturn    PlannedOrderReturn    `json:"PlannedOrder"`
	BPGeneralReturn       BPGeneralReturn       `json:"BusinessPartnerGeneral"`
	PlantGeneralReturn    PlantGeneralReturn    `json:"PlantGeneral"`
	CurrencyReturn        CurrencyReturn        `json:"Currency"`
	BatchReturn           BatchReturn           `json:"Batch"`
	StorageLocationReturn StorageLocationReturn `json:"StorageLocation"`
	ReservationReturn     ReservationReturn     `json:"Reservation"`
	ReservationItemReturn ReservationItemReturn `json:"ReservationItem"`
	APISchema             string                `json:"api_schema"`
	Accepter              []string              `json:"accepter"`
	Deleted               bool                  `json:"deleted"`
}

type ProductMasterReturn struct {
	General struct {
		Product       string `json:"Product"`
		ExistenceConf bool   `json:"ExistenceConf"`
	} `json:"General"`
}

type PlannedOrderReturn struct {
	PlannedOrder int `json:"PlannedOrder"`
}

type BPGeneralReturn struct {
	BusinessPartner int `json:"BusinessPartner"`
}

type PlantGeneralReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type CurrencyReturn struct {
	Currency string `json:"Currency"`
}

type BatchReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Product         string `json:"Product"`
	Plant           string `json:"Plant"`
	Batch           string `json:"Batch"`
}

type StorageLocationReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
}

type ReservationReturn struct {
	Reservation int `json:"Reservation"`
}

type ReservationItemReturn struct {
	Reservation     int `json:"Reservation"`
	ReservationItem int `json:"ReservationItem"`
}
