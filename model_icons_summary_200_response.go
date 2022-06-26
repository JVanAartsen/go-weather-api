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

// IconsSummary200Response struct for IconsSummary200Response
type IconsSummary200Response struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Icons map[string]IconsSummary200ResponseIconsValue `json:"icons"`
}

// NewIconsSummary200Response instantiates a new IconsSummary200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIconsSummary200Response(icons map[string]IconsSummary200ResponseIconsValue) *IconsSummary200Response {
	this := IconsSummary200Response{}
	this.Icons = icons
	return &this
}

// NewIconsSummary200ResponseWithDefaults instantiates a new IconsSummary200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIconsSummary200ResponseWithDefaults() *IconsSummary200Response {
	this := IconsSummary200Response{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *IconsSummary200Response) GetContext() JsonLdContext {
	if o == nil || o.Context == nil {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IconsSummary200Response) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *IconsSummary200Response) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *IconsSummary200Response) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetIcons returns the Icons field value
func (o *IconsSummary200Response) GetIcons() map[string]IconsSummary200ResponseIconsValue {
	if o == nil {
		var ret map[string]IconsSummary200ResponseIconsValue
		return ret
	}

	return o.Icons
}

// GetIconsOk returns a tuple with the Icons field value
// and a boolean to check if the value has been set.
func (o *IconsSummary200Response) GetIconsOk() (*map[string]IconsSummary200ResponseIconsValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Icons, true
}

// SetIcons sets field value
func (o *IconsSummary200Response) SetIcons(v map[string]IconsSummary200ResponseIconsValue) {
	o.Icons = v
}

func (o IconsSummary200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["@context"] = o.Context
	}
	if true {
		toSerialize["icons"] = o.Icons
	}
	return json.Marshal(toSerialize)
}

type NullableIconsSummary200Response struct {
	value *IconsSummary200Response
	isSet bool
}

func (v NullableIconsSummary200Response) Get() *IconsSummary200Response {
	return v.value
}

func (v *NullableIconsSummary200Response) Set(val *IconsSummary200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIconsSummary200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIconsSummary200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIconsSummary200Response(val *IconsSummary200Response) *NullableIconsSummary200Response {
	return &NullableIconsSummary200Response{value: val, isSet: true}
}

func (v NullableIconsSummary200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIconsSummary200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


