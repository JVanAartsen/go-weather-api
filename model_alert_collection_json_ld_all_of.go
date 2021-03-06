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

// AlertCollectionJsonLdAllOf struct for AlertCollectionJsonLdAllOf
type AlertCollectionJsonLdAllOf struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Graph []Alert `json:"@graph,omitempty"`
}

// NewAlertCollectionJsonLdAllOf instantiates a new AlertCollectionJsonLdAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertCollectionJsonLdAllOf() *AlertCollectionJsonLdAllOf {
	this := AlertCollectionJsonLdAllOf{}
	return &this
}

// NewAlertCollectionJsonLdAllOfWithDefaults instantiates a new AlertCollectionJsonLdAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertCollectionJsonLdAllOfWithDefaults() *AlertCollectionJsonLdAllOf {
	this := AlertCollectionJsonLdAllOf{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *AlertCollectionJsonLdAllOf) GetContext() JsonLdContext {
	if o == nil || o.Context == nil {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCollectionJsonLdAllOf) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *AlertCollectionJsonLdAllOf) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *AlertCollectionJsonLdAllOf) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetGraph returns the Graph field value if set, zero value otherwise.
func (o *AlertCollectionJsonLdAllOf) GetGraph() []Alert {
	if o == nil || o.Graph == nil {
		var ret []Alert
		return ret
	}
	return o.Graph
}

// GetGraphOk returns a tuple with the Graph field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertCollectionJsonLdAllOf) GetGraphOk() ([]Alert, bool) {
	if o == nil || o.Graph == nil {
		return nil, false
	}
	return o.Graph, true
}

// HasGraph returns a boolean if a field has been set.
func (o *AlertCollectionJsonLdAllOf) HasGraph() bool {
	if o != nil && o.Graph != nil {
		return true
	}

	return false
}

// SetGraph gets a reference to the given []Alert and assigns it to the Graph field.
func (o *AlertCollectionJsonLdAllOf) SetGraph(v []Alert) {
	o.Graph = v
}

func (o AlertCollectionJsonLdAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["@context"] = o.Context
	}
	if o.Graph != nil {
		toSerialize["@graph"] = o.Graph
	}
	return json.Marshal(toSerialize)
}

type NullableAlertCollectionJsonLdAllOf struct {
	value *AlertCollectionJsonLdAllOf
	isSet bool
}

func (v NullableAlertCollectionJsonLdAllOf) Get() *AlertCollectionJsonLdAllOf {
	return v.value
}

func (v *NullableAlertCollectionJsonLdAllOf) Set(val *AlertCollectionJsonLdAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertCollectionJsonLdAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertCollectionJsonLdAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertCollectionJsonLdAllOf(val *AlertCollectionJsonLdAllOf) *NullableAlertCollectionJsonLdAllOf {
	return &NullableAlertCollectionJsonLdAllOf{value: val, isSet: true}
}

func (v NullableAlertCollectionJsonLdAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertCollectionJsonLdAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


