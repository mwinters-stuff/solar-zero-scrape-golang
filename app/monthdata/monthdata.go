// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    monthData, err := UnmarshalMonthData(bytes)
//    bytes, err = monthData.Marshal()

package monthdata

import "encoding/json"

type MonthData []MonthDatum

func UnmarshalMonthData(data []byte) (MonthData, error) {
	var r MonthData
	err := json.Unmarshal(data, &r)
	return r, err
}

type MonthDatum struct {
	Day         int64    `json:"Day"`
	SolarUse    *float64 `json:"Solar use"`
	Grid        *float64 `json:"Grid"`
	Export      *float64 `json:"Export"`
	Solar       *float64 `json:"Solar"`
	BatteryGrid *float64 `json:"Battery grid"`
	HomeLoad    *float64 `json:"Home load"`
}

func (r *MonthDatum) GetInfluxFields() *map[string]interface{} {
	if r.SolarUse != nil {
		return &map[string]interface{}{
			"Day":         r.Day,
			"SolarUse":    *r.SolarUse,
			"Grid":        *r.Grid,
			"Export":      *r.Export,
			"Solar":       *r.Solar,
			"BatteryGrid": *r.BatteryGrid,
			"HomeLoad":    *r.HomeLoad,
		}
	} else {
		return nil
	}

}
