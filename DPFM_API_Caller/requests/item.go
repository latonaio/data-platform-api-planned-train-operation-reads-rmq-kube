package requests

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
