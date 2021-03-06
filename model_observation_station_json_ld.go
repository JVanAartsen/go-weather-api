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

// ObservationStationJsonLd struct for ObservationStationJsonLd
type ObservationStationJsonLd struct {
	Context JsonLdContext `json:"@context"`
	// A geometry represented in Well-Known Text (WKT) format.
	Geometry NullableString `json:"geometry"`
	Id *string `json:"@id,omitempty"`
	Type *string `json:"@type,omitempty"`
	Elevation *QuantitativeValue `json:"elevation,omitempty"`
	StationIdentifier *string `json:"stationIdentifier,omitempty"`
	Name *string `json:"name,omitempty"`
	TimeZone *string `json:"timeZone,omitempty"`
	// A link to the NWS public forecast zone containing this station.
	Forecast *string `json:"forecast,omitempty"`
	// A link to the NWS county zone containing this station.
	County *string `json:"county,omitempty"`
	// A link to the NWS fire weather forecast zone containing this station.
	FireWeatherZone *string `json:"fireWeatherZone,omitempty"`
}

// NewObservationStationJsonLd instantiates a new ObservationStationJsonLd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObservationStationJsonLd(context JsonLdContext, geometry NullableString) *ObservationStationJsonLd {
	this := ObservationStationJsonLd{}
	this.Context = context
	this.Geometry = geometry
	return &this
}

// NewObservationStationJsonLdWithDefaults instantiates a new ObservationStationJsonLd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObservationStationJsonLdWithDefaults() *ObservationStationJsonLd {
	this := ObservationStationJsonLd{}
	return &this
}

// GetContext returns the Context field value
func (o *ObservationStationJsonLd) GetContext() JsonLdContext {
	if o == nil {
		var ret JsonLdContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *ObservationStationJsonLd) GetContextOk() (*JsonLdContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *ObservationStationJsonLd) SetContext(v JsonLdContext) {
	o.Context = v
}

// GetGeometry returns the Geometry field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ObservationStationJsonLd) GetGeometry() string {
	if o == nil || o.Geometry.Get() == nil {
		var ret string
		return ret
	}

	return *o.Geometry.Get()
}

// GetGeometryOk returns a tuple with the Geometry field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObservationStationJsonLd) GetGeometryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Geometry.Get(), o.Geometry.IsSet()
}

// SetGeometry sets field value
func (o *ObservationStationJsonLd) SetGeometry(v string) {
	o.Geometry.Set(&v)
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ObservationStationJsonLd) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationJsonLd) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ObservationStationJsonLd) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ObservationStationJsonLd) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ObservationStationJsonLd) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationJsonLd) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ObservationStationJsonLd) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ObservationStationJsonLd) SetType(v string) {
	o.Type = &v
}

// GetElevation returns the Elevation field value if set, zero value otherwise.
func (o *ObservationStationJsonLd) GetElevation() QuantitativeValue {
	if o == nil || o.Elevation == nil {
		var ret QuantitativeValue
		return ret
	}
	return *o.Elevation
}

// GetElevationOk returns a tuple with the Elevation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationJsonLd) GetElevationOk() (*QuantitativeValue, bool) {
	if o == nil || o.Elevation == nil {
		return nil, false
	}
	return o.Elevation, true
}

// HasElevation returns a boolean if a field has been set.
func (o *ObservationStationJsonLd) HasElevation() bool {
	if o != nil && o.Elevation != nil {
		return true
	}

	return false
}

// SetElevation gets a reference to the given QuantitativeValue and assigns it to the Elevation field.
func (o *ObservationStationJsonLd) SetElevation(v QuantitativeValue) {
	o.Elevation = &v
}

// GetStationIdentifier returns the StationIdentifier field value if set, zero value otherwise.
func (o *ObservationStationJsonLd) GetStationIdentifier() string {
	if o == nil || o.StationIdentifier == nil {
		var ret string
		return ret
	}
	return *o.StationIdentifier
}

// GetStationIdentifierOk returns a tuple with the StationIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationJsonLd) GetStationIdentifierOk() (*string, bool) {
	if o == nil || o.StationIdentifier == nil {
		return nil, false
	}
	return o.StationIdentifier, true
}

// HasStationIdentifier returns a boolean if a field has been set.
func (o *ObservationStationJsonLd) HasStationIdentifier() bool {
	if o != nil && o.StationIdentifier != nil {
		return true
	}

	return false
}

// SetStationIdentifier gets a reference to the given string and assigns it to the StationIdentifier field.
func (o *ObservationStationJsonLd) SetStationIdentifier(v string) {
	o.StationIdentifier = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ObservationStationJsonLd) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationJsonLd) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ObservationStationJsonLd) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ObservationStationJsonLd) SetName(v string) {
	o.Name = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *ObservationStationJsonLd) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationJsonLd) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *ObservationStationJsonLd) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *ObservationStationJsonLd) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetForecast returns the Forecast field value if set, zero value otherwise.
func (o *ObservationStationJsonLd) GetForecast() string {
	if o == nil || o.Forecast == nil {
		var ret string
		return ret
	}
	return *o.Forecast
}

// GetForecastOk returns a tuple with the Forecast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationJsonLd) GetForecastOk() (*string, bool) {
	if o == nil || o.Forecast == nil {
		return nil, false
	}
	return o.Forecast, true
}

// HasForecast returns a boolean if a field has been set.
func (o *ObservationStationJsonLd) HasForecast() bool {
	if o != nil && o.Forecast != nil {
		return true
	}

	return false
}

// SetForecast gets a reference to the given string and assigns it to the Forecast field.
func (o *ObservationStationJsonLd) SetForecast(v string) {
	o.Forecast = &v
}

// GetCounty returns the County field value if set, zero value otherwise.
func (o *ObservationStationJsonLd) GetCounty() string {
	if o == nil || o.County == nil {
		var ret string
		return ret
	}
	return *o.County
}

// GetCountyOk returns a tuple with the County field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationJsonLd) GetCountyOk() (*string, bool) {
	if o == nil || o.County == nil {
		return nil, false
	}
	return o.County, true
}

// HasCounty returns a boolean if a field has been set.
func (o *ObservationStationJsonLd) HasCounty() bool {
	if o != nil && o.County != nil {
		return true
	}

	return false
}

// SetCounty gets a reference to the given string and assigns it to the County field.
func (o *ObservationStationJsonLd) SetCounty(v string) {
	o.County = &v
}

// GetFireWeatherZone returns the FireWeatherZone field value if set, zero value otherwise.
func (o *ObservationStationJsonLd) GetFireWeatherZone() string {
	if o == nil || o.FireWeatherZone == nil {
		var ret string
		return ret
	}
	return *o.FireWeatherZone
}

// GetFireWeatherZoneOk returns a tuple with the FireWeatherZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservationStationJsonLd) GetFireWeatherZoneOk() (*string, bool) {
	if o == nil || o.FireWeatherZone == nil {
		return nil, false
	}
	return o.FireWeatherZone, true
}

// HasFireWeatherZone returns a boolean if a field has been set.
func (o *ObservationStationJsonLd) HasFireWeatherZone() bool {
	if o != nil && o.FireWeatherZone != nil {
		return true
	}

	return false
}

// SetFireWeatherZone gets a reference to the given string and assigns it to the FireWeatherZone field.
func (o *ObservationStationJsonLd) SetFireWeatherZone(v string) {
	o.FireWeatherZone = &v
}

func (o ObservationStationJsonLd) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["@context"] = o.Context
	}
	if true {
		toSerialize["geometry"] = o.Geometry.Get()
	}
	if o.Id != nil {
		toSerialize["@id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["@type"] = o.Type
	}
	if o.Elevation != nil {
		toSerialize["elevation"] = o.Elevation
	}
	if o.StationIdentifier != nil {
		toSerialize["stationIdentifier"] = o.StationIdentifier
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.Forecast != nil {
		toSerialize["forecast"] = o.Forecast
	}
	if o.County != nil {
		toSerialize["county"] = o.County
	}
	if o.FireWeatherZone != nil {
		toSerialize["fireWeatherZone"] = o.FireWeatherZone
	}
	return json.Marshal(toSerialize)
}

type NullableObservationStationJsonLd struct {
	value *ObservationStationJsonLd
	isSet bool
}

func (v NullableObservationStationJsonLd) Get() *ObservationStationJsonLd {
	return v.value
}

func (v *NullableObservationStationJsonLd) Set(val *ObservationStationJsonLd) {
	v.value = val
	v.isSet = true
}

func (v NullableObservationStationJsonLd) IsSet() bool {
	return v.isSet
}

func (v *NullableObservationStationJsonLd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObservationStationJsonLd(val *ObservationStationJsonLd) *NullableObservationStationJsonLd {
	return &NullableObservationStationJsonLd{value: val, isSet: true}
}

func (v NullableObservationStationJsonLd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObservationStationJsonLd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


