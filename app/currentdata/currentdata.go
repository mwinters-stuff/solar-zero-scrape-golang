// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    CurrentData, err := UnmarshalCurrentData(bytes)
//    bytes, err = CurrentData.Marshal()

package currentdata

import "encoding/json"

func UnmarshalCurrentData(data []byte) (CurrentData, error) {
	var r CurrentData
	err := json.Unmarshal(data, &r)
	return r, err
}

type CurrentData struct {
	DeviceStatus    int64   `json:"deviceStatus"`
	DPowerFlow      int64   `json:"dPowerFlow"`
	Export          int64   `json:"export"`
	Import          int64   `json:"import"`
	Load            int64   `json:"load"`
	Solar           int64   `json:"solar"`
	Soc             int64   `json:"soc"`
	Charge          int64   `json:"charge"`
	GridPowerOutage int64   `json:"gridPowerOutage"`
	Temperature     float64 `json:"temperature"`
}

func (r *CurrentData) GetInfluxFields() map[string]interface{} {
	return map[string]interface{}{
		"DeviceStatus":    r.DeviceStatus,
		"DPowerFlow":      r.DPowerFlow,
		"Export":          r.Export,
		"Import":          r.Import,
		"Load":            r.Load,
		"Solar":           r.Solar,
		"Soc":             r.Soc,
		"Charge":          r.Charge,
		"GridPowerOutage": r.GridPowerOutage,
		"Temperature":     r.Temperature,
	}
}
