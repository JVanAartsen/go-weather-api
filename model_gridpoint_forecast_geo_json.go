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

// GridpointForecastGeoJson struct for GridpointForecastGeoJson
type GridpointForecastGeoJson struct {
	Context *JsonLdContext `json:"@context,omitempty"`
	Id *string `json:"id,omitempty"`
	Type string `json:"type"`
	Geometry NullableGeoJsonGeometry `json:"geometry"`
	Properties GridpointForecast `json:"properties"`
}

// NewGridpointForecastGeoJson instantiates a new GridpointForecastGeoJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridpointForecastGeoJson(type_ string, geometry NullableGeoJsonGeometry, properties GridpointForecast) *GridpointForecastGeoJson {
	this := GridpointForecastGeoJson{}
	this.Type = type_
	this.Geometry = geometry
	this.Properties = properties
	return &this
}

// NewGridpointForecastGeoJsonWithDefaults instantiates a new GridpointForecastGeoJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridpointForecastGeoJsonWithDefaults() *GridpointForecastGeoJson {
	this := GridpointForecastGeoJson{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *GridpointForecastGeoJson) GetContext() JsonLdContext {
	if o == nil || o.Context == nil {
		var ret JsonLdContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastGeoJson) GetContextOk() (*JsonLdContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *GridpointForecastGeoJson) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given JsonLdContext and assigns it to the Context field.
func (o *GridpointForecastGeoJson) SetContext(v JsonLdContext) {
	o.Context = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GridpointForecastGeoJson) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastGeoJson) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GridpointForecastGeoJson) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GridpointForecastGeoJson) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value
func (o *GridpointForecastGeoJson) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GridpointForecastGeoJson) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GridpointForecastGeoJson) SetType(v string) {
	o.Type = v
}

// GetGeometry returns the Geometry field value
// If the value is explicit nil, the zero value for GeoJsonGeometry will be returned
func (o *GridpointForecastGeoJson) GetGeometry() GeoJsonGeometry {
	if o == nil || o.Geometry.Get() == nil {
		var ret GeoJsonGeometry
		return ret
	}

	return *o.Geometry.Get()
}

// GetGeometryOk returns a tuple with the Geometry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GridpointForecastGeoJson) GetGeometryOk() (*GeoJsonGeometry, bool) {
	if o == nil {
		return nil, false
	}
	return o.Geometry.Get(), o.Geometry.IsSet()
}

// SetGeometry sets field value
func (o *GridpointForecastGeoJson) SetGeometry(v GeoJsonGeometry) {
	o.Geometry.Set(&v)
}

// GetProperties returns the Properties field value
func (o *GridpointForecastGeoJson) GetProperties() GridpointForecast {
	if o == nil {
		var ret GridpointForecast
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *GridpointForecastGeoJson) GetPropertiesOk() (*GridpointForecast, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *GridpointForecastGeoJson) SetProperties(v GridpointForecast) {
	o.Properties = v
}

func (o GridpointForecastGeoJson) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["@context"] = o.Context
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["geometry"] = o.Geometry.Get()
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableGridpointForecastGeoJson struct {
	value *GridpointForecastGeoJson
	isSet bool
}

func (v NullableGridpointForecastGeoJson) Get() *GridpointForecastGeoJson {
	return v.value
}

func (v *NullableGridpointForecastGeoJson) Set(val *GridpointForecastGeoJson) {
	v.value = val
	v.isSet = true
}

func (v NullableGridpointForecastGeoJson) IsSet() bool {
	return v.isSet
}

func (v *NullableGridpointForecastGeoJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridpointForecastGeoJson(val *GridpointForecastGeoJson) *NullableGridpointForecastGeoJson {
	return &NullableGridpointForecastGeoJson{value: val, isSet: true}
}

func (v NullableGridpointForecastGeoJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridpointForecastGeoJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


