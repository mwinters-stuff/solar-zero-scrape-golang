package currentdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeCurrentData(t *testing.T) {
	json := `
	{
		"deviceStatus": 1,
		"dPowerFlow": 5,
		"export": 0,
		"import": 1026,
		"load": 1026,
		"solar": 0,
		"soc": 19,
		"charge": 0,
		"gridPowerOutage": 0,
		"temperature": 23.1
	}`

	currentdata, err := UnmarshalCurrentData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	assert.Equal(t, int64(1), currentdata.DeviceStatus)
	assert.Equal(t, int64(5), currentdata.DPowerFlow)
	assert.Equal(t, int64(0), currentdata.Export)
	assert.Equal(t, int64(1026), currentdata.Import)
	assert.Equal(t, int64(1026), currentdata.Load)
	assert.Equal(t, int64(0), currentdata.Solar)
	assert.Equal(t, int64(19), currentdata.Soc)
	assert.Equal(t, int64(0), currentdata.Charge)
	assert.Equal(t, int64(0), currentdata.GridPowerOutage)
	assert.Equal(t, 23.1, currentdata.Temperature)

}
