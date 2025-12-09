package models

type RawDataEvent struct {
	NodeIdentifier   string  `json:"nodeIdentifier"`
	SensorIdentifier string  `json:"sensorIdentifier"`
	Temperature      float64 `json:"temperature"`
}
