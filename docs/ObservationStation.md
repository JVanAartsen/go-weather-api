# ObservationStation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Geometry** | Pointer to **NullableString** | A geometry represented in Well-Known Text (WKT) format. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Elevation** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**StationIdentifier** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**Forecast** | Pointer to **string** | A link to the NWS public forecast zone containing this station. | [optional] 
**County** | Pointer to **string** | A link to the NWS county zone containing this station. | [optional] 
**FireWeatherZone** | Pointer to **string** | A link to the NWS fire weather forecast zone containing this station. | [optional] 

## Methods

### NewObservationStation

`func NewObservationStation() *ObservationStation`

NewObservationStation instantiates a new ObservationStation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationStationWithDefaults

`func NewObservationStationWithDefaults() *ObservationStation`

NewObservationStationWithDefaults instantiates a new ObservationStation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ObservationStation) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ObservationStation) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ObservationStation) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ObservationStation) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetGeometry

`func (o *ObservationStation) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *ObservationStation) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *ObservationStation) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *ObservationStation) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### SetGeometryNil

`func (o *ObservationStation) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *ObservationStation) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetId

`func (o *ObservationStation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObservationStation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObservationStation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ObservationStation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ObservationStation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObservationStation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObservationStation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ObservationStation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetElevation

`func (o *ObservationStation) GetElevation() QuantitativeValue`

GetElevation returns the Elevation field if non-nil, zero value otherwise.

### GetElevationOk

`func (o *ObservationStation) GetElevationOk() (*QuantitativeValue, bool)`

GetElevationOk returns a tuple with the Elevation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevation

`func (o *ObservationStation) SetElevation(v QuantitativeValue)`

SetElevation sets Elevation field to given value.

### HasElevation

`func (o *ObservationStation) HasElevation() bool`

HasElevation returns a boolean if a field has been set.

### GetStationIdentifier

`func (o *ObservationStation) GetStationIdentifier() string`

GetStationIdentifier returns the StationIdentifier field if non-nil, zero value otherwise.

### GetStationIdentifierOk

`func (o *ObservationStation) GetStationIdentifierOk() (*string, bool)`

GetStationIdentifierOk returns a tuple with the StationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationIdentifier

`func (o *ObservationStation) SetStationIdentifier(v string)`

SetStationIdentifier sets StationIdentifier field to given value.

### HasStationIdentifier

`func (o *ObservationStation) HasStationIdentifier() bool`

HasStationIdentifier returns a boolean if a field has been set.

### GetName

`func (o *ObservationStation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ObservationStation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ObservationStation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ObservationStation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeZone

`func (o *ObservationStation) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ObservationStation) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ObservationStation) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ObservationStation) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetForecast

`func (o *ObservationStation) GetForecast() string`

GetForecast returns the Forecast field if non-nil, zero value otherwise.

### GetForecastOk

`func (o *ObservationStation) GetForecastOk() (*string, bool)`

GetForecastOk returns a tuple with the Forecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecast

`func (o *ObservationStation) SetForecast(v string)`

SetForecast sets Forecast field to given value.

### HasForecast

`func (o *ObservationStation) HasForecast() bool`

HasForecast returns a boolean if a field has been set.

### GetCounty

`func (o *ObservationStation) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *ObservationStation) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *ObservationStation) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *ObservationStation) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetFireWeatherZone

`func (o *ObservationStation) GetFireWeatherZone() string`

GetFireWeatherZone returns the FireWeatherZone field if non-nil, zero value otherwise.

### GetFireWeatherZoneOk

`func (o *ObservationStation) GetFireWeatherZoneOk() (*string, bool)`

GetFireWeatherZoneOk returns a tuple with the FireWeatherZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFireWeatherZone

`func (o *ObservationStation) SetFireWeatherZone(v string)`

SetFireWeatherZone sets FireWeatherZone field to given value.

### HasFireWeatherZone

`func (o *ObservationStation) HasFireWeatherZone() bool`

HasFireWeatherZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


