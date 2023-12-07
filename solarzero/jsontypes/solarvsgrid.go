// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    solarVsGrid, err := UnmarshalSolarVsGrid(bytes)
//    bytes, err = solarVsGrid.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalSolarVsGrid(data []byte) (SolarVsGrid, error) {
	var r SolarVsGrid
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SolarVsGrid) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SolarVsGrid struct {
	SolarPercent int64   `json:"solarPercent"`
	SolarAmount  float64 `json:"solarAmount"`
	GridPercent  int64   `json:"gridPercent"`
	GridAmount   float64 `json:"gridAmount"`
}
