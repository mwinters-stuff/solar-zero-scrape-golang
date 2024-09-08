// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    infoRequestData, err := UnmarshalInfoRequestData(bytes)
//    bytes, err = infoRequestData.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalDailyRequestData(data []byte) (DailyRequestData, error) {
	var r DailyRequestData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DailyRequestData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DailyRequestData struct {
	SiteID   string `json:"siteId"`
	Timezone string `json:"timezone"`
	HasTou   bool   `json:"hasTou"`
}
