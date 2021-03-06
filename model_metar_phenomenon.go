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

// MetarPhenomenon An object representing a decoded METAR phenomenon string.
type MetarPhenomenon struct {
	Intensity NullableString `json:"intensity"`
	Modifier NullableString `json:"modifier"`
	Weather string `json:"weather"`
	RawString string `json:"rawString"`
	InVicinity *bool `json:"inVicinity,omitempty"`
}

// NewMetarPhenomenon instantiates a new MetarPhenomenon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetarPhenomenon(intensity NullableString, modifier NullableString, weather string, rawString string) *MetarPhenomenon {
	this := MetarPhenomenon{}
	this.Intensity = intensity
	this.Modifier = modifier
	this.Weather = weather
	this.RawString = rawString
	return &this
}

// NewMetarPhenomenonWithDefaults instantiates a new MetarPhenomenon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetarPhenomenonWithDefaults() *MetarPhenomenon {
	this := MetarPhenomenon{}
	return &this
}

// GetIntensity returns the Intensity field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MetarPhenomenon) GetIntensity() string {
	if o == nil || o.Intensity.Get() == nil {
		var ret string
		return ret
	}

	return *o.Intensity.Get()
}

// GetIntensityOk returns a tuple with the Intensity field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetarPhenomenon) GetIntensityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Intensity.Get(), o.Intensity.IsSet()
}

// SetIntensity sets field value
func (o *MetarPhenomenon) SetIntensity(v string) {
	o.Intensity.Set(&v)
}

// GetModifier returns the Modifier field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MetarPhenomenon) GetModifier() string {
	if o == nil || o.Modifier.Get() == nil {
		var ret string
		return ret
	}

	return *o.Modifier.Get()
}

// GetModifierOk returns a tuple with the Modifier field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetarPhenomenon) GetModifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modifier.Get(), o.Modifier.IsSet()
}

// SetModifier sets field value
func (o *MetarPhenomenon) SetModifier(v string) {
	o.Modifier.Set(&v)
}

// GetWeather returns the Weather field value
func (o *MetarPhenomenon) GetWeather() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Weather
}

// GetWeatherOk returns a tuple with the Weather field value
// and a boolean to check if the value has been set.
func (o *MetarPhenomenon) GetWeatherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weather, true
}

// SetWeather sets field value
func (o *MetarPhenomenon) SetWeather(v string) {
	o.Weather = v
}

// GetRawString returns the RawString field value
func (o *MetarPhenomenon) GetRawString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RawString
}

// GetRawStringOk returns a tuple with the RawString field value
// and a boolean to check if the value has been set.
func (o *MetarPhenomenon) GetRawStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawString, true
}

// SetRawString sets field value
func (o *MetarPhenomenon) SetRawString(v string) {
	o.RawString = v
}

// GetInVicinity returns the InVicinity field value if set, zero value otherwise.
func (o *MetarPhenomenon) GetInVicinity() bool {
	if o == nil || o.InVicinity == nil {
		var ret bool
		return ret
	}
	return *o.InVicinity
}

// GetInVicinityOk returns a tuple with the InVicinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetarPhenomenon) GetInVicinityOk() (*bool, bool) {
	if o == nil || o.InVicinity == nil {
		return nil, false
	}
	return o.InVicinity, true
}

// HasInVicinity returns a boolean if a field has been set.
func (o *MetarPhenomenon) HasInVicinity() bool {
	if o != nil && o.InVicinity != nil {
		return true
	}

	return false
}

// SetInVicinity gets a reference to the given bool and assigns it to the InVicinity field.
func (o *MetarPhenomenon) SetInVicinity(v bool) {
	o.InVicinity = &v
}

func (o MetarPhenomenon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["intensity"] = o.Intensity.Get()
	}
	if true {
		toSerialize["modifier"] = o.Modifier.Get()
	}
	if true {
		toSerialize["weather"] = o.Weather
	}
	if true {
		toSerialize["rawString"] = o.RawString
	}
	if o.InVicinity != nil {
		toSerialize["inVicinity"] = o.InVicinity
	}
	return json.Marshal(toSerialize)
}

type NullableMetarPhenomenon struct {
	value *MetarPhenomenon
	isSet bool
}

func (v NullableMetarPhenomenon) Get() *MetarPhenomenon {
	return v.value
}

func (v *NullableMetarPhenomenon) Set(val *MetarPhenomenon) {
	v.value = val
	v.isSet = true
}

func (v NullableMetarPhenomenon) IsSet() bool {
	return v.isSet
}

func (v *NullableMetarPhenomenon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetarPhenomenon(val *MetarPhenomenon) *NullableMetarPhenomenon {
	return &NullableMetarPhenomenon{value: val, isSet: true}
}

func (v NullableMetarPhenomenon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetarPhenomenon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


