# GridpointForecast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Geometry** | Pointer to **NullableString** | A geometry represented in Well-Known Text (WKT) format. | [optional] 
**Units** | Pointer to [**GridpointForecastUnits**](GridpointForecastUnits.md) |  | [optional] [default to US]
**ForecastGenerator** | Pointer to **string** | The internal generator class used to create the forecast text (used for NWS debugging). | [optional] 
**GeneratedAt** | Pointer to **time.Time** | The time this forecast data was generated. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The last update time of the data this forecast was generated from. | [optional] 
**Updated** | Pointer to **time.Time** | This property is deprecated (use updateTime instead). | [optional] 
**ValidTimes** | Pointer to [**ISO8601Interval**](ISO8601Interval.md) |  | [optional] 
**Elevation** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**Periods** | Pointer to [**[]GridpointForecastPeriod**](GridpointForecastPeriod.md) | An array of forecast periods. | [optional] 

## Methods

### NewGridpointForecast

`func NewGridpointForecast() *GridpointForecast`

NewGridpointForecast instantiates a new GridpointForecast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridpointForecastWithDefaults

`func NewGridpointForecastWithDefaults() *GridpointForecast`

NewGridpointForecastWithDefaults instantiates a new GridpointForecast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *GridpointForecast) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GridpointForecast) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GridpointForecast) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *GridpointForecast) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetGeometry

`func (o *GridpointForecast) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *GridpointForecast) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *GridpointForecast) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *GridpointForecast) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### SetGeometryNil

`func (o *GridpointForecast) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *GridpointForecast) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetUnits

`func (o *GridpointForecast) GetUnits() GridpointForecastUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *GridpointForecast) GetUnitsOk() (*GridpointForecastUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *GridpointForecast) SetUnits(v GridpointForecastUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *GridpointForecast) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetForecastGenerator

`func (o *GridpointForecast) GetForecastGenerator() string`

GetForecastGenerator returns the ForecastGenerator field if non-nil, zero value otherwise.

### GetForecastGeneratorOk

`func (o *GridpointForecast) GetForecastGeneratorOk() (*string, bool)`

GetForecastGeneratorOk returns a tuple with the ForecastGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastGenerator

`func (o *GridpointForecast) SetForecastGenerator(v string)`

SetForecastGenerator sets ForecastGenerator field to given value.

### HasForecastGenerator

`func (o *GridpointForecast) HasForecastGenerator() bool`

HasForecastGenerator returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *GridpointForecast) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *GridpointForecast) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *GridpointForecast) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *GridpointForecast) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GridpointForecast) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GridpointForecast) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GridpointForecast) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GridpointForecast) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdated

`func (o *GridpointForecast) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GridpointForecast) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GridpointForecast) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GridpointForecast) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetValidTimes

`func (o *GridpointForecast) GetValidTimes() ISO8601Interval`

GetValidTimes returns the ValidTimes field if non-nil, zero value otherwise.

### GetValidTimesOk

`func (o *GridpointForecast) GetValidTimesOk() (*ISO8601Interval, bool)`

GetValidTimesOk returns a tuple with the ValidTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTimes

`func (o *GridpointForecast) SetValidTimes(v ISO8601Interval)`

SetValidTimes sets ValidTimes field to given value.

### HasValidTimes

`func (o *GridpointForecast) HasValidTimes() bool`

HasValidTimes returns a boolean if a field has been set.

### GetElevation

`func (o *GridpointForecast) GetElevation() QuantitativeValue`

GetElevation returns the Elevation field if non-nil, zero value otherwise.

### GetElevationOk

`func (o *GridpointForecast) GetElevationOk() (*QuantitativeValue, bool)`

GetElevationOk returns a tuple with the Elevation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevation

`func (o *GridpointForecast) SetElevation(v QuantitativeValue)`

SetElevation sets Elevation field to given value.

### HasElevation

`func (o *GridpointForecast) HasElevation() bool`

HasElevation returns a boolean if a field has been set.

### GetPeriods

`func (o *GridpointForecast) GetPeriods() []GridpointForecastPeriod`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *GridpointForecast) GetPeriodsOk() (*[]GridpointForecastPeriod, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *GridpointForecast) SetPeriods(v []GridpointForecastPeriod)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *GridpointForecast) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


