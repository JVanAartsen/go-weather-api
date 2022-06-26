# GridpointForecastJsonLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**JsonLdContext**](JsonLdContext.md) |  | 
**Geometry** | **NullableString** | A geometry represented in Well-Known Text (WKT) format. | 
**Units** | Pointer to [**GridpointForecastUnits**](GridpointForecastUnits.md) |  | [optional] [default to US]
**ForecastGenerator** | Pointer to **string** | The internal generator class used to create the forecast text (used for NWS debugging). | [optional] 
**GeneratedAt** | Pointer to **time.Time** | The time this forecast data was generated. | [optional] 
**UpdateTime** | Pointer to **time.Time** | The last update time of the data this forecast was generated from. | [optional] 
**Updated** | Pointer to **time.Time** | This property is deprecated (use updateTime instead). | [optional] 
**ValidTimes** | Pointer to [**ISO8601Interval**](ISO8601Interval.md) |  | [optional] 
**Elevation** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**Periods** | Pointer to [**[]GridpointForecastPeriod**](GridpointForecastPeriod.md) | An array of forecast periods. | [optional] 

## Methods

### NewGridpointForecastJsonLd

`func NewGridpointForecastJsonLd(context JsonLdContext, geometry NullableString, ) *GridpointForecastJsonLd`

NewGridpointForecastJsonLd instantiates a new GridpointForecastJsonLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridpointForecastJsonLdWithDefaults

`func NewGridpointForecastJsonLdWithDefaults() *GridpointForecastJsonLd`

NewGridpointForecastJsonLdWithDefaults instantiates a new GridpointForecastJsonLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *GridpointForecastJsonLd) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GridpointForecastJsonLd) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GridpointForecastJsonLd) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.


### GetGeometry

`func (o *GridpointForecastJsonLd) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *GridpointForecastJsonLd) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *GridpointForecastJsonLd) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.


### SetGeometryNil

`func (o *GridpointForecastJsonLd) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *GridpointForecastJsonLd) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetUnits

`func (o *GridpointForecastJsonLd) GetUnits() GridpointForecastUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *GridpointForecastJsonLd) GetUnitsOk() (*GridpointForecastUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *GridpointForecastJsonLd) SetUnits(v GridpointForecastUnits)`

SetUnits sets Units field to given value.

### HasUnits

`func (o *GridpointForecastJsonLd) HasUnits() bool`

HasUnits returns a boolean if a field has been set.

### GetForecastGenerator

`func (o *GridpointForecastJsonLd) GetForecastGenerator() string`

GetForecastGenerator returns the ForecastGenerator field if non-nil, zero value otherwise.

### GetForecastGeneratorOk

`func (o *GridpointForecastJsonLd) GetForecastGeneratorOk() (*string, bool)`

GetForecastGeneratorOk returns a tuple with the ForecastGenerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastGenerator

`func (o *GridpointForecastJsonLd) SetForecastGenerator(v string)`

SetForecastGenerator sets ForecastGenerator field to given value.

### HasForecastGenerator

`func (o *GridpointForecastJsonLd) HasForecastGenerator() bool`

HasForecastGenerator returns a boolean if a field has been set.

### GetGeneratedAt

`func (o *GridpointForecastJsonLd) GetGeneratedAt() time.Time`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *GridpointForecastJsonLd) GetGeneratedAtOk() (*time.Time, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *GridpointForecastJsonLd) SetGeneratedAt(v time.Time)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *GridpointForecastJsonLd) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GridpointForecastJsonLd) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GridpointForecastJsonLd) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GridpointForecastJsonLd) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GridpointForecastJsonLd) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdated

`func (o *GridpointForecastJsonLd) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GridpointForecastJsonLd) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GridpointForecastJsonLd) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GridpointForecastJsonLd) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetValidTimes

`func (o *GridpointForecastJsonLd) GetValidTimes() ISO8601Interval`

GetValidTimes returns the ValidTimes field if non-nil, zero value otherwise.

### GetValidTimesOk

`func (o *GridpointForecastJsonLd) GetValidTimesOk() (*ISO8601Interval, bool)`

GetValidTimesOk returns a tuple with the ValidTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTimes

`func (o *GridpointForecastJsonLd) SetValidTimes(v ISO8601Interval)`

SetValidTimes sets ValidTimes field to given value.

### HasValidTimes

`func (o *GridpointForecastJsonLd) HasValidTimes() bool`

HasValidTimes returns a boolean if a field has been set.

### GetElevation

`func (o *GridpointForecastJsonLd) GetElevation() QuantitativeValue`

GetElevation returns the Elevation field if non-nil, zero value otherwise.

### GetElevationOk

`func (o *GridpointForecastJsonLd) GetElevationOk() (*QuantitativeValue, bool)`

GetElevationOk returns a tuple with the Elevation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevation

`func (o *GridpointForecastJsonLd) SetElevation(v QuantitativeValue)`

SetElevation sets Elevation field to given value.

### HasElevation

`func (o *GridpointForecastJsonLd) HasElevation() bool`

HasElevation returns a boolean if a field has been set.

### GetPeriods

`func (o *GridpointForecastJsonLd) GetPeriods() []GridpointForecastPeriod`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *GridpointForecastJsonLd) GetPeriodsOk() (*[]GridpointForecastPeriod, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *GridpointForecastJsonLd) SetPeriods(v []GridpointForecastPeriod)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *GridpointForecastJsonLd) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


