package sapi

type SensorsStatus struct {
	Temperature int `json:"temperature"`
	Pressure    int `json:"pressure"`
	Humidity    int `json:"humidity"`
}
