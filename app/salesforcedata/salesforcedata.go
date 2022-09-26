// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    salesForceData, err := UnmarshalSalesForceData(bytes)
//    bytes, err = salesForceData.Marshal()

package salesforcedata

import "encoding/json"

func UnmarshalSalesForceData(data []byte) (SalesForceData, error) {
	var r SalesForceData
	err := json.Unmarshal(data, &r)
	return r, err
}

type SalesForceData struct {
	Contact     Contact       `json:"contact"`
	Account     Account       `json:"account"`
	Opportunity []Opportunity `json:"opportunity"`
	Token       string        `json:"token"`
	Asset       Asset         `json:"asset"`
}

type Account struct {
	AccountID  string      `json:"accountId"`
	SiteNumber string      `json:"siteNumber"`
	Activated  interface{} `json:"activated"`
	Property   Property    `json:"property"`
}

type Property struct {
	Address            string      `json:"address"`
	BuildingName       interface{} `json:"buildingName"`
	City               string      `json:"city"`
	Country            string      `json:"country"`
	Dpid               interface{} `json:"dpid"`
	ExistingPanels     interface{} `json:"existingPanels"`
	FloorLevel         interface{} `json:"floorLevel"`
	GeocodeAccuracy    interface{} `json:"geocodeAccuracy"`
	PostCode           string      `json:"postCode"`
	RDNumber           interface{} `json:"rdNumber"`
	Region             interface{} `json:"region"`
	SolarCapableRoof   interface{} `json:"solarCapableRoof"`
	StreetNameSuffix   interface{} `json:"streetNameSuffix"`
	StreetName         string      `json:"streetName"`
	StreetNumberSuffix interface{} `json:"streetNumberSuffix"`
	StreetNumber       string      `json:"streetNumber"`
	StreetType         interface{} `json:"streetType"`
	Street             string      `json:"street"`
	Suburb             string      `json:"suburb"`
	SwimmingPool       interface{} `json:"swimmingPool"`
	UnitNumber         interface{} `json:"unitNumber"`
	UnitType           interface{} `json:"unitType"`
	XCoordinate        interface{} `json:"xCoordinate"`
	YCoordinate        interface{} `json:"yCoordinate"`
	ParcelID           interface{} `json:"parcelId"`
	ValidationError    interface{} `json:"validationError"`
	ValidationStatus   interface{} `json:"validationStatus"`
	OwnershipType      string      `json:"ownershipType"`
}

type Asset struct {
	Attributes    Attributes `json:"attributes"`
	ID            string     `json:"Id"`
	SerialNumberC string     `json:"Serial_Number__c"`
	AssetR        AssetR     `json:"Asset__r"`
	BrandC        string     `json:"Brand__c"`
	TypeC         string     `json:"Type__c"`
	SubTypeC      string     `json:"Sub_Type__c"`
	StatusC       string     `json:"Status__c"`
	CreatedDate   string     `json:"CreatedDate"`
}

type AssetR struct {
	Attributes Attributes `json:"attributes"`
	AccountID  string     `json:"AccountId"`
}

type Attributes struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Contact struct {
	Attributes Attributes `json:"attributes"`
	ID         string     `json:"Id"`
	AccountID  string     `json:"AccountId"`
	Email      string     `json:"Email"`
	FirstName  string     `json:"FirstName"`
	LastName   string     `json:"LastName"`
}

type Opportunity struct {
	Attributes      Attributes      `json:"attributes"`
	ID              string          `json:"Id"`
	PrimaryContactR PrimaryContactR `json:"Primary_Contact__r"`
}

type PrimaryContactR struct {
	Attributes            Attributes `json:"attributes"`
	PrimaryContactPhoneFC string     `json:"Primary_Contact_Phone_f__c"`
}
