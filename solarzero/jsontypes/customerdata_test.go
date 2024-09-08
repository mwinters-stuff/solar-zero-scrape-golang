package jsontypes_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero/jsontypes"
)

func TestUnmarshalCustomerDataResponseData(t *testing.T) {
	jsonData := `{
    "name": "Joe Blogs",
    "firstName": "Joe",
    "lastName": "Bloggs",
    "address": "1 Some Place, Christchurch",
    "email": "example@example.com",
    "referralCode": "ABCDEF",
    "bankDirectDebit": null,
    "account": {
        "id": "123445345t4",
        "timezone": "Pacific/Auckland",
        "siteId": "SC-22-999999",
        "hasCustomerPortal": true,
        "isActivated": true
    },
    "provider": {
        "id": "nz-ecotricity",
        "plan": "NZSZ_ZER0_STD_FXTOU_240901_T1",
        "isDynamicGridPricing": true,
        "details": {
            "weekdays": {
                "notes": "Rates for per kWh of electricity used or exported. Excluding GST, lines and network charges.",
                "unit": "$",
                "offPeak": {
                    "label": "Off-peak",
                    "rateText": "0.10",
                    "rate": 0.1
                },
                "shoulder": {
                    "label": "Shoulder",
                    "rateText": "0.14",
                    "rate": 0.14
                },
                "peak": {
                    "label": "Peak",
                    "rateText": "0.26",
                    "rate": 0.26
                },
                "hours": [
                    {
                        "idx": 0,
                        "type": "offPeak",
                        "numberOfHours": 7,
                        "from": 0,
                        "to": 6
                    },
                    {
                        "idx": 7,
                        "type": "peak",
                        "numberOfHours": 2,
                        "from": 7,
                        "to": 8
                    },
                    {
                        "idx": 9,
                        "type": "shoulder",
                        "numberOfHours": 8,
                        "from": 9,
                        "to": 16
                    },
                    {
                        "idx": 17,
                        "type": "peak",
                        "numberOfHours": 4,
                        "from": 17,
                        "to": 20
                    },
                    {
                        "idx": 21,
                        "type": "shoulder",
                        "numberOfHours": 3,
                        "from": 21,
                        "to": 23
                    }
                ]
            },
            "weekends": {
                "notes": "Rates for per kWh of electricity used or exported. Excluding GST, lines and network charges.",
                "unit": "$",
                "offPeak": {
                    "label": "Off-peak",
                    "rateText": "0.10",
                    "rate": 0.1
                },
                "shoulder": {
                    "label": "Shoulder",
                    "rateText": "0.14",
                    "rate": 0.14
                },
                "peak": null,
                "hours": [
                    {
                        "idx": 0,
                        "type": "offPeak",
                        "numberOfHours": 7,
                        "from": 0,
                        "to": 6
                    },
                    {
                        "idx": 7,
                        "type": "shoulder",
                        "numberOfHours": 17,
                        "from": 7,
                        "to": 23
                    }
                ]
            }
        }
    },
    "system": {
        "brand": "panasonic",
        "type": "ICON"
    },
    "city": "Christchurch",
    "hasUsageData": true,
    "isNewAccount": false,
    "isCarbonTrack": false,
    "isEnPhase": false,
    "isEMU": false
}`

	// Unmarshal the JSON into the struct
	var data jsontypes.CustomerData
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
	var unmarshaledData jsontypes.CustomerData
	err = json.Unmarshal(marshaledData, &unmarshaledData)
	if err != nil {
		t.Fatalf("Error unmarshaling marshaled JSON: %v", err)
	}

	// Use reflection to compare the original and unmarshaled data
	if !reflect.DeepEqual(data, unmarshaledData) {
		t.Errorf("Original and unmarshaled data do not match. Got: %+v, Want: %+v", unmarshaledData, data)
	}
}
