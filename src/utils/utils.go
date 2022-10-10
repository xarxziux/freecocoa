package utils

import (
	"fmt"
	"freecocoa/src/models"
	"freecocoa/src/stats"
)

func ConvertDefender(data models.DefenderUnit) (*models.UnitDetails, error) {
	// answer := models.UnitDetails{}

	unit, ok := stats.UnitStats[data.Name]
	if !ok {
		return nil, fmt.Errorf("unit \"%s\" not found", data.Name)
	}

	// answer.Name= unit.Name

	return &unit, nil
}
