// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    solarUse, err := UnmarshalSolarUse(bytes)
//    bytes, err = solarUse.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalSolarUse(data []byte) (SolarUse, error) {
	var r SolarUse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SolarUse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SolarUse struct {
	SelfUsePercent int64   `json:"selfUsePercent"`
	SelfUseAmount  float64 `json:"selfUseAmount"`
	ExportPercent  int64   `json:"exportPercent"`
	ExportAmount   float64 `json:"exportAmount"`
}
