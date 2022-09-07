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

func (r *CurrentData) Marshal() ([]byte, error) {
	return json.Marshal(r)
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
		"deviceStatus":    r.DeviceStatus,
		"dPowerFlow":      r.DPowerFlow,
		"export":          r.Export,
		"import":          r.Import,
		"Load":            r.Load,
		"solar":           r.Solar,
		"soc":             r.Soc,
		"charge":          r.Charge,
		"gridPowerOutage": r.GridPowerOutage,
		"temperature":     r.Temperature,
	}
}

func (r *CurrentData) Equals(o *CurrentData) bool {
	return r.DeviceStatus == o.DeviceStatus &&
		r.DPowerFlow == o.DPowerFlow &&
		r.Export == o.Export &&
		r.Import == o.Import &&
		r.Load == o.Load &&
		r.Solar == o.Solar &&
		r.Soc == o.Soc &&
		r.Charge == o.Charge &&
		r.GridPowerOutage == o.GridPowerOutage &&
		r.Temperature == o.Temperature

}
