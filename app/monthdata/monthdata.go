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

func (r *MonthData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MonthDatum struct {
	Day         int64    `json:"Day"`
	SolarUse    *float64 `json:"Solar use"`
	Grid        float64  `json:"Grid"`
	Export      float64  `json:"Export"`
	Solar       *float64 `json:"Solar"`
	BatteryGrid *float64 `json:"Battery grid"`
	HomeLoad    *float64 `json:"Home load"`
}
