# ZoneForecast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Geometry** | Pointer to **NullableString** | A geometry represented in Well-Known Text (WKT) format. | [optional] 
**Zone** | Pointer to **string** | An API link to the zone this forecast is for. | [optional] 
**Updated** | Pointer to **time.Time** | The time this zone forecast product was published. | [optional] 
**Periods** | Pointer to [**[]ZoneForecastPeriodsInner**](ZoneForecastPeriodsInner.md) | An array of forecast periods. | [optional] 

## Methods

### NewZoneForecast

`func NewZoneForecast() *ZoneForecast`

NewZoneForecast instantiates a new ZoneForecast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneForecastWithDefaults

`func NewZoneForecastWithDefaults() *ZoneForecast`

NewZoneForecastWithDefaults instantiates a new ZoneForecast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ZoneForecast) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ZoneForecast) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ZoneForecast) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ZoneForecast) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetGeometry

`func (o *ZoneForecast) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *ZoneForecast) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *ZoneForecast) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *ZoneForecast) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### SetGeometryNil

`func (o *ZoneForecast) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *ZoneForecast) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetZone

`func (o *ZoneForecast) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ZoneForecast) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ZoneForecast) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ZoneForecast) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetUpdated

`func (o *ZoneForecast) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ZoneForecast) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ZoneForecast) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ZoneForecast) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetPeriods

`func (o *ZoneForecast) GetPeriods() []ZoneForecastPeriodsInner`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *ZoneForecast) GetPeriodsOk() (*[]ZoneForecastPeriodsInner, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *ZoneForecast) SetPeriods(v []ZoneForecastPeriodsInner)`

SetPeriods sets Periods field to given value.

### HasPeriods

`func (o *ZoneForecast) HasPeriods() bool`

HasPeriods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


