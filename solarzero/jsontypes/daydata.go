// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dayData, err := UnmarshalDayData(bytes)
//    bytes, err = dayData.Marshal()

package jsontypes

import "encoding/json"

type DayData []DayDatum

func UnmarshalDayData(data []byte) (DayData, error) {
	var r DayData
	err := json.Unmarshal(data, &r)
	return r, err
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

func (r *DayDatum) GetInfluxFields() *map[string]interface{} {
	if r.Export != nil {
		return &map[string]interface{}{
			"Hour":        r.Hour,
			"Export":      *r.Export,
			"Grid":        *r.Grid,
			"SolarUse":    *r.SolarUse,
			"SoC":         *r.SoC,
			"Charge":      *r.Charge,
			"Discharge":   *r.Discharge,
			"Solar":       *r.Solar,
			"BatteryGrid": *r.BatteryGrid,
			"HomeLoad":    *r.HomeLoad,
		}
	} else {
		return nil
	}
}