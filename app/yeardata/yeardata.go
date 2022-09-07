// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    yearData, err := UnmarshalYearData(bytes)
//    bytes, err = yearData.Marshal()

package yeardata

import "encoding/json"

type YearData []YearDatum

func UnmarshalYearData(data []byte) (YearData, error) {
	var r YearData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *YearData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type YearDatum struct {
	Month       string   `json:"Month"`
	SolarUse    *float64 `json:"Solar use"`
	Grid        float64  `json:"Grid"`
	Export      float64  `json:"Export"`
	BatteryGrid *float64 `json:"Battery grid"`
	HomeLoad    *float64 `json:"Home load"`
}
