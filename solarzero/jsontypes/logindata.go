// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    LoginData, err := UnmarshalLoginData(bytes)
//    bytes, err = LoginData.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalLoginData(data []byte) (LoginData, error) {
	var r LoginData
	err := json.Unmarshal(data, &r)
	return r, err
}

type LoginData struct {
	Auth         LoginAuth         `json:"auth"`
	DeviceID     LoginDeviceID     `json:"deviceID"`
	ModalStatus  LoginModalStatus  `json:"modalStatus"`
	CurrentData  LoginCurrentData  `json:"currentData"`
	DayData      LoginDayData      `json:"dayData"`
	MonthData    LoginMonthData    `json:"monthData"`
	YearData     LoginYearData     `json:"yearData"`
	HotWaterData LoginHotWaterData `json:"hotWaterData"`
}

type LoginAuth struct {
	LoggedIn        bool   `json:"loggedIn"`
	StagingLoggedIn bool   `json:"stagingLoggedIn"`
	API             string `json:"API"`
	EMSAPI          string `json:"EmsApi"`
	FirstName       string `json:"firstName"`
	UserID          string `json:"userId"`
	Version         string `json:"version"`
}

type LoginCurrentData struct {
	Soc             string      `json:"soc"`
	Solar           int64       `json:"solar"`
	GridImport      int64       `json:"gridImport"`
	GridExport      int64       `json:"gridExport"`
	Load            int64       `json:"load"`
	Fetching        bool        `json:"fetching"`
	Error           interface{} `json:"error"`
	DPowerFlow      int64       `json:"dPowerFlow"`
	DeviceStatus    string      `json:"deviceStatus"`
	Charge          int64       `json:"charge"`
	GridPowerOutage int64       `json:"gridPowerOutage"`
	Temperature     int64       `json:"temperature"`
}

type LoginDayData struct {
	SolarUseToday   int64         `json:"solarUseToday"`
	GridExportToday int64         `json:"gridExportToday"`
	GridImportToday int64         `json:"gridImportToday"`
	HomeLoadToday   int64         `json:"homeLoadToday"`
	Data            []interface{} `json:"data"`
	Fetching        bool          `json:"fetching"`
	Error           interface{}   `json:"error"`
}

type LoginDeviceID struct {
	ID string `json:"ID"`
}

type LoginHotWaterData struct {
	HotWater int64 `json:"hotWater"`
	Boost    int64 `json:"boost"`
}

type LoginModalStatus struct {
	ModalStatus bool `json:"modalStatus"`
}

type LoginMonthData struct {
	SolarUseMonth    int64         `json:"solarUseMonth"`
	GridExportMonth  int64         `json:"gridExportMonth"`
	GridImportMonth  int64         `json:"gridImportMonth"`
	BatteryGridMonth int64         `json:"batteryGridMonth"`
	HomeLoadMonth    int64         `json:"homeLoadMonth"`
	Data             []interface{} `json:"data"`
	Fetching         bool          `json:"fetching"`
	Error            interface{}   `json:"error"`
}

type LoginYearData struct {
	SolarUseYear    int64         `json:"solarUseYear"`
	GridExportYear  int64         `json:"gridExportYear"`
	GridImportYear  int64         `json:"gridImportYear"`
	BatteryGridYear int64         `json:"batteryGridYear"`
	HomeLoadYear    int64         `json:"homeLoadYear"`
	Data            []interface{} `json:"data"`
	Fetching        bool          `json:"fetching"`
	Error           interface{}   `json:"error"`
}
