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

// AlertCollectionGeoJsonAllOfFeatures struct for AlertCollectionGeoJsonAllOfFeatures
type AlertCollectionGeoJsonAllOfFeatures struct {
	Properties *Alert `json:"properties,omitempty"`
}

// NewAlertCollectionGeoJsonAllOfFeatures instantiates a new AlertCollectionGeoJsonAllOfFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertCollectionGeoJsonAllOfFeatures() *AlertCollectionGeoJsonAllOfFeatures {
	this := AlertCollectionGeoJsonAllOfFeatures{}
	return &this
}

// NewAlertCollectionGeoJsonAllOfFeaturesWithDefaults instantiates a new AlertCollectionGeoJsonAllOfFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertCollectionGeoJsonAllOfFeaturesWithDefaults() *AlertCollectionGeoJsonAllOfFeatures {
	this := AlertCollectionGeoJsonAllOfFeatures{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AlertCollectionGeoJsonAllOfFeatures) GetProperties() Alert {
	if o == nil || o.Properties == nil {
		var ret Alert
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCollectionGeoJsonAllOfFeatures) GetPropertiesOk() (*Alert, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AlertCollectionGeoJsonAllOfFeatures) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given Alert and assigns it to the Properties field.
func (o *AlertCollectionGeoJsonAllOfFeatures) SetProperties(v Alert) {
	o.Properties = &v
}

func (o AlertCollectionGeoJsonAllOfFeatures) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableAlertCollectionGeoJsonAllOfFeatures struct {
	value *AlertCollectionGeoJsonAllOfFeatures
	isSet bool
}

func (v NullableAlertCollectionGeoJsonAllOfFeatures) Get() *AlertCollectionGeoJsonAllOfFeatures {
	return v.value
}

func (v *NullableAlertCollectionGeoJsonAllOfFeatures) Set(val *AlertCollectionGeoJsonAllOfFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertCollectionGeoJsonAllOfFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertCollectionGeoJsonAllOfFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertCollectionGeoJsonAllOfFeatures(val *AlertCollectionGeoJsonAllOfFeatures) *NullableAlertCollectionGeoJsonAllOfFeatures {
	return &NullableAlertCollectionGeoJsonAllOfFeatures{value: val, isSet: true}
}

func (v NullableAlertCollectionGeoJsonAllOfFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertCollectionGeoJsonAllOfFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

