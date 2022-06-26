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

// OfficeAddress struct for OfficeAddress
type OfficeAddress struct {
	Type *string `json:"@type,omitempty"`
	StreetAddress *string `json:"streetAddress,omitempty"`
	AddressLocality *string `json:"addressLocality,omitempty"`
	AddressRegion *string `json:"addressRegion,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
}

// NewOfficeAddress instantiates a new OfficeAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfficeAddress() *OfficeAddress {
	this := OfficeAddress{}
	return &this
}

// NewOfficeAddressWithDefaults instantiates a new OfficeAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfficeAddressWithDefaults() *OfficeAddress {
	this := OfficeAddress{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OfficeAddress) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeAddress) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OfficeAddress) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OfficeAddress) SetType(v string) {
	o.Type = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *OfficeAddress) GetStreetAddress() string {
	if o == nil || o.StreetAddress == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeAddress) GetStreetAddressOk() (*string, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *OfficeAddress) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given string and assigns it to the StreetAddress field.
func (o *OfficeAddress) SetStreetAddress(v string) {
	o.StreetAddress = &v
}

// GetAddressLocality returns the AddressLocality field value if set, zero value otherwise.
func (o *OfficeAddress) GetAddressLocality() string {
	if o == nil || o.AddressLocality == nil {
		var ret string
		return ret
	}
	return *o.AddressLocality
}

// GetAddressLocalityOk returns a tuple with the AddressLocality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeAddress) GetAddressLocalityOk() (*string, bool) {
	if o == nil || o.AddressLocality == nil {
		return nil, false
	}
	return o.AddressLocality, true
}

// HasAddressLocality returns a boolean if a field has been set.
func (o *OfficeAddress) HasAddressLocality() bool {
	if o != nil && o.AddressLocality != nil {
		return true
	}

	return false
}

// SetAddressLocality gets a reference to the given string and assigns it to the AddressLocality field.
func (o *OfficeAddress) SetAddressLocality(v string) {
	o.AddressLocality = &v
}

// GetAddressRegion returns the AddressRegion field value if set, zero value otherwise.
func (o *OfficeAddress) GetAddressRegion() string {
	if o == nil || o.AddressRegion == nil {
		var ret string
		return ret
	}
	return *o.AddressRegion
}

// GetAddressRegionOk returns a tuple with the AddressRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeAddress) GetAddressRegionOk() (*string, bool) {
	if o == nil || o.AddressRegion == nil {
		return nil, false
	}
	return o.AddressRegion, true
}

// HasAddressRegion returns a boolean if a field has been set.
func (o *OfficeAddress) HasAddressRegion() bool {
	if o != nil && o.AddressRegion != nil {
		return true
	}

	return false
}

// SetAddressRegion gets a reference to the given string and assigns it to the AddressRegion field.
func (o *OfficeAddress) SetAddressRegion(v string) {
	o.AddressRegion = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *OfficeAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *OfficeAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *OfficeAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

func (o OfficeAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["@type"] = o.Type
	}
	if o.StreetAddress != nil {
		toSerialize["streetAddress"] = o.StreetAddress
	}
	if o.AddressLocality != nil {
		toSerialize["addressLocality"] = o.AddressLocality
	}
	if o.AddressRegion != nil {
		toSerialize["addressRegion"] = o.AddressRegion
	}
	if o.PostalCode != nil {
		toSerialize["postalCode"] = o.PostalCode
	}
	return json.Marshal(toSerialize)
}

type NullableOfficeAddress struct {
	value *OfficeAddress
	isSet bool
}

func (v NullableOfficeAddress) Get() *OfficeAddress {
	return v.value
}

func (v *NullableOfficeAddress) Set(val *OfficeAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableOfficeAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableOfficeAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfficeAddress(val *OfficeAddress) *NullableOfficeAddress {
	return &NullableOfficeAddress{value: val, isSet: true}
}

func (v NullableOfficeAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfficeAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

