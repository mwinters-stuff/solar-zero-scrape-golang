// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dataResponseData, err := UnmarshalDataResponseData(bytes)
//    bytes, err = dataResponseData.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalDataResponseData(data []byte) (DataResponseData, error) {
	var r DataResponseData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DataResponseData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DataResponseData struct {
	Cards      DataCards      `json:"cards"`
	Monitor    DataMonitor    `json:"monitor"`
	Hotwater   DataHotwater   `json:"hotwater"`
	EnergyFlow DataEnergyFlow `json:"energyFlow"`
	Tou        DataTou        `json:"tou"`
}

type DataCards struct {
	HomeUsage        DataGridValue       `json:"homeUsage"`
	SolarUtilization DataGridValue       `json:"solarUtilization"`
	HomeUsageTotal   DataGridExportTotal `json:"homeUsageTotal"`
	SolarUtilTotal   DataGridExportTotal `json:"solarUtilTotal"`
	GridImportTotal  DataGridExportTotal `json:"gridImportTotal"`
	GridExportTotal  DataGridExportTotal `json:"gridExportTotal"`
}

type DataGridExportTotal struct {
	Value float64 `json:"value"`
}

type DataEnergyFlow struct {
	LastUpdate     int64         `json:"lastUpdate"`
	Operation      DataOperation `json:"operation"`
	Home           float64       `json:"home"`
	Solar          float64       `json:"solar"`
	Grid           float64       `json:"grid"`
	Battery        float64       `json:"battery"`
	GridImport     bool          `json:"gridImport"`
	GridExport     bool          `json:"gridExport"`
	BatteryUsed    bool          `json:"batteryUsed"`
	BatteryCharged bool          `json:"batteryCharged"`
	Flows          DataFlows     `json:"flows"`
}

type DataFlows struct {
	Threshold      int64   `json:"threshold"`
	Solartohome    float64 `json:"solartohome"`
	Solartobattery float64 `json:"solartobattery"`
	Solartogrid    float64 `json:"solartogrid"`
	Gridtohome     float64 `json:"gridtohome"`
	Batterytohome  float64 `json:"batterytohome"`
	Batterytogrid  float64 `json:"batterytogrid"`
	Gridtobattery  float64 `json:"gridtobattery"`
}

type DataOperation struct {
	Mode        string `json:"mode"`
	Comment     string `json:"comment"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type DataHotwater struct {
	State              string              `json:"state"`
	EnergySavingStatus string              `json:"energySavingStatus"`
	StateTitle         string              `json:"stateTitle"`
	Voltage            float64             `json:"voltage"`
	Current            float64             `json:"current"`
	Wattage            float64             `json:"wattage"`
	Plan               string              `json:"plan"`
	HourlyUsage        []float64           `json:"hourlyUsage"`
	IsAvailable        bool                `json:"isAvailable"`
	Total              float64             `json:"total"`
	AverageDailyUsage  float64             `json:"averageDailyUsage"`
	Comment            string              `json:"comment"`
	MorningPeakFrom    string              `json:"morningPeakFrom"`
	MorningPeakTo      string              `json:"morningPeakTo"`
	EveningPeakFrom    string              `json:"eveningPeakFrom"`
	EveningPeakTo      string              `json:"eveningPeakTo"`
	CircuitOn          bool                `json:"circuitOn"`
	HeatingOn          bool                `json:"heatingOn"`
	Troubleshooting    DataTroubleshooting `json:"troubleshooting"`
}

type DataTroubleshooting struct {
	Title    string       `json:"title"`
	Reminder DataReminder `json:"reminder"`
	Faqs     []DataFAQ    `json:"faqs"`
}

type DataFAQ struct {
	ID       string `json:"id"`
	Question string `json:"question"`
	Type     string `json:"type"`
}

type DataReminder struct {
	Title                 string     `json:"title"`
	Note                  string     `json:"note"`
	Description           string     `json:"description"`
	AcknowledgementAction DataAction `json:"acknowledgementAction"`
	DeferAction           DataAction `json:"deferAction"`
}

type DataAction struct {
	Title string `json:"title"`
}

type DataMonitor struct {
	Home    DataHome           `json:"home"`
	Solar   DataHome           `json:"solar"`
	Battery DataMonitorBattery `json:"battery"`
	Carbon  DataCarbon         `json:"carbon"`
}

type DataMonitorBattery struct {
	Capacity float64   `json:"capacity"`
	Charged  float64   `json:"charged"`
	Series   []float64 `json:"series"`
}

type DataCarbon struct {
	Value    float64 `json:"value"`
	Desc     string  `json:"desc"`
	Comments string  `json:"comments"`
}

type DataGridValue struct {
	Value int64 `json:"value"`
}

type DataHome struct {
	Comments string        `json:"comments"`
	Value1   DataGridValue `json:"value1"`
	Value2   DataGridValue `json:"value2"`
}

type DataTou struct {
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	MoreDescription string        `json:"moreDescription"`
	Grid            DataGridClass `json:"grid"`
	Battery         DataGridClass `json:"battery"`
	Distribution    bool          `json:"distribution"`
	OffPeak         float64       `json:"offPeak"`
	Shoulder        float64       `json:"shoulder"`
	Peak            float64       `json:"peak"`
}

type DataGridClass struct {
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	State    string `json:"state"`
}
