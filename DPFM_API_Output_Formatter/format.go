package dpfm_api_output_formatter

import (
	"data-platform-api-planned-train-operation-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.RailwayLine,
			&pm.TrainOperationVersion,
			&pm.WeekdayType,
			&pm.PlannedTrainOperationID,
			&pm.ExpressType,
			&pm.DepartureStation,
			&pm.DestinationStation,
			&pm.PlannedDepartureTime,
			&pm.PlannedArrivalTime,
			&pm.Description,
			&pm.OperationRemarks,
			&pm.OperationCode,
			&pm.IsSuspended,
			&pm.ValidityStartDate,
			&pm.ValidityEndDate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			RailwayLine:					data.RailwayLine,
			TrainOperationVersion:			data.TrainOperationVersion,
			WeekdayType:					data.WeekdayType,
			PlannedTrainOperationID:		data.PlannedTrainOperationID,
			ExpressType:					data.ExpressType,
			DepartureStation:				data.DepartureStation,
			DestinationStation:				data.DestinationStation,
			PlannedDepartureTime:			data.PlannedDepartureTime,
			PlannedArrivalTime:				data.PlannedArrivalTime,
			Description:					data.Description,
			OperationRemarks:				data.OperationRemarks,
			OperationCode:					data.OperationCode,
			IsSuspended:					data.IsSuspended,
			ValidityStartDate:				data.ValidityStartDate,
			ValidityEndDate:				data.ValidityEndDate,
			CreationDate:					data.CreationDate,
			CreationTime:					data.CreationTime,
			LastChangeDate:					data.LastChangeDate,
			LastChangeTime:					data.LastChangeTime,
			CreateUser:						data.CreateUser,
			LastChangeUser:					data.LastChangeUser,
			IsReleased:						data.IsReleased,
			IsMarkedForDeletion:			data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.RailwayLine,
			&pm.TrainOperationVersion,
			&pm.WeekdayType,
			&pm.PlannedTrainOperationID,
			&pm.RailwayLineStationID,
			&pm.StopStation,
			&pm.PlannedPlatformNumber,
			&pm.ExpressType,
			&pm.PlannedArrivalTime,
			&pm.PlannedDepartureTime,
			&pm.Description,
			&pm.OperationRemarks,
			&pm.IsSuspended,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.CreateUser,
			&pm.LastChangeUser,
			&pm.IsReleased,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}

		data := pm
		item = append(item, Item{
			RailwayLine:					data.RailwayLine,
			TrainOperationVersion:			data.TrainOperationVersion,
			WeekdayType:					data.WeekdayType,
			PlannedTrainOperationID:		data.PlannedTrainOperationID,
			RailwayLineStationID:			data.RailwayLineStationID,
			StopStation:					data.StopStation,
			PlannedPlatformNumber:			data.PlannedPlatformNumber,
			ExpressType:					data.ExpressType,
			PlannedArrivalTime:				data.PlannedArrivalTime,
			PlannedDepartureTime:			data.PlannedDepartureTime,
			Description:					data.Description,
			OperationRemarks:				data.OperationRemarks,
			IsSuspended:					data.IsSuspended,
			CreationDate:					data.CreationDate,
			CreationTime:					data.CreationTime,
			LastChangeDate:					data.LastChangeDate,
			LastChangeTime:					data.LastChangeTime,
			CreateUser:						data.CreateUser,
			LastChangeUser:					data.LastChangeUser,
			IsReleased:						data.IsReleased,
			IsMarkedForDeletion:			data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}
