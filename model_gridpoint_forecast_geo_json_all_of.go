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

// GridpointForecastGeoJsonAllOf struct for GridpointForecastGeoJsonAllOf
type GridpointForecastGeoJsonAllOf struct {
	Properties *GridpointForecast `json:"properties,omitempty"`
}

// NewGridpointForecastGeoJsonAllOf instantiates a new GridpointForecastGeoJsonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridpointForecastGeoJsonAllOf() *GridpointForecastGeoJsonAllOf {
	this := GridpointForecastGeoJsonAllOf{}
	return &this
}

// NewGridpointForecastGeoJsonAllOfWithDefaults instantiates a new GridpointForecastGeoJsonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridpointForecastGeoJsonAllOfWithDefaults() *GridpointForecastGeoJsonAllOf {
	this := GridpointForecastGeoJsonAllOf{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *GridpointForecastGeoJsonAllOf) GetProperties() GridpointForecast {
	if o == nil || o.Properties == nil {
		var ret GridpointForecast
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastGeoJsonAllOf) GetPropertiesOk() (*GridpointForecast, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *GridpointForecastGeoJsonAllOf) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given GridpointForecast and assigns it to the Properties field.
func (o *GridpointForecastGeoJsonAllOf) SetProperties(v GridpointForecast) {
	o.Properties = &v
}

func (o GridpointForecastGeoJsonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableGridpointForecastGeoJsonAllOf struct {
	value *GridpointForecastGeoJsonAllOf
	isSet bool
}

func (v NullableGridpointForecastGeoJsonAllOf) Get() *GridpointForecastGeoJsonAllOf {
	return v.value
}

func (v *NullableGridpointForecastGeoJsonAllOf) Set(val *GridpointForecastGeoJsonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGridpointForecastGeoJsonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGridpointForecastGeoJsonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridpointForecastGeoJsonAllOf(val *GridpointForecastGeoJsonAllOf) *NullableGridpointForecastGeoJsonAllOf {
	return &NullableGridpointForecastGeoJsonAllOf{value: val, isSet: true}
}

func (v NullableGridpointForecastGeoJsonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridpointForecastGeoJsonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


