// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cookieResult, err := UnmarshalCookieResult(bytes)
//    bytes, err = cookieResult.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalCookieResult(data []byte) (CookieResult, error) {
	var r CookieResult
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CookieResult) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CookieResult struct {
	Auth       bool   `json:"auth"`
	CustomerID string `json:"customerId"`
}
