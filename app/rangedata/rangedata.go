// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    rangeExportData, err := UnmarshalRangeExportData(bytes)
//    bytes, err = rangeExportData.Marshal()

package rangedata

import "encoding/json"

type RangeExportData []RangeExportDatum

func UnmarshalRangeExportData(data []byte) (RangeExportData, error) {
	var r RangeExportData
	err := json.Unmarshal(data, &r)
	return r, err
}

type RangeExportDatum struct {
	Date      string  `json:"Date"`
	SolarUse  float64 `json:"Solar use"`
	Grid      float64 `json:"Grid"`
	Export    float64 `json:"Export"`
	Load      float64 `json:"Load"`
	Solar     float64 `json:"Solar"`
	Discharge float64 `json:"Discharge"`
	Charge    float64 `json:"Charge"`
}
