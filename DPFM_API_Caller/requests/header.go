package requests

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
