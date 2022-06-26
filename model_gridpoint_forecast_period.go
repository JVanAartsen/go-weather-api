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

// GridpointForecastPeriod An object containing forecast information for a specific time period (generally 12-hour or 1-hour). 
type GridpointForecastPeriod struct {
	// Sequential period number.
	Number *int32 `json:"number,omitempty"`
	// A textual identifier for the period. This value will not be present for hourly forecasts. 
	Name *string `json:"name,omitempty"`
	// The starting time that this forecast period is valid for.
	StartTime *time.Time `json:"startTime,omitempty"`
	// The ending time that this forecast period is valid for.
	EndTime *time.Time `json:"endTime,omitempty"`
	// Indicates whether this period is daytime or nighttime.
	IsDaytime *bool `json:"isDaytime,omitempty"`
	Temperature *GridpointForecastPeriodTemperature `json:"temperature,omitempty"`
	// The unit of the temperature value (Fahrenheit or Celsius). This property is deprecated. Future versions will indicate the unit within the quantitative value object for the temperature property. To make use of the future standard format now, set the \"forecast_temperature_qv\" feature flag on the request. 
	// Deprecated
	TemperatureUnit *string `json:"temperatureUnit,omitempty"`
	// If not null, indicates a non-diurnal temperature trend for the period (either rising temperature overnight, or falling temperature during the day) 
	TemperatureTrend NullableString `json:"temperatureTrend,omitempty"`
	WindSpeed *GridpointForecastPeriodWindSpeed `json:"windSpeed,omitempty"`
	WindGust NullableGridpointForecastPeriodWindGust `json:"windGust,omitempty"`
	// The prevailing direction of the wind for the period, using a 16-point compass.
	WindDirection *string `json:"windDirection,omitempty"`
	// A link to an icon representing the forecast summary.
	// Deprecated
	Icon *string `json:"icon,omitempty"`
	// A brief textual forecast summary for the period.
	ShortForecast *string `json:"shortForecast,omitempty"`
	// A detailed textual forecast for the period.
	DetailedForecast *string `json:"detailedForecast,omitempty"`
}

// NewGridpointForecastPeriod instantiates a new GridpointForecastPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGridpointForecastPeriod() *GridpointForecastPeriod {
	this := GridpointForecastPeriod{}
	return &this
}

// NewGridpointForecastPeriodWithDefaults instantiates a new GridpointForecastPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridpointForecastPeriodWithDefaults() *GridpointForecastPeriod {
	this := GridpointForecastPeriod{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *GridpointForecastPeriod) GetNumber() int32 {
	if o == nil || o.Number == nil {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastPeriod) GetNumberOk() (*int32, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *GridpointForecastPeriod) SetNumber(v int32) {
	o.Number = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GridpointForecastPeriod) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastPeriod) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GridpointForecastPeriod) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GridpointForecastPeriod) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastPeriod) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *GridpointForecastPeriod) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GridpointForecastPeriod) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastPeriod) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *GridpointForecastPeriod) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetIsDaytime returns the IsDaytime field value if set, zero value otherwise.
func (o *GridpointForecastPeriod) GetIsDaytime() bool {
	if o == nil || o.IsDaytime == nil {
		var ret bool
		return ret
	}
	return *o.IsDaytime
}

// GetIsDaytimeOk returns a tuple with the IsDaytime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastPeriod) GetIsDaytimeOk() (*bool, bool) {
	if o == nil || o.IsDaytime == nil {
		return nil, false
	}
	return o.IsDaytime, true
}

// HasIsDaytime returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasIsDaytime() bool {
	if o != nil && o.IsDaytime != nil {
		return true
	}

	return false
}

// SetIsDaytime gets a reference to the given bool and assigns it to the IsDaytime field.
func (o *GridpointForecastPeriod) SetIsDaytime(v bool) {
	o.IsDaytime = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *GridpointForecastPeriod) GetTemperature() GridpointForecastPeriodTemperature {
	if o == nil || o.Temperature == nil {
		var ret GridpointForecastPeriodTemperature
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastPeriod) GetTemperatureOk() (*GridpointForecastPeriodTemperature, bool) {
	if o == nil || o.Temperature == nil {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasTemperature() bool {
	if o != nil && o.Temperature != nil {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given GridpointForecastPeriodTemperature and assigns it to the Temperature field.
func (o *GridpointForecastPeriod) SetTemperature(v GridpointForecastPeriodTemperature) {
	o.Temperature = &v
}

// GetTemperatureUnit returns the TemperatureUnit field value if set, zero value otherwise.
// Deprecated
func (o *GridpointForecastPeriod) GetTemperatureUnit() string {
	if o == nil || o.TemperatureUnit == nil {
		var ret string
		return ret
	}
	return *o.TemperatureUnit
}

// GetTemperatureUnitOk returns a tuple with the TemperatureUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GridpointForecastPeriod) GetTemperatureUnitOk() (*string, bool) {
	if o == nil || o.TemperatureUnit == nil {
		return nil, false
	}
	return o.TemperatureUnit, true
}

// HasTemperatureUnit returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasTemperatureUnit() bool {
	if o != nil && o.TemperatureUnit != nil {
		return true
	}

	return false
}

// SetTemperatureUnit gets a reference to the given string and assigns it to the TemperatureUnit field.
// Deprecated
func (o *GridpointForecastPeriod) SetTemperatureUnit(v string) {
	o.TemperatureUnit = &v
}

// GetTemperatureTrend returns the TemperatureTrend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GridpointForecastPeriod) GetTemperatureTrend() string {
	if o == nil || o.TemperatureTrend.Get() == nil {
		var ret string
		return ret
	}
	return *o.TemperatureTrend.Get()
}

// GetTemperatureTrendOk returns a tuple with the TemperatureTrend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GridpointForecastPeriod) GetTemperatureTrendOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TemperatureTrend.Get(), o.TemperatureTrend.IsSet()
}

// HasTemperatureTrend returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasTemperatureTrend() bool {
	if o != nil && o.TemperatureTrend.IsSet() {
		return true
	}

	return false
}

// SetTemperatureTrend gets a reference to the given NullableString and assigns it to the TemperatureTrend field.
func (o *GridpointForecastPeriod) SetTemperatureTrend(v string) {
	o.TemperatureTrend.Set(&v)
}
// SetTemperatureTrendNil sets the value for TemperatureTrend to be an explicit nil
func (o *GridpointForecastPeriod) SetTemperatureTrendNil() {
	o.TemperatureTrend.Set(nil)
}

// UnsetTemperatureTrend ensures that no value is present for TemperatureTrend, not even an explicit nil
func (o *GridpointForecastPeriod) UnsetTemperatureTrend() {
	o.TemperatureTrend.Unset()
}

// GetWindSpeed returns the WindSpeed field value if set, zero value otherwise.
func (o *GridpointForecastPeriod) GetWindSpeed() GridpointForecastPeriodWindSpeed {
	if o == nil || o.WindSpeed == nil {
		var ret GridpointForecastPeriodWindSpeed
		return ret
	}
	return *o.WindSpeed
}

// GetWindSpeedOk returns a tuple with the WindSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastPeriod) GetWindSpeedOk() (*GridpointForecastPeriodWindSpeed, bool) {
	if o == nil || o.WindSpeed == nil {
		return nil, false
	}
	return o.WindSpeed, true
}

// HasWindSpeed returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasWindSpeed() bool {
	if o != nil && o.WindSpeed != nil {
		return true
	}

	return false
}

// SetWindSpeed gets a reference to the given GridpointForecastPeriodWindSpeed and assigns it to the WindSpeed field.
func (o *GridpointForecastPeriod) SetWindSpeed(v GridpointForecastPeriodWindSpeed) {
	o.WindSpeed = &v
}

// GetWindGust returns the WindGust field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GridpointForecastPeriod) GetWindGust() GridpointForecastPeriodWindGust {
	if o == nil || o.WindGust.Get() == nil {
		var ret GridpointForecastPeriodWindGust
		return ret
	}
	return *o.WindGust.Get()
}

// GetWindGustOk returns a tuple with the WindGust field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GridpointForecastPeriod) GetWindGustOk() (*GridpointForecastPeriodWindGust, bool) {
	if o == nil {
		return nil, false
	}
	return o.WindGust.Get(), o.WindGust.IsSet()
}

// HasWindGust returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasWindGust() bool {
	if o != nil && o.WindGust.IsSet() {
		return true
	}

	return false
}

// SetWindGust gets a reference to the given NullableGridpointForecastPeriodWindGust and assigns it to the WindGust field.
func (o *GridpointForecastPeriod) SetWindGust(v GridpointForecastPeriodWindGust) {
	o.WindGust.Set(&v)
}
// SetWindGustNil sets the value for WindGust to be an explicit nil
func (o *GridpointForecastPeriod) SetWindGustNil() {
	o.WindGust.Set(nil)
}

// UnsetWindGust ensures that no value is present for WindGust, not even an explicit nil
func (o *GridpointForecastPeriod) UnsetWindGust() {
	o.WindGust.Unset()
}

// GetWindDirection returns the WindDirection field value if set, zero value otherwise.
func (o *GridpointForecastPeriod) GetWindDirection() string {
	if o == nil || o.WindDirection == nil {
		var ret string
		return ret
	}
	return *o.WindDirection
}

// GetWindDirectionOk returns a tuple with the WindDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastPeriod) GetWindDirectionOk() (*string, bool) {
	if o == nil || o.WindDirection == nil {
		return nil, false
	}
	return o.WindDirection, true
}

// HasWindDirection returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasWindDirection() bool {
	if o != nil && o.WindDirection != nil {
		return true
	}

	return false
}

// SetWindDirection gets a reference to the given string and assigns it to the WindDirection field.
func (o *GridpointForecastPeriod) SetWindDirection(v string) {
	o.WindDirection = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
// Deprecated
func (o *GridpointForecastPeriod) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GridpointForecastPeriod) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
// Deprecated
func (o *GridpointForecastPeriod) SetIcon(v string) {
	o.Icon = &v
}

// GetShortForecast returns the ShortForecast field value if set, zero value otherwise.
func (o *GridpointForecastPeriod) GetShortForecast() string {
	if o == nil || o.ShortForecast == nil {
		var ret string
		return ret
	}
	return *o.ShortForecast
}

// GetShortForecastOk returns a tuple with the ShortForecast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastPeriod) GetShortForecastOk() (*string, bool) {
	if o == nil || o.ShortForecast == nil {
		return nil, false
	}
	return o.ShortForecast, true
}

// HasShortForecast returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasShortForecast() bool {
	if o != nil && o.ShortForecast != nil {
		return true
	}

	return false
}

// SetShortForecast gets a reference to the given string and assigns it to the ShortForecast field.
func (o *GridpointForecastPeriod) SetShortForecast(v string) {
	o.ShortForecast = &v
}

// GetDetailedForecast returns the DetailedForecast field value if set, zero value otherwise.
func (o *GridpointForecastPeriod) GetDetailedForecast() string {
	if o == nil || o.DetailedForecast == nil {
		var ret string
		return ret
	}
	return *o.DetailedForecast
}

// GetDetailedForecastOk returns a tuple with the DetailedForecast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GridpointForecastPeriod) GetDetailedForecastOk() (*string, bool) {
	if o == nil || o.DetailedForecast == nil {
		return nil, false
	}
	return o.DetailedForecast, true
}

// HasDetailedForecast returns a boolean if a field has been set.
func (o *GridpointForecastPeriod) HasDetailedForecast() bool {
	if o != nil && o.DetailedForecast != nil {
		return true
	}

	return false
}

// SetDetailedForecast gets a reference to the given string and assigns it to the DetailedForecast field.
func (o *GridpointForecastPeriod) SetDetailedForecast(v string) {
	o.DetailedForecast = &v
}

func (o GridpointForecastPeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StartTime != nil {
		toSerialize["startTime"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["endTime"] = o.EndTime
	}
	if o.IsDaytime != nil {
		toSerialize["isDaytime"] = o.IsDaytime
	}
	if o.Temperature != nil {
		toSerialize["temperature"] = o.Temperature
	}
	if o.TemperatureUnit != nil {
		toSerialize["temperatureUnit"] = o.TemperatureUnit
	}
	if o.TemperatureTrend.IsSet() {
		toSerialize["temperatureTrend"] = o.TemperatureTrend.Get()
	}
	if o.WindSpeed != nil {
		toSerialize["windSpeed"] = o.WindSpeed
	}
	if o.WindGust.IsSet() {
		toSerialize["windGust"] = o.WindGust.Get()
	}
	if o.WindDirection != nil {
		toSerialize["windDirection"] = o.WindDirection
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.ShortForecast != nil {
		toSerialize["shortForecast"] = o.ShortForecast
	}
	if o.DetailedForecast != nil {
		toSerialize["detailedForecast"] = o.DetailedForecast
	}
	return json.Marshal(toSerialize)
}

type NullableGridpointForecastPeriod struct {
	value *GridpointForecastPeriod
	isSet bool
}

func (v NullableGridpointForecastPeriod) Get() *GridpointForecastPeriod {
	return v.value
}

func (v *NullableGridpointForecastPeriod) Set(val *GridpointForecastPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableGridpointForecastPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableGridpointForecastPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGridpointForecastPeriod(val *GridpointForecastPeriod) *NullableGridpointForecastPeriod {
	return &NullableGridpointForecastPeriod{value: val, isSet: true}
}

func (v NullableGridpointForecastPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGridpointForecastPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

