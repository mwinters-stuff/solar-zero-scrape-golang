// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    customerData, err := UnmarshalCustomerData(bytes)
//    bytes, err = customerData.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalCustomerData(data []byte) (CustomerData, error) {
	var r CustomerData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CustomerData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CustomerData struct {
	Name            string      `json:"name"`
	FirstName       string      `json:"firstName"`
	LastName        string      `json:"lastName"`
	Address         string      `json:"address"`
	Email           string      `json:"email"`
	ReferralCode    string      `json:"referralCode"`
	BankDirectDebit interface{} `json:"bankDirectDebit"`
	Account         Account     `json:"account"`
	Provider        Provider    `json:"provider"`
	System          System      `json:"system"`
	City            string      `json:"city"`
	HasUsageData    bool        `json:"hasUsageData"`
	IsNewAccount    bool        `json:"isNewAccount"`
	IsCarbonTrack   bool        `json:"isCarbonTrack"`
	IsEnPhase       bool        `json:"isEnPhase"`
	IsEMU           bool        `json:"isEMU"`
}

type Account struct {
	ID                string `json:"id"`
	Timezone          string `json:"timezone"`
	SiteID            string `json:"siteId"`
	HasCustomerPortal bool   `json:"hasCustomerPortal"`
	IsActivated       bool   `json:"isActivated"`
}

type Provider struct {
	ID                   string  `json:"id"`
	Plan                 string  `json:"plan"`
	IsDynamicGridPricing bool    `json:"isDynamicGridPricing"`
	Details              Details `json:"details"`
}

type Details struct {
	Weekdays Week `json:"weekdays"`
	Weekends Week `json:"weekends"`
}

type Week struct {
	Notes    string   `json:"notes"`
	Unit     string   `json:"unit"`
	OffPeak  OffPeak  `json:"offPeak"`
	Shoulder OffPeak  `json:"shoulder"`
	Peak     *OffPeak `json:"peak"`
	Hours    []Hour   `json:"hours"`
}

type Hour struct {
	Idx           int64  `json:"idx"`
	Type          string `json:"type"`
	NumberOfHours int64  `json:"numberOfHours"`
	From          int64  `json:"from"`
	To            int64  `json:"to"`
}

type OffPeak struct {
	Label    string  `json:"label"`
	RateText string  `json:"rateText"`
	Rate     float64 `json:"rate"`
}

type System struct {
	Brand string `json:"brand"`
	Type  string `json:"type"`
}
