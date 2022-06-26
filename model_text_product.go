/*
weather.gov API

weather.gov API

API version: 1.8.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package weatherApi

import (
	"encoding/json"
	"time"
)

// TextProduct struct for TextProduct
type TextProduct struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Id *string `json:"@id,omitempty"`
	Id *string `json:"id,omitempty"`
	WmoCollectiveId *string `json:"wmoCollectiveId,omitempty"`
	IssuingOffice *string `json:"issuingOffice,omitempty"`
	IssuanceTime *time.Time `json:"issuanceTime,omitempty"`
	ProductCode *string `json:"productCode,omitempty"`
	ProductName *string `json:"productName,omitempty"`
	ProductText *string `json:"productText,omitempty"`
}

// NewTextProduct instantiates a new TextProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextProduct() *TextProduct {
	this := TextProduct{}
	return &this
}

// NewTextProductWithDefaults instantiates a new TextProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextProductWithDefaults() *TextProduct {
	this := TextProduct{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TextProduct) GetContext() JsonLdContext {
	if o == nil || o.Context == nil {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextProduct) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TextProduct) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *TextProduct) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TextProduct) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextProduct) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TextProduct) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TextProduct) SetId(v string) {
	o.Id = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TextProduct) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextProduct) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TextProduct) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TextProduct) SetId(v string) {
	o.Id = &v
}

// GetWmoCollectiveId returns the WmoCollectiveId field value if set, zero value otherwise.
func (o *TextProduct) GetWmoCollectiveId() string {
	if o == nil || o.WmoCollectiveId == nil {
		var ret string
		return ret
	}
	return *o.WmoCollectiveId
}

// GetWmoCollectiveIdOk returns a tuple with the WmoCollectiveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextProduct) GetWmoCollectiveIdOk() (*string, bool) {
	if o == nil || o.WmoCollectiveId == nil {
		return nil, false
	}
	return o.WmoCollectiveId, true
}

// HasWmoCollectiveId returns a boolean if a field has been set.
func (o *TextProduct) HasWmoCollectiveId() bool {
	if o != nil && o.WmoCollectiveId != nil {
		return true
	}

	return false
}

// SetWmoCollectiveId gets a reference to the given string and assigns it to the WmoCollectiveId field.
func (o *TextProduct) SetWmoCollectiveId(v string) {
	o.WmoCollectiveId = &v
}

// GetIssuingOffice returns the IssuingOffice field value if set, zero value otherwise.
func (o *TextProduct) GetIssuingOffice() string {
	if o == nil || o.IssuingOffice == nil {
		var ret string
		return ret
	}
	return *o.IssuingOffice
}

// GetIssuingOfficeOk returns a tuple with the IssuingOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextProduct) GetIssuingOfficeOk() (*string, bool) {
	if o == nil || o.IssuingOffice == nil {
		return nil, false
	}
	return o.IssuingOffice, true
}

// HasIssuingOffice returns a boolean if a field has been set.
func (o *TextProduct) HasIssuingOffice() bool {
	if o != nil && o.IssuingOffice != nil {
		return true
	}

	return false
}

// SetIssuingOffice gets a reference to the given string and assigns it to the IssuingOffice field.
func (o *TextProduct) SetIssuingOffice(v string) {
	o.IssuingOffice = &v
}

// GetIssuanceTime returns the IssuanceTime field value if set, zero value otherwise.
func (o *TextProduct) GetIssuanceTime() time.Time {
	if o == nil || o.IssuanceTime == nil {
		var ret time.Time
		return ret
	}
	return *o.IssuanceTime
}

// GetIssuanceTimeOk returns a tuple with the IssuanceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextProduct) GetIssuanceTimeOk() (*time.Time, bool) {
	if o == nil || o.IssuanceTime == nil {
		return nil, false
	}
	return o.IssuanceTime, true
}

// HasIssuanceTime returns a boolean if a field has been set.
func (o *TextProduct) HasIssuanceTime() bool {
	if o != nil && o.IssuanceTime != nil {
		return true
	}

	return false
}

// SetIssuanceTime gets a reference to the given time.Time and assigns it to the IssuanceTime field.
func (o *TextProduct) SetIssuanceTime(v time.Time) {
	o.IssuanceTime = &v
}

// GetProductCode returns the ProductCode field value if set, zero value otherwise.
func (o *TextProduct) GetProductCode() string {
	if o == nil || o.ProductCode == nil {
		var ret string
		return ret
	}
	return *o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextProduct) GetProductCodeOk() (*string, bool) {
	if o == nil || o.ProductCode == nil {
		return nil, false
	}
	return o.ProductCode, true
}

// HasProductCode returns a boolean if a field has been set.
func (o *TextProduct) HasProductCode() bool {
	if o != nil && o.ProductCode != nil {
		return true
	}

	return false
}

// SetProductCode gets a reference to the given string and assigns it to the ProductCode field.
func (o *TextProduct) SetProductCode(v string) {
	o.ProductCode = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *TextProduct) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextProduct) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *TextProduct) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *TextProduct) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductText returns the ProductText field value if set, zero value otherwise.
func (o *TextProduct) GetProductText() string {
	if o == nil || o.ProductText == nil {
		var ret string
		return ret
	}
	return *o.ProductText
}

// GetProductTextOk returns a tuple with the ProductText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextProduct) GetProductTextOk() (*string, bool) {
	if o == nil || o.ProductText == nil {
		return nil, false
	}
	return o.ProductText, true
}

// HasProductText returns a boolean if a field has been set.
func (o *TextProduct) HasProductText() bool {
	if o != nil && o.ProductText != nil {
		return true
	}

	return false
}

// SetProductText gets a reference to the given string and assigns it to the ProductText field.
func (o *TextProduct) SetProductText(v string) {
	o.ProductText = &v
}

func (o TextProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["@context"] = o.Context
	}
	if o.Id != nil {
		toSerialize["@id"] = o.Id
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.WmoCollectiveId != nil {
		toSerialize["wmoCollectiveId"] = o.WmoCollectiveId
	}
	if o.IssuingOffice != nil {
		toSerialize["issuingOffice"] = o.IssuingOffice
	}
	if o.IssuanceTime != nil {
		toSerialize["issuanceTime"] = o.IssuanceTime
	}
	if o.ProductCode != nil {
		toSerialize["productCode"] = o.ProductCode
	}
	if o.ProductName != nil {
		toSerialize["productName"] = o.ProductName
	}
	if o.ProductText != nil {
		toSerialize["productText"] = o.ProductText
	}
	return json.Marshal(toSerialize)
}

type NullableTextProduct struct {
	value *TextProduct
	isSet bool
}

func (v NullableTextProduct) Get() *TextProduct {
	return v.value
}

func (v *NullableTextProduct) Set(val *TextProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableTextProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableTextProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextProduct(val *TextProduct) *NullableTextProduct {
	return &NullableTextProduct{value: val, isSet: true}
}

func (v NullableTextProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

