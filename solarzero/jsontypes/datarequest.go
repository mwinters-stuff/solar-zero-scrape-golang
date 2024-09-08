// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    infoRequestData, err := UnmarshalInfoRequestData(bytes)
//    bytes, err = infoRequestData.Marshal()

package jsontypes

import "encoding/json"

func UnmarshalInfoRequestData(data []byte) (DataRequestData, error) {
	var r DataRequestData
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DataRequestData) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DataRequestData struct {
	SiteID     string `json:"siteId"`
	Timezone   string `json:"timezone"`
	ProviderID string `json:"providerId"`
	HasTou     bool   `json:"hasTou"`
}
