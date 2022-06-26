/*
weather.gov API

weather.gov API

API version: 1.8.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package weatherApi

import (
	"encoding/json"
	"fmt"
)

// LandRegionCode Land region code. These correspond to the six NWS regional headquarters: * AR: Alaska Region * CR: Central Region * ER: Eastern Region * PR: Pacific Region * SR: Southern Region * WR: Western Region 
type LandRegionCode string

// List of LandRegionCode
const (
	AR LandRegionCode = "AR"
	CR LandRegionCode = "CR"
	ER LandRegionCode = "ER"
	PR LandRegionCode = "PR"
	SR LandRegionCode = "SR"
	WR LandRegionCode = "WR"
)

// All allowed values of LandRegionCode enum
var AllowedLandRegionCodeEnumValues = []LandRegionCode{
	"AR",
	"CR",
	"ER",
	"PR",
	"SR",
	"WR",
}

func (v *LandRegionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LandRegionCode(value)
	for _, existing := range AllowedLandRegionCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LandRegionCode", value)
}

// NewLandRegionCodeFromValue returns a pointer to a valid LandRegionCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLandRegionCodeFromValue(v string) (*LandRegionCode, error) {
	ev := LandRegionCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LandRegionCode: valid values are %v", v, AllowedLandRegionCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LandRegionCode) IsValid() bool {
	for _, existing := range AllowedLandRegionCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LandRegionCode value
func (v LandRegionCode) Ptr() *LandRegionCode {
	return &v
}

type NullableLandRegionCode struct {
	value *LandRegionCode
	isSet bool
}

func (v NullableLandRegionCode) Get() *LandRegionCode {
	return v.value
}

func (v *NullableLandRegionCode) Set(val *LandRegionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableLandRegionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableLandRegionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLandRegionCode(val *LandRegionCode) *NullableLandRegionCode {
	return &NullableLandRegionCode{value: val, isSet: true}
}

func (v NullableLandRegionCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLandRegionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
