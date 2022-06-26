# GridpointForecastPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Sequential period number. | [optional] 
**Name** | Pointer to **string** | A textual identifier for the period. This value will not be present for hourly forecasts.  | [optional] 
**StartTime** | Pointer to **time.Time** | The starting time that this forecast period is valid for. | [optional] 
**EndTime** | Pointer to **time.Time** | The ending time that this forecast period is valid for. | [optional] 
**IsDaytime** | Pointer to **bool** | Indicates whether this period is daytime or nighttime. | [optional] 
**Temperature** | Pointer to [**GridpointForecastPeriodTemperature**](GridpointForecastPeriodTemperature.md) |  | [optional] 
**TemperatureUnit** | Pointer to **string** | The unit of the temperature value (Fahrenheit or Celsius). This property is deprecated. Future versions will indicate the unit within the quantitative value object for the temperature property. To make use of the future standard format now, set the \&quot;forecast_temperature_qv\&quot; feature flag on the request.  | [optional] 
**TemperatureTrend** | Pointer to **NullableString** | If not null, indicates a non-diurnal temperature trend for the period (either rising temperature overnight, or falling temperature during the day)  | [optional] 
**WindSpeed** | Pointer to [**GridpointForecastPeriodWindSpeed**](GridpointForecastPeriodWindSpeed.md) |  | [optional] 
**WindGust** | Pointer to [**NullableGridpointForecastPeriodWindGust**](GridpointForecastPeriodWindGust.md) |  | [optional] 
**WindDirection** | Pointer to **string** | The prevailing direction of the wind for the period, using a 16-point compass. | [optional] 
**Icon** | Pointer to **string** | A link to an icon representing the forecast summary. | [optional] 
**ShortForecast** | Pointer to **string** | A brief textual forecast summary for the period. | [optional] 
**DetailedForecast** | Pointer to **string** | A detailed textual forecast for the period. | [optional] 

## Methods

### NewGridpointForecastPeriod

`func NewGridpointForecastPeriod() *GridpointForecastPeriod`

NewGridpointForecastPeriod instantiates a new GridpointForecastPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridpointForecastPeriodWithDefaults

`func NewGridpointForecastPeriodWithDefaults() *GridpointForecastPeriod`

NewGridpointForecastPeriodWithDefaults instantiates a new GridpointForecastPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *GridpointForecastPeriod) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *GridpointForecastPeriod) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *GridpointForecastPeriod) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *GridpointForecastPeriod) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetName

`func (o *GridpointForecastPeriod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GridpointForecastPeriod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GridpointForecastPeriod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GridpointForecastPeriod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *GridpointForecastPeriod) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GridpointForecastPeriod) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GridpointForecastPeriod) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GridpointForecastPeriod) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *GridpointForecastPeriod) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *GridpointForecastPeriod) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *GridpointForecastPeriod) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *GridpointForecastPeriod) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetIsDaytime

`func (o *GridpointForecastPeriod) GetIsDaytime() bool`

GetIsDaytime returns the IsDaytime field if non-nil, zero value otherwise.

### GetIsDaytimeOk

`func (o *GridpointForecastPeriod) GetIsDaytimeOk() (*bool, bool)`

GetIsDaytimeOk returns a tuple with the IsDaytime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDaytime

`func (o *GridpointForecastPeriod) SetIsDaytime(v bool)`

SetIsDaytime sets IsDaytime field to given value.

### HasIsDaytime

`func (o *GridpointForecastPeriod) HasIsDaytime() bool`

HasIsDaytime returns a boolean if a field has been set.

### GetTemperature

`func (o *GridpointForecastPeriod) GetTemperature() GridpointForecastPeriodTemperature`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *GridpointForecastPeriod) GetTemperatureOk() (*GridpointForecastPeriodTemperature, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *GridpointForecastPeriod) SetTemperature(v GridpointForecastPeriodTemperature)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *GridpointForecastPeriod) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTemperatureUnit

`func (o *GridpointForecastPeriod) GetTemperatureUnit() string`

GetTemperatureUnit returns the TemperatureUnit field if non-nil, zero value otherwise.

### GetTemperatureUnitOk

`func (o *GridpointForecastPeriod) GetTemperatureUnitOk() (*string, bool)`

GetTemperatureUnitOk returns a tuple with the TemperatureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureUnit

`func (o *GridpointForecastPeriod) SetTemperatureUnit(v string)`

SetTemperatureUnit sets TemperatureUnit field to given value.

### HasTemperatureUnit

`func (o *GridpointForecastPeriod) HasTemperatureUnit() bool`

HasTemperatureUnit returns a boolean if a field has been set.

### GetTemperatureTrend

`func (o *GridpointForecastPeriod) GetTemperatureTrend() string`

GetTemperatureTrend returns the TemperatureTrend field if non-nil, zero value otherwise.

### GetTemperatureTrendOk

`func (o *GridpointForecastPeriod) GetTemperatureTrendOk() (*string, bool)`

GetTemperatureTrendOk returns a tuple with the TemperatureTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperatureTrend

`func (o *GridpointForecastPeriod) SetTemperatureTrend(v string)`

SetTemperatureTrend sets TemperatureTrend field to given value.

### HasTemperatureTrend

`func (o *GridpointForecastPeriod) HasTemperatureTrend() bool`

HasTemperatureTrend returns a boolean if a field has been set.

### SetTemperatureTrendNil

`func (o *GridpointForecastPeriod) SetTemperatureTrendNil(b bool)`

 SetTemperatureTrendNil sets the value for TemperatureTrend to be an explicit nil

### UnsetTemperatureTrend
`func (o *GridpointForecastPeriod) UnsetTemperatureTrend()`

UnsetTemperatureTrend ensures that no value is present for TemperatureTrend, not even an explicit nil
### GetWindSpeed

`func (o *GridpointForecastPeriod) GetWindSpeed() GridpointForecastPeriodWindSpeed`

GetWindSpeed returns the WindSpeed field if non-nil, zero value otherwise.

### GetWindSpeedOk

`func (o *GridpointForecastPeriod) GetWindSpeedOk() (*GridpointForecastPeriodWindSpeed, bool)`

GetWindSpeedOk returns a tuple with the WindSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindSpeed

`func (o *GridpointForecastPeriod) SetWindSpeed(v GridpointForecastPeriodWindSpeed)`

SetWindSpeed sets WindSpeed field to given value.

### HasWindSpeed

`func (o *GridpointForecastPeriod) HasWindSpeed() bool`

HasWindSpeed returns a boolean if a field has been set.

### GetWindGust

`func (o *GridpointForecastPeriod) GetWindGust() GridpointForecastPeriodWindGust`

GetWindGust returns the WindGust field if non-nil, zero value otherwise.

### GetWindGustOk

`func (o *GridpointForecastPeriod) GetWindGustOk() (*GridpointForecastPeriodWindGust, bool)`

GetWindGustOk returns a tuple with the WindGust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindGust

`func (o *GridpointForecastPeriod) SetWindGust(v GridpointForecastPeriodWindGust)`

SetWindGust sets WindGust field to given value.

### HasWindGust

`func (o *GridpointForecastPeriod) HasWindGust() bool`

HasWindGust returns a boolean if a field has been set.

### SetWindGustNil

`func (o *GridpointForecastPeriod) SetWindGustNil(b bool)`

 SetWindGustNil sets the value for WindGust to be an explicit nil

### UnsetWindGust
`func (o *GridpointForecastPeriod) UnsetWindGust()`

UnsetWindGust ensures that no value is present for WindGust, not even an explicit nil
### GetWindDirection

`func (o *GridpointForecastPeriod) GetWindDirection() string`

GetWindDirection returns the WindDirection field if non-nil, zero value otherwise.

### GetWindDirectionOk

`func (o *GridpointForecastPeriod) GetWindDirectionOk() (*string, bool)`

GetWindDirectionOk returns a tuple with the WindDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindDirection

`func (o *GridpointForecastPeriod) SetWindDirection(v string)`

SetWindDirection sets WindDirection field to given value.

### HasWindDirection

`func (o *GridpointForecastPeriod) HasWindDirection() bool`

HasWindDirection returns a boolean if a field has been set.

### GetIcon

`func (o *GridpointForecastPeriod) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *GridpointForecastPeriod) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *GridpointForecastPeriod) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *GridpointForecastPeriod) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetShortForecast

`func (o *GridpointForecastPeriod) GetShortForecast() string`

GetShortForecast returns the ShortForecast field if non-nil, zero value otherwise.

### GetShortForecastOk

`func (o *GridpointForecastPeriod) GetShortForecastOk() (*string, bool)`

GetShortForecastOk returns a tuple with the ShortForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortForecast

`func (o *GridpointForecastPeriod) SetShortForecast(v string)`

SetShortForecast sets ShortForecast field to given value.

### HasShortForecast

`func (o *GridpointForecastPeriod) HasShortForecast() bool`

HasShortForecast returns a boolean if a field has been set.

### GetDetailedForecast

`func (o *GridpointForecastPeriod) GetDetailedForecast() string`

GetDetailedForecast returns the DetailedForecast field if non-nil, zero value otherwise.

### GetDetailedForecastOk

`func (o *GridpointForecastPeriod) GetDetailedForecastOk() (*string, bool)`

GetDetailedForecastOk returns a tuple with the DetailedForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedForecast

`func (o *GridpointForecastPeriod) SetDetailedForecast(v string)`

SetDetailedForecast sets DetailedForecast field to given value.

### HasDetailedForecast

`func (o *GridpointForecastPeriod) HasDetailedForecast() bool`

HasDetailedForecast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


