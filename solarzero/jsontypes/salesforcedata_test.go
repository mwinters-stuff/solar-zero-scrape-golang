package jsontypes_test

import (
	"testing"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/app/jsontypes"
	"github.com/stretchr/testify/assert"
)

func TestDecodeSalesForceData(t *testing.T) {
	json := `
	{
		"contact": {
			"attributes": {
				"type": "Contact",
				"url": "/services/data/v42.0/sobjects/Contact/0036XXXXXX"
			},
			"Id": "0036F00003nvYX0QAM",
			"AccountId": "0016F00003umEDSQA2",
			"Email": "some@email",
			"FirstName": "Some",
			"LastName": "Person"
		},
		"account": {
			"accountId": "00XXXXXXXXXXXX",
			"siteNumber": "SC-11-111111",
			"activated": null,
			"property": {
				"address": "1 Some Place, Brooks, Christchurch,  8083, New Zealand",
				"buildingName": null,
				"city": "Christchurch",
				"country": "New Zealand",
				"dpid": null,
				"existingPanels": null,
				"floorLevel": null,
				"geocodeAccuracy": null,
				"postCode": "8083",
				"rdNumber": null,
				"region": null,
				"solarCapableRoof": null,
				"streetNameSuffix": null,
				"streetName": "Some Place",
				"streetNumberSuffix": null,
				"streetNumber": "1",
				"streetType": null,
				"street": "1 Some Place",
				"suburb": "Brooks",
				"swimmingPool": null,
				"unitNumber": null,
				"unitType": null,
				"xCoordinate": null,
				"yCoordinate": null,
				"parcelId": null,
				"validationError": null,
				"validationStatus": null,
				"ownershipType": "Multiple Owners"
			}
		},
		"opportunity": [
			{
				"attributes": {
					"type": "Opportunity",
					"url": "/services/data/v42.0/sobjects/Opportunity/0066FXXXXXX"
				},
				"Id": "0066FXXXXXX",
				"Primary_Contact__r": {
					"attributes": {
						"type": "Contact",
						"url": "/services/data/v42.0/sobjects/Contact/0036FXXXXXX"
					},
					"Primary_Contact_Phone_f__c": "555566667777"
				}
			},
			{
				"attributes": {
					"type": "Opportunity",
					"url": "/services/data/v42.0/sobjects/Opportunity/0066FXXXXXX"
				},
				"Id": "0066XXXXXX",
				"Primary_Contact__r": {
					"attributes": {
						"type": "Contact",
						"url": "/services/data/v42.0/sobjects/Contact/0036FXXXXX"
					},
					"Primary_Contact_Phone_f__c": "555666777"
				}
			}
		],
		"token": "some token",
		"asset": {
			"attributes": {
				"type": "Component__c",
				"url": "/services/data/v42.0/sobjects/Component__c/a2C6FXXXXXX"
			},
			"Id": "a2C6FXXXXXX",
			"Serial_Number__c": "XXXXXX",
			"Asset__r": {
				"attributes": {
					"type": "Asset",
					"url": "/services/data/v42.0/sobjects/Asset/02i6F00000JZK6ZQAX"
				},
				"AccountId": "00XXXXXXX"
			},
			"Brand__c": "Panasonic",
			"Type__c": "Monitoring Asset",
			"Sub_Type__c": "ICON",
			"Status__c": "Deployed",
			"CreatedDate": "2022-04-07T22:10:40.000+0000"
		}
	}
	`

	data, err := jsontypes.UnmarshalSalesForceData([]byte(json))
	assert.Nil(t, err, "Err is not nil")

	// only interested in token
	assert.Equal(t, "some token", data.Token)

}

//
