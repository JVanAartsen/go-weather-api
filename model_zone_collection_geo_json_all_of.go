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

// ZoneCollectionGeoJsonAllOf struct for ZoneCollectionGeoJsonAllOf
type ZoneCollectionGeoJsonAllOf struct {
	Features []ZoneGeoJsonAllOf `json:"features,omitempty"`
}

// NewZoneCollectionGeoJsonAllOf instantiates a new ZoneCollectionGeoJsonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneCollectionGeoJsonAllOf() *ZoneCollectionGeoJsonAllOf {
	this := ZoneCollectionGeoJsonAllOf{}
	return &this
}

// NewZoneCollectionGeoJsonAllOfWithDefaults instantiates a new ZoneCollectionGeoJsonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneCollectionGeoJsonAllOfWithDefaults() *ZoneCollectionGeoJsonAllOf {
	this := ZoneCollectionGeoJsonAllOf{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ZoneCollectionGeoJsonAllOf) GetFeatures() []ZoneGeoJsonAllOf {
	if o == nil || o.Features == nil {
		var ret []ZoneGeoJsonAllOf
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZoneCollectionGeoJsonAllOf) GetFeaturesOk() ([]ZoneGeoJsonAllOf, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ZoneCollectionGeoJsonAllOf) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []ZoneGeoJsonAllOf and assigns it to the Features field.
func (o *ZoneCollectionGeoJsonAllOf) SetFeatures(v []ZoneGeoJsonAllOf) {
	o.Features = v
}

func (o ZoneCollectionGeoJsonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	return json.Marshal(toSerialize)
}

type NullableZoneCollectionGeoJsonAllOf struct {
	value *ZoneCollectionGeoJsonAllOf
	isSet bool
}

func (v NullableZoneCollectionGeoJsonAllOf) Get() *ZoneCollectionGeoJsonAllOf {
	return v.value
}

func (v *NullableZoneCollectionGeoJsonAllOf) Set(val *ZoneCollectionGeoJsonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneCollectionGeoJsonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneCollectionGeoJsonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneCollectionGeoJsonAllOf(val *ZoneCollectionGeoJsonAllOf) *NullableZoneCollectionGeoJsonAllOf {
	return &NullableZoneCollectionGeoJsonAllOf{value: val, isSet: true}
}

func (v NullableZoneCollectionGeoJsonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneCollectionGeoJsonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


