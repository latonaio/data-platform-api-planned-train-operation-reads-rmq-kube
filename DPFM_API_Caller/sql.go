package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-planned-train-operation-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-planned-train-operation-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeadersByRailwayLine":
			func() {
				header = c.HeadersByRailwayLine(mtx, input, output, errs, log)
			}()
		case "HeadersByRailwayLineWeekdayType":
			func() {
				header = c.HeadersByRailwayLineWeekdayType(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "ItemsByRailwayLine":
			func() {
				item = c.ItemsByRailwayLine(mtx, input, output, errs, log)
		case "ItemsByRailwayLineWeekdayType":
			func() {
				item = c.ItemsByRailwayLineWeekdayType(mtx, input, output, errs, log)
			}()
		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:		header,
		Item:		item,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	
	where := fmt.Sprintf("%s\nAND header.RailwayLine = %d", input.Header.RailwayLine)

	where := fmt.Sprintf("%s\nAND header.TrainOperationVersion = %d", where, input.Header.TrainOperationVersion)

	where := fmt.Sprintf("%s\nAND header.WeekdayType = \"%s\"", where, input.Header.WeekdayType)

	where := fmt.Sprintf("%s\nAND header.PlannedTrainOperationID = %d", where, input.Header.PlannedTrainOperationID)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}

	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_train_operation_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByRailwayLine(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.RailwayLine = %d", input.Header.RailwayLine)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}

	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_transaction_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.TrainOperationVersion ASC, header.WeekdayType ASC, header.PlannedTrainOperationID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByRailwayLineWeekdayType(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {

	where := fmt.Sprintf("WHERE header.RailwayLine = %d", input.Header.RailwayLine)

	where := fmt.Sprintf("%s\nAND header.WeekdayType = \"%s\"", where, input.Header.WeekdayType)

	if input.Header.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND header.IsReleased = %v", where, *input.Header.IsReleased)
	}

	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_transaction_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsReleased ASC, header.TrainOperationVersion ASC, header.WeekdayType ASC, header.PlannedTrainOperationID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	
	where := fmt.Sprintf("%s\nAND item.RailwayLine = %d", input.Item.RailwayLine)

	where := fmt.Sprintf("%s\nAND item.TrainOperationVersion = %d", where, input.Item.TrainOperationVersion)

	where := fmt.Sprintf("%s\nAND item.WeekdayType = \"%s\"", where, input.Item.WeekdayType)

	where := fmt.Sprintf("%s\nAND item.PlannedTrainOperationID = %d", where, input.Item.PlannedTrainOperationID)

	where := fmt.Sprintf("%s\nAND item.RailwayLineStationID = %d", where, input.Item.RailwayLineStationID)

	if input.Item.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND item.IsReleased = %v", where, *input.Item.IsReleased)
	}

	if input.Item.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND item.IsMarkedForDeletion = %v", where, *input.Item.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_planned_train_operation_item_data AS item
		` + where + ` ORDER BY item.IsMarkedForDeletion ASC, item.IsReleased ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemsByRailwayLine(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {

	where := fmt.Sprintf("WHERE item.RailwayLine = %d", input.Item.RailwayLine)

	if input.Item.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND item.IsReleased = %v", where, *input.Item.IsReleased)
	}

	if input.Item.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND item.IsMarkedForDeletion = %v", where, *input.Item.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_transaction_item_data AS item
		` + where + ` ORDER BY item.IsMarkedForDeletion ASC, item.IsReleased ASC, item.TrainOperationVersion ASC, item.WeekdayType ASC, item.PlannedTrainOperationID ASC, item.RailwayLineStationID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) ItemsByRailwayLineWeekdayType(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {

	where := fmt.Sprintf("WHERE item.RailwayLine = %d", input.Item.RailwayLine)

	where := fmt.Sprintf("%s\nAND item.WeekdayType = \"%s\"", where, input.Item.WeekdayType)

	if input.Item.IsReleased != nil {
		where = fmt.Sprintf("%s\nAND item.IsReleased = %v", where, *input.Item.IsReleased)
	}

	if input.Item.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND item.IsMarkedForDeletion = %v", where, *input.Item.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_point_transaction_item_data AS item
		` + where + ` ORDER BY item.IsMarkedForDeletion ASC, item.IsReleased ASC, item.TrainOperationVersion ASC, item.WeekdayType ASC, item.PlannedTrainOperationID ASC, item.RailwayLineStationID ASC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
