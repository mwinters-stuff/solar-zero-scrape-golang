// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dayData, err := UnmarshalDayData(bytes)
//    bytes, err = dayData.Marshal()

package jsontypes

import "encoding/json"

type DayData []DayDatum

func UnmarshalDayData(data []byte) (DayData, error) {
	var r DayData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DayData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DayDatum struct {
	ReceivedDate  string  `json:"receivedDate"`
	SolarUse      float64 `json:"solarUse"`
	Grid          float64 `json:"grid"`
	Export        float64 `json:"export"`
	StateOfCharge int64   `json:"stateOfCharge"`
}
