// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dailyResponseData, err := UnmarshalDailyResponseData(bytes)
//    bytes, err = dailyResponseData.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalDailyResponseData(data []byte) (DailyResponseData, error) {
	var r DailyResponseData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DailyResponseData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DailyResponseData struct {
	Reports []DailyReport `json:"reports"`
}

type DailyReport struct {
	Day     string            `json:"day"`
	Home    DailyHomeClass    `json:"home"`
	Solar   DailyHomeClass    `json:"solar"`
	Battery DailyBatteryClass `json:"battery"`
	Grid    DailyGridClass    `json:"grid"`
	Tou     *DailyTou         `json:"tou"`
}

type DailyBatteryClass struct {
	HighestDischarge float64            `json:"highestDischarge"`
	LowestCharge     float64            `json:"lowestCharge"`
	Total            float64            `json:"total"`
	Total2           float64            `json:"total2"`
	Series           DailyBatterySeries `json:"series"`
	Stack            []DailyStack       `json:"stack"`
	SummaryHome      float64            `json:"summaryHome"`
	SummaryToGrid    float64            `json:"summaryToGrid"`
	SummarySolar     float64            `json:"summarySolar"`
	SummaryFromGrid  float64            `json:"summaryFromGrid"`
	Summary          []DailySummary     `json:"summary"`
	Summary2         []DailySummary     `json:"summary2"`
}

type DailyBatterySeries struct {
	Charge    []float64 `json:"charge"`
	Discharge []float64 `json:"discharge"`
}

type DailyStack struct {
	ID     ID        `json:"id"`
	Series []float64 `json:"series"`
}

type DailySummary struct {
	ID      ID      `json:"id"`
	Name    Name    `json:"name"`
	Kwh     float64 `json:"kwh"`
	Percent float64 `json:"percent"`
}

type DailyGridClass struct {
	HighestImport      float64         `json:"highestImport"`
	LowestExport       float64         `json:"lowestExport"`
	Total              float64         `json:"total"`
	Total2             float64         `json:"total2"`
	Series             DailyGridSeries `json:"series"`
	Stack              []DailyStack    `json:"stack"`
	SummaryHome        float64         `json:"summaryHome"`
	SummaryToBattery   float64         `json:"summaryToBattery"`
	SummarySolar       float64         `json:"summarySolar"`
	SummaryFromBattery float64         `json:"summaryFromBattery"`
	Summary            []DailySummary  `json:"summary"`
	Summary2           []DailySummary  `json:"summary2"`
}

type DailyGridSeries struct {
	Import []float64 `json:"import"`
	Export []float64 `json:"export"`
}

type DailyHomeClass struct {
	Total          float64        `json:"total"`
	Series         []float64      `json:"series"`
	Stack          []DailyStack   `json:"stack"`
	SummarySolar   *float64       `json:"summarySolar,omitempty"`
	SummaryBattery float64        `json:"summaryBattery"`
	SummaryGrid    float64        `json:"summaryGrid"`
	Summary        []DailySummary `json:"summary"`
	SummaryHome    *float64       `json:"summaryHome,omitempty"`
}

type DailyTou struct {
	OffPeak   float64 `json:"offPeak"`
	Peak      float64 `json:"peak"`
	Shoulder  float64 `json:"shoulder"`
	IsWeekend bool    `json:"isWeekend"`
}

type ID string

const (
	Battery     ID = "battery"
	Frombattery ID = "frombattery"
	Fromgrid    ID = "fromgrid"
	Grid        ID = "grid"
	Home        ID = "home"
	Solar       ID = "solar"
	Tobattery   ID = "tobattery"
	Togrid      ID = "togrid"
)

type Name string

const (
	NameBattery Name = "Battery"
	NameGrid    Name = "Grid"
	NameHome    Name = "Home"
	NameSolar   Name = "Solar"
)
