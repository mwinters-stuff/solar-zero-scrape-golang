package jsontypes_test

import (
	"encoding/json"
	"testing"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
)

func TestUnmarshalInfoRequestData(t *testing.T) {
	// Example JSON input
	jsonData := `{
		"siteId": "123",
		"timezone": "PST",
		"providerId": "abc-provider",
		"hasTou": true
	}`

	// Expected struct based on the JSON input
	expected := jsontypes.DataRequestData{
		SiteID:     "123",
		Timezone:   "PST",
		ProviderID: "abc-provider",
		HasTou:     true,
	}

	// Test Unmarshal function
	result, err := jsontypes.UnmarshalInfoRequestData([]byte(jsonData))
	if err != nil {
		t.Errorf("UnmarshalInfoRequestData failed: %v", err)
	}

	// Check if the result matches the expected output
	if result != expected {
		t.Errorf("UnmarshalInfoRequestData produced incorrect result. Expected: %v, Got: %v", expected, result)
	}
}

func TestMarshalInfoRequestData(t *testing.T) {
	// Struct to marshal
	data := jsontypes.DataRequestData{
		SiteID:     "123",
		Timezone:   "PST",
		ProviderID: "abc-provider",
		HasTou:     true,
	}

	// Expected JSON output
	expectedJSON := `{"siteId":"123","timezone":"PST","providerId":"abc-provider","hasTou":true}`

	// Test Marshal function
	marshaled, err := data.Marshal()
	if err != nil {
		t.Errorf("Marshal failed: %v", err)
	}

	// Check if the marshaled output matches the expected JSON
	var marshaledMap map[string]interface{}
	var expectedMap map[string]interface{}
	json.Unmarshal(marshaled, &marshaledMap)
	json.Unmarshal([]byte(expectedJSON), &expectedMap)

	if !equalMaps(marshaledMap, expectedMap) {
		t.Errorf("Marshal produced incorrect result. Expected: %s, Got: %s", expectedJSON, string(marshaled))
	}
}

// Helper function to compare two maps
func equalMaps(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for key, valA := range a {
		valB, ok := b[key]
		if !ok || valA != valB {
			return false
		}
	}
	return true
}

func TestDataRequestData(t *testing.T) {
	data := jsontypes.DataRequestData{
		SiteID:     "test-site",
		Timezone:   "UTC",
		ProviderID: "test-provider",
		HasTou:     true,
	}

	jsonStr, err := json.Marshal(data)
	if err != nil {
		t.Errorf("Error while marshaling DataRequestData to JSON: %v", err)
	}

	expectedJsonStr := `{"siteId":"test-site","timezone":"UTC","providerId":"test-provider","hasTou":true}`
	if string(jsonStr) != expectedJsonStr {
		t.Errorf("Expected JSON string %s, but got %s", expectedJsonStr, jsonStr)
	}

	var decodedData jsontypes.DataRequestData
	err = json.Unmarshal([]byte(expectedJsonStr), &decodedData)
	if err != nil {
		t.Errorf("Error while unmarshaling JSON to DataRequestData: %v", err)
	}

	if decodedData != data {
		t.Errorf("Expected decoded data to be equal to the original data")
	}
}
