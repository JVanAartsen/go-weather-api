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

// AlertCollectionPagination Links for retrieving more data
type AlertCollectionPagination struct {
	// A link to the next set of alerts
	Next string `json:"next"`
}

// NewAlertCollectionPagination instantiates a new AlertCollectionPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertCollectionPagination(next string) *AlertCollectionPagination {
	this := AlertCollectionPagination{}
	this.Next = next
	return &this
}

// NewAlertCollectionPaginationWithDefaults instantiates a new AlertCollectionPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertCollectionPaginationWithDefaults() *AlertCollectionPagination {
	this := AlertCollectionPagination{}
	return &this
}

// GetNext returns the Next field value
func (o *AlertCollectionPagination) GetNext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *AlertCollectionPagination) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *AlertCollectionPagination) SetNext(v string) {
	o.Next = v
}

func (o AlertCollectionPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["next"] = o.Next
	}
	return json.Marshal(toSerialize)
}

type NullableAlertCollectionPagination struct {
	value *AlertCollectionPagination
	isSet bool
}

func (v NullableAlertCollectionPagination) Get() *AlertCollectionPagination {
	return v.value
}

func (v *NullableAlertCollectionPagination) Set(val *AlertCollectionPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertCollectionPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertCollectionPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertCollectionPagination(val *AlertCollectionPagination) *NullableAlertCollectionPagination {
	return &NullableAlertCollectionPagination{value: val, isSet: true}
}

func (v NullableAlertCollectionPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertCollectionPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


