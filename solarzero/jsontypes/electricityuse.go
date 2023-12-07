// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    electricityUse, err := UnmarshalElectricityUse(bytes)
//    bytes, err = electricityUse.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalElectricityUse(data []byte) (ElectricityUse, error) {
	var r ElectricityUse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ElectricityUse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ElectricityUse struct {
	ElectricityUse float64 `json:"electricityUse"`
}
