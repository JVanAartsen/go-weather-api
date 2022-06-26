# Gridpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Geometry** | Pointer to **NullableString** | A geometry represented in Well-Known Text (WKT) format. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 
**ValidTimes** | Pointer to [**ISO8601Interval**](ISO8601Interval.md) |  | [optional] 
**Elevation** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**ForecastOffice** | Pointer to **string** |  | [optional] 
**GridId** | Pointer to **string** |  | [optional] 
**GridX** | Pointer to **int32** |  | [optional] 
**GridY** | Pointer to **int32** |  | [optional] 
**Weather** | Pointer to [**GridpointWeather**](GridpointWeather.md) |  | [optional] 
**Hazards** | Pointer to [**GridpointHazards**](GridpointHazards.md) |  | [optional] 

## Methods

### NewGridpoint

`func NewGridpoint() *Gridpoint`

NewGridpoint instantiates a new Gridpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridpointWithDefaults

`func NewGridpointWithDefaults() *Gridpoint`

NewGridpointWithDefaults instantiates a new Gridpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Gridpoint) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Gridpoint) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Gridpoint) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *Gridpoint) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetGeometry

`func (o *Gridpoint) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *Gridpoint) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *Gridpoint) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *Gridpoint) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### SetGeometryNil

`func (o *Gridpoint) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *Gridpoint) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetId

`func (o *Gridpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gridpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gridpoint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Gridpoint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Gridpoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Gridpoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Gridpoint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Gridpoint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateTime

`func (o *Gridpoint) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *Gridpoint) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *Gridpoint) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *Gridpoint) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetValidTimes

`func (o *Gridpoint) GetValidTimes() ISO8601Interval`

GetValidTimes returns the ValidTimes field if non-nil, zero value otherwise.

### GetValidTimesOk

`func (o *Gridpoint) GetValidTimesOk() (*ISO8601Interval, bool)`

GetValidTimesOk returns a tuple with the ValidTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTimes

`func (o *Gridpoint) SetValidTimes(v ISO8601Interval)`

SetValidTimes sets ValidTimes field to given value.

### HasValidTimes

`func (o *Gridpoint) HasValidTimes() bool`

HasValidTimes returns a boolean if a field has been set.

### GetElevation

`func (o *Gridpoint) GetElevation() QuantitativeValue`

GetElevation returns the Elevation field if non-nil, zero value otherwise.

### GetElevationOk

`func (o *Gridpoint) GetElevationOk() (*QuantitativeValue, bool)`

GetElevationOk returns a tuple with the Elevation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevation

`func (o *Gridpoint) SetElevation(v QuantitativeValue)`

SetElevation sets Elevation field to given value.

### HasElevation

`func (o *Gridpoint) HasElevation() bool`

HasElevation returns a boolean if a field has been set.

### GetForecastOffice

`func (o *Gridpoint) GetForecastOffice() string`

GetForecastOffice returns the ForecastOffice field if non-nil, zero value otherwise.

### GetForecastOfficeOk

`func (o *Gridpoint) GetForecastOfficeOk() (*string, bool)`

GetForecastOfficeOk returns a tuple with the ForecastOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastOffice

`func (o *Gridpoint) SetForecastOffice(v string)`

SetForecastOffice sets ForecastOffice field to given value.

### HasForecastOffice

`func (o *Gridpoint) HasForecastOffice() bool`

HasForecastOffice returns a boolean if a field has been set.

### GetGridId

`func (o *Gridpoint) GetGridId() string`

GetGridId returns the GridId field if non-nil, zero value otherwise.

### GetGridIdOk

`func (o *Gridpoint) GetGridIdOk() (*string, bool)`

GetGridIdOk returns a tuple with the GridId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridId

`func (o *Gridpoint) SetGridId(v string)`

SetGridId sets GridId field to given value.

### HasGridId

`func (o *Gridpoint) HasGridId() bool`

HasGridId returns a boolean if a field has been set.

### GetGridX

`func (o *Gridpoint) GetGridX() int32`

GetGridX returns the GridX field if non-nil, zero value otherwise.

### GetGridXOk

`func (o *Gridpoint) GetGridXOk() (*int32, bool)`

GetGridXOk returns a tuple with the GridX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridX

`func (o *Gridpoint) SetGridX(v int32)`

SetGridX sets GridX field to given value.

### HasGridX

`func (o *Gridpoint) HasGridX() bool`

HasGridX returns a boolean if a field has been set.

### GetGridY

`func (o *Gridpoint) GetGridY() int32`

GetGridY returns the GridY field if non-nil, zero value otherwise.

### GetGridYOk

`func (o *Gridpoint) GetGridYOk() (*int32, bool)`

GetGridYOk returns a tuple with the GridY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridY

`func (o *Gridpoint) SetGridY(v int32)`

SetGridY sets GridY field to given value.

### HasGridY

`func (o *Gridpoint) HasGridY() bool`

HasGridY returns a boolean if a field has been set.

### GetWeather

`func (o *Gridpoint) GetWeather() GridpointWeather`

GetWeather returns the Weather field if non-nil, zero value otherwise.

### GetWeatherOk

`func (o *Gridpoint) GetWeatherOk() (*GridpointWeather, bool)`

GetWeatherOk returns a tuple with the Weather field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeather

`func (o *Gridpoint) SetWeather(v GridpointWeather)`

SetWeather sets Weather field to given value.

### HasWeather

`func (o *Gridpoint) HasWeather() bool`

HasWeather returns a boolean if a field has been set.

### GetHazards

`func (o *Gridpoint) GetHazards() GridpointHazards`

GetHazards returns the Hazards field if non-nil, zero value otherwise.

### GetHazardsOk

`func (o *Gridpoint) GetHazardsOk() (*GridpointHazards, bool)`

GetHazardsOk returns a tuple with the Hazards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHazards

`func (o *Gridpoint) SetHazards(v GridpointHazards)`

SetHazards sets Hazards field to given value.

### HasHazards

`func (o *Gridpoint) HasHazards() bool`

HasHazards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


