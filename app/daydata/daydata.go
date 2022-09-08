// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dayData, err := UnmarshalDayData(bytes)
//    bytes, err = dayData.Marshal()

package daydata

import "encoding/json"

type DayData []DayDatum

func UnmarshalDayData(data []byte) (DayData, error) {
	var r DayData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DayData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DayDatum struct {
	Hour        string   `json:"Hour"`
	Export      *float64 `json:"Export"`
	Grid        *float64 `json:"Grid"`
	SolarUse    *float64 `json:"Solar use"`
	SoC         *float64 `json:"SoC"`
	Charge      *float64 `json:"Charge"`
	Discharge   *float64 `json:"Discharge"`
	Solar       *float64 `json:"Solar"`
	BatteryGrid *float64 `json:"Battery grid"`
	HomeLoad    *float64 `json:"Home load"`
}

func (r *DayDatum) GetInfluxFields() map[string]interface{} {
	return map[string]interface{}{
		"hour":        r.Hour,
		"export":      r.Export,
		"grid":        r.Grid,
		"solarUse":    r.SolarUse,
		"soc":         r.SoC,
		"charge":      r.Charge,
		"discharge":   r.Discharge,
		"solar":       r.Solar,
		"batteryGrid": r.BatteryGrid,
		"homeLoad":    r.HomeLoad,
	}
}
