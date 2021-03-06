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

// IconsSizeParameterAnyOf the model 'IconsSizeParameterAnyOf'
type IconsSizeParameterAnyOf string

// List of icons_size_parameter_anyOf
const (
	SMALL IconsSizeParameterAnyOf = "small"
	MEDIUM IconsSizeParameterAnyOf = "medium"
	LARGE IconsSizeParameterAnyOf = "large"
)

// All allowed values of IconsSizeParameterAnyOf enum
var AllowedIconsSizeParameterAnyOfEnumValues = []IconsSizeParameterAnyOf{
	"small",
	"medium",
	"large",
}

func (v *IconsSizeParameterAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IconsSizeParameterAnyOf(value)
	for _, existing := range AllowedIconsSizeParameterAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IconsSizeParameterAnyOf", value)
}

// NewIconsSizeParameterAnyOfFromValue returns a pointer to a valid IconsSizeParameterAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIconsSizeParameterAnyOfFromValue(v string) (*IconsSizeParameterAnyOf, error) {
	ev := IconsSizeParameterAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IconsSizeParameterAnyOf: valid values are %v", v, AllowedIconsSizeParameterAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IconsSizeParameterAnyOf) IsValid() bool {
	for _, existing := range AllowedIconsSizeParameterAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to icons_size_parameter_anyOf value
func (v IconsSizeParameterAnyOf) Ptr() *IconsSizeParameterAnyOf {
	return &v
}

type NullableIconsSizeParameterAnyOf struct {
	value *IconsSizeParameterAnyOf
	isSet bool
}

func (v NullableIconsSizeParameterAnyOf) Get() *IconsSizeParameterAnyOf {
	return v.value
}

func (v *NullableIconsSizeParameterAnyOf) Set(val *IconsSizeParameterAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIconsSizeParameterAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIconsSizeParameterAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIconsSizeParameterAnyOf(val *IconsSizeParameterAnyOf) *NullableIconsSizeParameterAnyOf {
	return &NullableIconsSizeParameterAnyOf{value: val, isSet: true}
}

func (v NullableIconsSizeParameterAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIconsSizeParameterAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

