/*
weather.gov API

weather.gov API

API version: 1.8.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package weatherApi

import (
	"encoding/json"
)

// PointGeoJsonAllOf struct for PointGeoJsonAllOf
type PointGeoJsonAllOf struct {
	Properties *Point `json:"properties,omitempty"`
}

// NewPointGeoJsonAllOf instantiates a new PointGeoJsonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointGeoJsonAllOf() *PointGeoJsonAllOf {
	this := PointGeoJsonAllOf{}
	return &this
}

// NewPointGeoJsonAllOfWithDefaults instantiates a new PointGeoJsonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointGeoJsonAllOfWithDefaults() *PointGeoJsonAllOf {
	this := PointGeoJsonAllOf{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *PointGeoJsonAllOf) GetProperties() Point {
	if o == nil || o.Properties == nil {
		var ret Point
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PointGeoJsonAllOf) GetPropertiesOk() (*Point, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *PointGeoJsonAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given Point and assigns it to the Properties field.
func (o *PointGeoJsonAllOf) SetProperties(v Point) {
	o.Properties = &v
}

func (o PointGeoJsonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullablePointGeoJsonAllOf struct {
	value *PointGeoJsonAllOf
	isSet bool
}

func (v NullablePointGeoJsonAllOf) Get() *PointGeoJsonAllOf {
	return v.value
}

func (v *NullablePointGeoJsonAllOf) Set(val *PointGeoJsonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePointGeoJsonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePointGeoJsonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointGeoJsonAllOf(val *PointGeoJsonAllOf) *NullablePointGeoJsonAllOf {
	return &NullablePointGeoJsonAllOf{value: val, isSet: true}
}

func (v NullablePointGeoJsonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointGeoJsonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


