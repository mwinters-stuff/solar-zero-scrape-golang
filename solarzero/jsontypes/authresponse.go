// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    authResponse, err := UnmarshalAuthResponse(bytes)
//    bytes, err = authResponse.Marshal()

package jsontypes

import "time"

import "encoding/json"

func UnmarshalAuthResponse(data []byte) (AuthResponse, error) {
	var r AuthResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AuthResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type AuthResponse struct {
	Tokens Tokens `json:"tokens"`
}

type Tokens struct {
	AccessToken  string    `json:"accessToken"`
	IDToken      string    `json:"idToken"`
	RefreshToken string    `json:"refreshToken"`
	TokenType    string    `json:"tokenType"`
	ExpiresIn    int64     `json:"expiresIn"`
	ExpiresAt    time.Time `json:"expiresAt"`
	SessionID    string    `json:"sessionId"`
}
