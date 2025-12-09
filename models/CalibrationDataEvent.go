package models

import "github.com/gofrs/uuid/v5"

type CalibrationDataEvent struct {
	NodeIdentifier   string    `json:"nodeIdentifier"`
	NodeID           uuid.UUID `json:"nodeId"`
	SensorIdentifier string    `json:"sensorIdentifier"`
	SensorID         uuid.UUID `json:"sensorId"`
	RawDataID        uuid.UUID `json:"rawDataId"`
	RawTemperature   float64   `json:"rawTemperature"`
}
