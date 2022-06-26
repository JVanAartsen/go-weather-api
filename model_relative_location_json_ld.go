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

// RelativeLocationJsonLd struct for RelativeLocationJsonLd
type RelativeLocationJsonLd struct {
	City *string `json:"city,omitempty"`
	State *string `json:"state,omitempty"`
	Distance *QuantitativeValue `json:"distance,omitempty"`
	Bearing *QuantitativeValue `json:"bearing,omitempty"`
	// A geometry represented in Well-Known Text (WKT) format.
	Geometry NullableString `json:"geometry"`
}

// NewRelativeLocationJsonLd instantiates a new RelativeLocationJsonLd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelativeLocationJsonLd(geometry NullableString) *RelativeLocationJsonLd {
	this := RelativeLocationJsonLd{}
	this.Geometry = geometry
	return &this
}

// NewRelativeLocationJsonLdWithDefaults instantiates a new RelativeLocationJsonLd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelativeLocationJsonLdWithDefaults() *RelativeLocationJsonLd {
	this := RelativeLocationJsonLd{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *RelativeLocationJsonLd) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelativeLocationJsonLd) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *RelativeLocationJsonLd) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *RelativeLocationJsonLd) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RelativeLocationJsonLd) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelativeLocationJsonLd) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RelativeLocationJsonLd) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *RelativeLocationJsonLd) SetState(v string) {
	o.State = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *RelativeLocationJsonLd) GetDistance() QuantitativeValue {
	if o == nil || o.Distance == nil {
		var ret QuantitativeValue
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelativeLocationJsonLd) GetDistanceOk() (*QuantitativeValue, bool) {
	if o == nil || o.Distance == nil {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *RelativeLocationJsonLd) HasDistance() bool {
	if o != nil && o.Distance != nil {
		return true
	}

	return false
}

// SetDistance gets a reference to the given QuantitativeValue and assigns it to the Distance field.
func (o *RelativeLocationJsonLd) SetDistance(v QuantitativeValue) {
	o.Distance = &v
}

// GetBearing returns the Bearing field value if set, zero value otherwise.
func (o *RelativeLocationJsonLd) GetBearing() QuantitativeValue {
	if o == nil || o.Bearing == nil {
		var ret QuantitativeValue
		return ret
	}
	return *o.Bearing
}

// GetBearingOk returns a tuple with the Bearing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelativeLocationJsonLd) GetBearingOk() (*QuantitativeValue, bool) {
	if o == nil || o.Bearing == nil {
		return nil, false
	}
	return o.Bearing, true
}

// HasBearing returns a boolean if a field has been set.
func (o *RelativeLocationJsonLd) HasBearing() bool {
	if o != nil && o.Bearing != nil {
		return true
	}

	return false
}

// SetBearing gets a reference to the given QuantitativeValue and assigns it to the Bearing field.
func (o *RelativeLocationJsonLd) SetBearing(v QuantitativeValue) {
	o.Bearing = &v
}

// GetGeometry returns the Geometry field value
// If the value is explicit nil, the zero value for string will be returned
func (o *RelativeLocationJsonLd) GetGeometry() string {
	if o == nil || o.Geometry.Get() == nil {
		var ret string
		return ret
	}

	return *o.Geometry.Get()
}

// GetGeometryOk returns a tuple with the Geometry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RelativeLocationJsonLd) GetGeometryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Geometry.Get(), o.Geometry.IsSet()
}

// SetGeometry sets field value
func (o *RelativeLocationJsonLd) SetGeometry(v string) {
	o.Geometry.Set(&v)
}

func (o RelativeLocationJsonLd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Distance != nil {
		toSerialize["distance"] = o.Distance
	}
	if o.Bearing != nil {
		toSerialize["bearing"] = o.Bearing
	}
	if true {
		toSerialize["geometry"] = o.Geometry.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRelativeLocationJsonLd struct {
	value *RelativeLocationJsonLd
	isSet bool
}

func (v NullableRelativeLocationJsonLd) Get() *RelativeLocationJsonLd {
	return v.value
}

func (v *NullableRelativeLocationJsonLd) Set(val *RelativeLocationJsonLd) {
	v.value = val
	v.isSet = true
}

func (v NullableRelativeLocationJsonLd) IsSet() bool {
	return v.isSet
}

func (v *NullableRelativeLocationJsonLd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelativeLocationJsonLd(val *RelativeLocationJsonLd) *NullableRelativeLocationJsonLd {
	return &NullableRelativeLocationJsonLd{value: val, isSet: true}
}

func (v NullableRelativeLocationJsonLd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelativeLocationJsonLd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

