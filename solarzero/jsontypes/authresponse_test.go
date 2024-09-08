package jsontypes_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
)

func TestUnmarshalAuthResponseData(t *testing.T) {
	jsonData := `{
    "tokens": {
        "accessToken": "accesstokenthing",
        "idToken": "idtokenthing",
        "refreshToken": "refreshtokenthing",
        "tokenType": "Bearer",
        "expiresIn": 3600,
        "expiresAt": "2024-09-06T23:08:15.174Z",
        "sessionId": "8569376a12759125e901c621bf6be96a"
    }
}`

	// Unmarshal the JSON into the struct
	var data jsontypes.AuthResponse
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Marshal the struct back to JSON
	marshaledData, err := json.Marshal(&data)
	if err != nil {
		t.Fatalf("Error marshaling data: %v", err)
	}

	// Unmarshal the marshaled JSON back to the struct and compare with the original
	var unmarshaledData jsontypes.AuthResponse
	err = json.Unmarshal(marshaledData, &unmarshaledData)
	if err != nil {
		t.Fatalf("Error unmarshaling marshaled JSON: %v", err)
	}

	// Use reflection to compare the original and unmarshaled data
	if !reflect.DeepEqual(data, unmarshaledData) {
		t.Errorf("Original and unmarshaled data do not match. Got: %+v, Want: %+v", unmarshaledData, data)
	}
}
