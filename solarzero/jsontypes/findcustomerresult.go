// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    findCustomerResult, err := UnmarshalFindCustomerResult(bytes)
//    bytes, err = findCustomerResult.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalFindCustomerResult(data []byte) (FindCustomerResult, error) {
	var r FindCustomerResult
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *FindCustomerResult) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type FindCustomerResult struct {
	ID         string      `json:"id"`
	CustomerID string      `json:"customerId"`
	Theme      string      `json:"theme"`
	Solar      int64       `json:"solar"`
	Battery    int64       `json:"battery"`
	Evcharger  interface{} `json:"evcharger"`
	HotWater   int64       `json:"hotWater"`
	Aircon     interface{} `json:"aircon"`
}
