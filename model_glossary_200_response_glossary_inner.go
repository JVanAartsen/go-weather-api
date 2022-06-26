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

// Glossary200ResponseGlossaryInner struct for Glossary200ResponseGlossaryInner
type Glossary200ResponseGlossaryInner struct {
	// The term being defined
	Term *string `json:"term,omitempty"`
	// A definition for the term
	Definition *string `json:"definition,omitempty"`
}

// NewGlossary200ResponseGlossaryInner instantiates a new Glossary200ResponseGlossaryInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlossary200ResponseGlossaryInner() *Glossary200ResponseGlossaryInner {
	this := Glossary200ResponseGlossaryInner{}
	return &this
}

// NewGlossary200ResponseGlossaryInnerWithDefaults instantiates a new Glossary200ResponseGlossaryInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlossary200ResponseGlossaryInnerWithDefaults() *Glossary200ResponseGlossaryInner {
	this := Glossary200ResponseGlossaryInner{}
	return &this
}

// GetTerm returns the Term field value if set, zero value otherwise.
func (o *Glossary200ResponseGlossaryInner) GetTerm() string {
	if o == nil || o.Term == nil {
		var ret string
		return ret
	}
	return *o.Term
}

// GetTermOk returns a tuple with the Term field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Glossary200ResponseGlossaryInner) GetTermOk() (*string, bool) {
	if o == nil || o.Term == nil {
		return nil, false
	}
	return o.Term, true
}

// HasTerm returns a boolean if a field has been set.
func (o *Glossary200ResponseGlossaryInner) HasTerm() bool {
	if o != nil && o.Term != nil {
		return true
	}

	return false
}

// SetTerm gets a reference to the given string and assigns it to the Term field.
func (o *Glossary200ResponseGlossaryInner) SetTerm(v string) {
	o.Term = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *Glossary200ResponseGlossaryInner) GetDefinition() string {
	if o == nil || o.Definition == nil {
		var ret string
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Glossary200ResponseGlossaryInner) GetDefinitionOk() (*string, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *Glossary200ResponseGlossaryInner) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given string and assigns it to the Definition field.
func (o *Glossary200ResponseGlossaryInner) SetDefinition(v string) {
	o.Definition = &v
}

func (o Glossary200ResponseGlossaryInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Term != nil {
		toSerialize["term"] = o.Term
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	return json.Marshal(toSerialize)
}

type NullableGlossary200ResponseGlossaryInner struct {
	value *Glossary200ResponseGlossaryInner
	isSet bool
}

func (v NullableGlossary200ResponseGlossaryInner) Get() *Glossary200ResponseGlossaryInner {
	return v.value
}

func (v *NullableGlossary200ResponseGlossaryInner) Set(val *Glossary200ResponseGlossaryInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGlossary200ResponseGlossaryInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGlossary200ResponseGlossaryInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlossary200ResponseGlossaryInner(val *Glossary200ResponseGlossaryInner) *NullableGlossary200ResponseGlossaryInner {
	return &NullableGlossary200ResponseGlossaryInner{value: val, isSet: true}
}

func (v NullableGlossary200ResponseGlossaryInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlossary200ResponseGlossaryInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

