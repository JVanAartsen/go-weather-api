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

// OfficeHeadlineCollection struct for OfficeHeadlineCollection
type OfficeHeadlineCollection struct {
	Context JsonLdContext `json:"@context"`
	Graph []OfficeHeadline `json:"@graph"`
}

// NewOfficeHeadlineCollection instantiates a new OfficeHeadlineCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfficeHeadlineCollection(context JsonLdContext, graph []OfficeHeadline) *OfficeHeadlineCollection {
	this := OfficeHeadlineCollection{}
	this.Context = context
	this.Graph = graph
	return &this
}

// NewOfficeHeadlineCollectionWithDefaults instantiates a new OfficeHeadlineCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfficeHeadlineCollectionWithDefaults() *OfficeHeadlineCollection {
	this := OfficeHeadlineCollection{}
	return &this
}

// GetContext returns the Context field value
func (o *OfficeHeadlineCollection) GetContext() JsonLdContext {
	if o == nil {
		var ret JsonLdContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *OfficeHeadlineCollection) GetContextOk() (*JsonLdContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *OfficeHeadlineCollection) SetContext(v JsonLdContext) {
	o.Context = v
}

// GetGraph returns the Graph field value
func (o *OfficeHeadlineCollection) GetGraph() []OfficeHeadline {
	if o == nil {
		var ret []OfficeHeadline
		return ret
	}

	return o.Graph
}

// GetGraphOk returns a tuple with the Graph field value
// and a boolean to check if the value has been set.
func (o *OfficeHeadlineCollection) GetGraphOk() ([]OfficeHeadline, bool) {
	if o == nil {
		return nil, false
	}
	return o.Graph, true
}

// SetGraph sets field value
func (o *OfficeHeadlineCollection) SetGraph(v []OfficeHeadline) {
	o.Graph = v
}

func (o OfficeHeadlineCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["@context"] = o.Context
	}
	if true {
		toSerialize["@graph"] = o.Graph
	}
	return json.Marshal(toSerialize)
}

type NullableOfficeHeadlineCollection struct {
	value *OfficeHeadlineCollection
	isSet bool
}

func (v NullableOfficeHeadlineCollection) Get() *OfficeHeadlineCollection {
	return v.value
}

func (v *NullableOfficeHeadlineCollection) Set(val *OfficeHeadlineCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableOfficeHeadlineCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableOfficeHeadlineCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfficeHeadlineCollection(val *OfficeHeadlineCollection) *NullableOfficeHeadlineCollection {
	return &NullableOfficeHeadlineCollection{value: val, isSet: true}
}

func (v NullableOfficeHeadlineCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfficeHeadlineCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


