package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header	*[]Header	`json:"Header"`
	Item	*[]Item		`json:"Item"`
}

type Header struct {
	RailwayLine					int		`json:"RailwayLine"`
	TrainOperationVersion		int		`json:"TrainOperationVersion"`
	WeekdayType					string	`json:"WeekdayType"`
	PlannedTrainOperationID		int		`json:"PlannedTrainOperationID"`
	ExpressType					string	`json:"ExpressType"`
	DepartureStation			int		`json:"DepartureStation"`
	DestinationStation			int		`json:"DestinationStation"`
	PlannedDepartureTime		string	`json:"PlannedDepartureTime"`
	PlannedArrivalTime			string	`json:"PlannedArrivalTime"`
	Description					string	`json:"Description"`
	OperationRemarks			*string	`json:"OperationRemarks"`
	OperationCode				*string	`json:"OperationCode"`
	IsSuspended					*bool	`json:"IsSuspended"`
	ValidityStartDate			string	`json:"ValidityStartDate"`
	ValidityEndDate				string	`json:"ValidityEndDate"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	CreateUser					int		`json:"CreateUser"`
	LastChangeUser				int		`json:"LastChangeUser"`
	IsReleased					*bool	`json:"IsReleased"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}

type Item struct {
	RailwayLine					int		`json:"RailwayLine"`
	TrainOperationVersion		int		`json:"TrainOperationVersion"`
	WeekdayType					string	`json:"WeekdayType"`
	PlannedTrainOperationID		int		`json:"PlannedTrainOperationID"`
	RailwayLineStationID		int		`json:"RailwayLineStationID"`
	StopStation					int		`json:"StopStation"`
	PlannedPlatformNumber		string	`json:"PlannedPlatformNumber"`
	ExpressType					string	`json:"ExpressType"`
	PlannedArrivalTime			string	`json:"PlannedArrivalTime"`
	PlannedDepartureTime		*string	`json:"PlannedDepartureTime"`
	Description					string	`json:"Description"`
	OperationRemarks			*string	`json:"OperationRemarks"`
	IsSuspended					*bool	`json:"IsSuspended"`
	CreationDate				string	`json:"CreationDate"`
	CreationTime				string	`json:"CreationTime"`
	LastChangeDate				string	`json:"LastChangeDate"`
	LastChangeTime				string	`json:"LastChangeTime"`
	CreateUser					int		`json:"CreateUser"`
	LastChangeUser				int		`json:"LastChangeUser"`
	IsReleased					*bool	`json:"IsReleased"`
	IsMarkedForDeletion			*bool	`json:"IsMarkedForDeletion"`
}
