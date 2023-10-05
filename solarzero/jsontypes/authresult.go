// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    authResult, err := UnmarshalAuthResult(bytes)
//    bytes, err = authResult.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalAuthResult(data []byte) (AuthResult, error) {
	var r AuthResult
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthResult) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AuthResult struct {
	URL         string `json:"url"`
	TokenString string `json:"tokenString"`
}
