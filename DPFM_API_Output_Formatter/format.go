package dpfm_api_output_formatter

import (
	"data-platform-api-kanban-container-reads-rmq-kube/DPFM_API_Caller/requests"
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
			&pm.KanbanContainer,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.StorageLocation,
			&pm.StorageBin,
			&pm.KanbanControlCycle,
			&pm.Product,
			&pm.Capacity,
			&pm.CapacityUnit,
			&pm.CapacityWeight,
			&pm.CapacityWeightUnit,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			KanbanContainer:     data.KanbanContainer,
			BusinessPartner:     data.BusinessPartner,
			Plant:               data.Plant,
			StorageLocation:     data.StorageLocation,
			StorageBin:          data.StorageBin,
			KanbanControlCycle:  data.KanbanControlCycle,
			Product:             data.Product,
			Capacity:            data.Capacity,
			CapacityUnit:        data.CapacityUnit,
			CapacityWeight:      data.CapacityWeight,
			CapacityWeightUnit:  data.CapacityWeightUnit,
			CreationDate:        data.CreationDate,
			LastChangeDate:      data.LastChangeDate,
			IsMarkedForDeletion: data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}
