// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    yearData, err := UnmarshalYearData(bytes)
//    bytes, err = yearData.Marshal()

package jsontypes

import (
	"encoding/json"
	"time"
)

type YearData []YearDatum

func UnmarshalYearData(data []byte) (YearData, error) {
	var r YearData
	err := json.Unmarshal(data, &r)
	return r, err
}

type YearDatum struct {
	Month       string   `json:"Month"`
	SolarUse    *float64 `json:"Solar use"`
	Grid        *float64 `json:"Grid"`
	Export      *float64 `json:"Export"`
	BatteryGrid *float64 `json:"Battery grid"`
	HomeLoad    *float64 `json:"Home load"`
}

func (r *YearDatum) GetInfluxFields() *map[string]interface{} {
	if r.SolarUse != nil {
		return &map[string]interface{}{
			"Month":       r.Month,
			"SolarUse":    *r.SolarUse,
			"Grid":        *r.Grid,
			"Export":      *r.Export,
			"BatteryGrid": *r.BatteryGrid,
			"HomeLoad":    *r.HomeLoad,
		}
	} else {
		return nil
	}
}

func (r *YearDatum) GetMonthNum() time.Month {
	switch r.Month {
	case "Jan":
		return 1
	case "Feb":
		return 2
	case "Mar":
		return 3
	case "Apr":
		return 4
	case "May":
		return 5
	case "Jun":
		return 6
	case "Jul":
		return 7
	case "Aug":
		return 8
	case "Sep":
		return 9
	case "Oct":
		return 10
	case "Nov":
		return 11
	case "Dec":
		return 12
	}
	panic("invalid month name")
}
