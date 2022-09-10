package logindata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginData(t *testing.T) {
	json := `
	{
		"auth": {
			"loggedIn": true,
			"stagingLoggedIn": true,
			"API": "https://solarzero.pnz.technology/api",
			"EmsApi": "https://ems.solarcity.panabattery.com/api",
			"firstName": "Jack",
			"userId": "SC-99-999999",
			"version": "1.0.14"
		},
		"deviceID": {
			"ID": "999999"
		},
		"modalStatus": {
			"modalStatus": false
		},
		"currentData": {
			"soc": "0",
			"solar": 0,
			"gridImport": 0,
			"gridExport": 0,
			"load": 0,
			"fetching": true,
			"error": null,
			"dPowerFlow": 1,
			"deviceStatus": "1",
			"charge": 0,
			"gridPowerOutage": 0,
			"temperature": 0
		},
		"dayData": {
			"solarUseToday": 0,
			"gridExportToday": 0,
			"gridImportToday": 0,
			"homeLoadToday": 0,
			"data": [],
			"fetching": true,
			"error": null
		},
		"monthData": {
			"solarUseMonth": 0,
			"gridExportMonth": 0,
			"gridImportMonth": 0,
			"batteryGridMonth": 0,
			"homeLoadMonth": 0,
			"data": [],
			"fetching": true,
			"error": null
		},
		"yearData": {
			"solarUseYear": 0,
			"gridExportYear": 0,
			"gridImportYear": 0,
			"batteryGridYear": 0,
			"homeLoadYear": 0,
			"data": [],
			"fetching": true,
			"error": null
		},
		"hotWaterData": {
			"hotWater": 0,
			"boost": 0
		}
	}	`

	data, err := UnmarshalLoginData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	// only interested in api and emsapi
	assert.Equal(t, "https://solarzero.pnz.technology/api", data.Auth.API)
	assert.Equal(t, "https://ems.solarcity.panabattery.com/api", data.Auth.EMSAPI)

}

//
