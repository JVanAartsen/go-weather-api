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

// ZoneForecastPeriodsInner struct for ZoneForecastPeriodsInner
type ZoneForecastPeriodsInner struct {
	// A sequential identifier number.
	Number int32 `json:"number"`
	// A textual description of the period.
	Name string `json:"name"`
	// A detailed textual forecast for the period.
	DetailedForecast string `json:"detailedForecast"`
}

// NewZoneForecastPeriodsInner instantiates a new ZoneForecastPeriodsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZoneForecastPeriodsInner(number int32, name string, detailedForecast string) *ZoneForecastPeriodsInner {
	this := ZoneForecastPeriodsInner{}
	this.Number = number
	this.Name = name
	this.DetailedForecast = detailedForecast
	return &this
}

// NewZoneForecastPeriodsInnerWithDefaults instantiates a new ZoneForecastPeriodsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneForecastPeriodsInnerWithDefaults() *ZoneForecastPeriodsInner {
	this := ZoneForecastPeriodsInner{}
	return &this
}

// GetNumber returns the Number field value
func (o *ZoneForecastPeriodsInner) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *ZoneForecastPeriodsInner) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *ZoneForecastPeriodsInner) SetNumber(v int32) {
	o.Number = v
}

// GetName returns the Name field value
func (o *ZoneForecastPeriodsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ZoneForecastPeriodsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ZoneForecastPeriodsInner) SetName(v string) {
	o.Name = v
}

// GetDetailedForecast returns the DetailedForecast field value
func (o *ZoneForecastPeriodsInner) GetDetailedForecast() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DetailedForecast
}

// GetDetailedForecastOk returns a tuple with the DetailedForecast field value
// and a boolean to check if the value has been set.
func (o *ZoneForecastPeriodsInner) GetDetailedForecastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetailedForecast, true
}

// SetDetailedForecast sets field value
func (o *ZoneForecastPeriodsInner) SetDetailedForecast(v string) {
	o.DetailedForecast = v
}

func (o ZoneForecastPeriodsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["detailedForecast"] = o.DetailedForecast
	}
	return json.Marshal(toSerialize)
}

type NullableZoneForecastPeriodsInner struct {
	value *ZoneForecastPeriodsInner
	isSet bool
}

func (v NullableZoneForecastPeriodsInner) Get() *ZoneForecastPeriodsInner {
	return v.value
}

func (v *NullableZoneForecastPeriodsInner) Set(val *ZoneForecastPeriodsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableZoneForecastPeriodsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableZoneForecastPeriodsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZoneForecastPeriodsInner(val *ZoneForecastPeriodsInner) *NullableZoneForecastPeriodsInner {
	return &NullableZoneForecastPeriodsInner{value: val, isSet: true}
}

func (v NullableZoneForecastPeriodsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZoneForecastPeriodsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


