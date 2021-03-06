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

// GridpointHazardsValuesInnerValueInner A value object representing an expected hazard.
type GridpointHazardsValuesInnerValueInner struct {
	// Hazard code. This value will correspond to a P-VTEC phenomenon code as defined in NWS Directive 10-1703. 
	Phenomenon string `json:"phenomenon"`
	// Significance code. This value will correspond to a P-VTEC significance code as defined in NWS Directive 10-1703. This will most frequently be \"A\" for a watch or \"Y\" for an advisory. 
	Significance string `json:"significance"`
	// Event number. If this hazard refers to a national or regional center product (such as a Storm Prediction Center convective watch), this value will be the sequence number of that product. 
	EventNumber NullableInt32 `json:"event_number"`
}

// NewGridpointHazardsValuesInnerValueInner instantiates a new GridpointHazardsValuesInnerValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridpointHazardsValuesInnerValueInner(phenomenon string, significance string, eventNumber NullableInt32) *GridpointHazardsValuesInnerValueInner {
	this := GridpointHazardsValuesInnerValueInner{}
	this.Phenomenon = phenomenon
	this.Significance = significance
	this.EventNumber = eventNumber
	return &this
}

// NewGridpointHazardsValuesInnerValueInnerWithDefaults instantiates a new GridpointHazardsValuesInnerValueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridpointHazardsValuesInnerValueInnerWithDefaults() *GridpointHazardsValuesInnerValueInner {
	this := GridpointHazardsValuesInnerValueInner{}
	return &this
}

// GetPhenomenon returns the Phenomenon field value
func (o *GridpointHazardsValuesInnerValueInner) GetPhenomenon() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phenomenon
}

// GetPhenomenonOk returns a tuple with the Phenomenon field value
// and a boolean to check if the value has been set.
func (o *GridpointHazardsValuesInnerValueInner) GetPhenomenonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phenomenon, true
}

// SetPhenomenon sets field value
func (o *GridpointHazardsValuesInnerValueInner) SetPhenomenon(v string) {
	o.Phenomenon = v
}

// GetSignificance returns the Significance field value
func (o *GridpointHazardsValuesInnerValueInner) GetSignificance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Significance
}

// GetSignificanceOk returns a tuple with the Significance field value
// and a boolean to check if the value has been set.
func (o *GridpointHazardsValuesInnerValueInner) GetSignificanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Significance, true
}

// SetSignificance sets field value
func (o *GridpointHazardsValuesInnerValueInner) SetSignificance(v string) {
	o.Significance = v
}

// GetEventNumber returns the EventNumber field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *GridpointHazardsValuesInnerValueInner) GetEventNumber() int32 {
	if o == nil || o.EventNumber.Get() == nil {
		var ret int32
		return ret
	}

	return *o.EventNumber.Get()
}

// GetEventNumberOk returns a tuple with the EventNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GridpointHazardsValuesInnerValueInner) GetEventNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventNumber.Get(), o.EventNumber.IsSet()
}

// SetEventNumber sets field value
func (o *GridpointHazardsValuesInnerValueInner) SetEventNumber(v int32) {
	o.EventNumber.Set(&v)
}

func (o GridpointHazardsValuesInnerValueInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["phenomenon"] = o.Phenomenon
	}
	if true {
		toSerialize["significance"] = o.Significance
	}
	if true {
		toSerialize["event_number"] = o.EventNumber.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableGridpointHazardsValuesInnerValueInner struct {
	value *GridpointHazardsValuesInnerValueInner
	isSet bool
}

func (v NullableGridpointHazardsValuesInnerValueInner) Get() *GridpointHazardsValuesInnerValueInner {
	return v.value
}

func (v *NullableGridpointHazardsValuesInnerValueInner) Set(val *GridpointHazardsValuesInnerValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGridpointHazardsValuesInnerValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGridpointHazardsValuesInnerValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridpointHazardsValuesInnerValueInner(val *GridpointHazardsValuesInnerValueInner) *NullableGridpointHazardsValuesInnerValueInner {
	return &NullableGridpointHazardsValuesInnerValueInner{value: val, isSet: true}
}

func (v NullableGridpointHazardsValuesInnerValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridpointHazardsValuesInnerValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


