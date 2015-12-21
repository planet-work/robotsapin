package sapi

type SensorsStatus struct {
	Temperature int               `json:"temperature"`
	Pressure    int               `json:"pressure"`
	Humidity    int               `json:"humidity"`
	Orientation SensorOrientation `json:"orientation"`
	Updated     int               `json:"updated"`
}

type SensorOrientation struct {
	Pitch float64 `json:"pitch"`
	Yaw   float64 `json:"yaw"`
	Roll  float64 `json:"roll"`
}

func Sensors() (*SensorsStatus, error) {
	var sensors *SensorsStatus
	sensors.Temperature = 19
	sensors.Pressure = 1012
	sensors.Humidity = 40
	sensors.Updated = 0
	return sensors, nil
}
