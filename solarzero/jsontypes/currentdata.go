// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    currentData, err := UnmarshalCurrentData(bytes)
//    bytes, err = currentData.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalCurrentData(data []byte) (CurrentData, error) {
	var r CurrentData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CurrentData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CurrentData struct {
	Ppv1                 float64     `json:"ppv1"`
	Ppv2                 float64     `json:"ppv2"`
	ReceivedDate         string      `json:"receivedDate"`
	Soc                  int64       `json:"soc"`
	Load                 float64     `json:"load"`
	DeviceStatus         int64       `json:"deviceStatus"`
	Temperature          float64     `json:"temperature"`
	Import               float64     `json:"import"`
	Export               float64     `json:"export"`
	BatteryCurrent       float64     `json:"batteryCurrent"`
	BatteryVoltage       float64     `json:"batteryVoltage"`
	Charge               float64     `json:"charge"`
	Discharge            float64     `json:"discharge"`
	PowerFlow            int64       `json:"powerFlow"`
	GridPowerMode        int64       `json:"gridPowerMode"`
	GridPowerOutage      int64       `json:"gridPowerOutage"`
	InterconnectionState interface{} `json:"interconnectionState"`
	TotalCapacity        int64       `json:"totalCapacity"`
}
