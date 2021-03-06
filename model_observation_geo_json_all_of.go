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

// ObservationGeoJsonAllOf struct for ObservationGeoJsonAllOf
type ObservationGeoJsonAllOf struct {
	Properties *Observation `json:"properties,omitempty"`
}

// NewObservationGeoJsonAllOf instantiates a new ObservationGeoJsonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationGeoJsonAllOf() *ObservationGeoJsonAllOf {
	this := ObservationGeoJsonAllOf{}
	return &this
}

// NewObservationGeoJsonAllOfWithDefaults instantiates a new ObservationGeoJsonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationGeoJsonAllOfWithDefaults() *ObservationGeoJsonAllOf {
	this := ObservationGeoJsonAllOf{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ObservationGeoJsonAllOf) GetProperties() Observation {
	if o == nil || o.Properties == nil {
		var ret Observation
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationGeoJsonAllOf) GetPropertiesOk() (*Observation, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ObservationGeoJsonAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given Observation and assigns it to the Properties field.
func (o *ObservationGeoJsonAllOf) SetProperties(v Observation) {
	o.Properties = &v
}

func (o ObservationGeoJsonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableObservationGeoJsonAllOf struct {
	value *ObservationGeoJsonAllOf
	isSet bool
}

func (v NullableObservationGeoJsonAllOf) Get() *ObservationGeoJsonAllOf {
	return v.value
}

func (v *NullableObservationGeoJsonAllOf) Set(val *ObservationGeoJsonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationGeoJsonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationGeoJsonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationGeoJsonAllOf(val *ObservationGeoJsonAllOf) *NullableObservationGeoJsonAllOf {
	return &NullableObservationGeoJsonAllOf{value: val, isSet: true}
}

func (v NullableObservationGeoJsonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationGeoJsonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


