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

// ObservationCollectionJsonLd struct for ObservationCollectionJsonLd
type ObservationCollectionJsonLd struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Graph []Observation `json:"@graph,omitempty"`
}

// NewObservationCollectionJsonLd instantiates a new ObservationCollectionJsonLd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationCollectionJsonLd() *ObservationCollectionJsonLd {
	this := ObservationCollectionJsonLd{}
	return &this
}

// NewObservationCollectionJsonLdWithDefaults instantiates a new ObservationCollectionJsonLd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationCollectionJsonLdWithDefaults() *ObservationCollectionJsonLd {
	this := ObservationCollectionJsonLd{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ObservationCollectionJsonLd) GetContext() JsonLdContext {
	if o == nil || o.Context == nil {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationCollectionJsonLd) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ObservationCollectionJsonLd) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *ObservationCollectionJsonLd) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetGraph returns the Graph field value if set, zero value otherwise.
func (o *ObservationCollectionJsonLd) GetGraph() []Observation {
	if o == nil || o.Graph == nil {
		var ret []Observation
		return ret
	}
	return o.Graph
}

// GetGraphOk returns a tuple with the Graph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationCollectionJsonLd) GetGraphOk() ([]Observation, bool) {
	if o == nil || o.Graph == nil {
		return nil, false
	}
	return o.Graph, true
}

// HasGraph returns a boolean if a field has been set.
func (o *ObservationCollectionJsonLd) HasGraph() bool {
	if o != nil && o.Graph != nil {
		return true
	}

	return false
}

// SetGraph gets a reference to the given []Observation and assigns it to the Graph field.
func (o *ObservationCollectionJsonLd) SetGraph(v []Observation) {
	o.Graph = v
}

func (o ObservationCollectionJsonLd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["@context"] = o.Context
	}
	if o.Graph != nil {
		toSerialize["@graph"] = o.Graph
	}
	return json.Marshal(toSerialize)
}

type NullableObservationCollectionJsonLd struct {
	value *ObservationCollectionJsonLd
	isSet bool
}

func (v NullableObservationCollectionJsonLd) Get() *ObservationCollectionJsonLd {
	return v.value
}

func (v *NullableObservationCollectionJsonLd) Set(val *ObservationCollectionJsonLd) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationCollectionJsonLd) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationCollectionJsonLd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationCollectionJsonLd(val *ObservationCollectionJsonLd) *NullableObservationCollectionJsonLd {
	return &NullableObservationCollectionJsonLd{value: val, isSet: true}
}

func (v NullableObservationCollectionJsonLd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationCollectionJsonLd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


