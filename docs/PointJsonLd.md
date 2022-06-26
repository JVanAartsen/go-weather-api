# PointJsonLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | [**JsonLdContext**](JsonLdContext.md) |  | 
**Geometry** | **NullableString** | A geometry represented in Well-Known Text (WKT) format. | 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Cwa** | Pointer to [**NWSForecastOfficeId**](NWSForecastOfficeId.md) |  | [optional] 
**ForecastOffice** | Pointer to **string** |  | [optional] 
**GridId** | Pointer to [**NWSForecastOfficeId**](NWSForecastOfficeId.md) |  | [optional] 
**GridX** | Pointer to **int32** |  | [optional] 
**GridY** | Pointer to **int32** |  | [optional] 
**Forecast** | Pointer to **string** |  | [optional] 
**ForecastHourly** | Pointer to **string** |  | [optional] 
**ForecastGridData** | Pointer to **string** |  | [optional] 
**ObservationStations** | Pointer to **string** |  | [optional] 
**RelativeLocation** | Pointer to [**PointRelativeLocation**](PointRelativeLocation.md) |  | [optional] 
**ForecastZone** | Pointer to **string** |  | [optional] 
**County** | Pointer to **string** |  | [optional] 
**FireWeatherZone** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** |  | [optional] 
**RadarStation** | Pointer to **string** |  | [optional] 

## Methods

### NewPointJsonLd

`func NewPointJsonLd(context JsonLdContext, geometry NullableString, ) *PointJsonLd`

NewPointJsonLd instantiates a new PointJsonLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointJsonLdWithDefaults

`func NewPointJsonLdWithDefaults() *PointJsonLd`

NewPointJsonLdWithDefaults instantiates a new PointJsonLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *PointJsonLd) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PointJsonLd) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PointJsonLd) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.


### GetGeometry

`func (o *PointJsonLd) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *PointJsonLd) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *PointJsonLd) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.


### SetGeometryNil

`func (o *PointJsonLd) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *PointJsonLd) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetId

`func (o *PointJsonLd) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PointJsonLd) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PointJsonLd) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PointJsonLd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PointJsonLd) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PointJsonLd) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PointJsonLd) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PointJsonLd) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCwa

`func (o *PointJsonLd) GetCwa() NWSForecastOfficeId`

GetCwa returns the Cwa field if non-nil, zero value otherwise.

### GetCwaOk

`func (o *PointJsonLd) GetCwaOk() (*NWSForecastOfficeId, bool)`

GetCwaOk returns a tuple with the Cwa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwa

`func (o *PointJsonLd) SetCwa(v NWSForecastOfficeId)`

SetCwa sets Cwa field to given value.

### HasCwa

`func (o *PointJsonLd) HasCwa() bool`

HasCwa returns a boolean if a field has been set.

### GetForecastOffice

`func (o *PointJsonLd) GetForecastOffice() string`

GetForecastOffice returns the ForecastOffice field if non-nil, zero value otherwise.

### GetForecastOfficeOk

`func (o *PointJsonLd) GetForecastOfficeOk() (*string, bool)`

GetForecastOfficeOk returns a tuple with the ForecastOffice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastOffice

`func (o *PointJsonLd) SetForecastOffice(v string)`

SetForecastOffice sets ForecastOffice field to given value.

### HasForecastOffice

`func (o *PointJsonLd) HasForecastOffice() bool`

HasForecastOffice returns a boolean if a field has been set.

### GetGridId

`func (o *PointJsonLd) GetGridId() NWSForecastOfficeId`

GetGridId returns the GridId field if non-nil, zero value otherwise.

### GetGridIdOk

`func (o *PointJsonLd) GetGridIdOk() (*NWSForecastOfficeId, bool)`

GetGridIdOk returns a tuple with the GridId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridId

`func (o *PointJsonLd) SetGridId(v NWSForecastOfficeId)`

SetGridId sets GridId field to given value.

### HasGridId

`func (o *PointJsonLd) HasGridId() bool`

HasGridId returns a boolean if a field has been set.

### GetGridX

`func (o *PointJsonLd) GetGridX() int32`

GetGridX returns the GridX field if non-nil, zero value otherwise.

### GetGridXOk

`func (o *PointJsonLd) GetGridXOk() (*int32, bool)`

GetGridXOk returns a tuple with the GridX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridX

`func (o *PointJsonLd) SetGridX(v int32)`

SetGridX sets GridX field to given value.

### HasGridX

`func (o *PointJsonLd) HasGridX() bool`

HasGridX returns a boolean if a field has been set.

### GetGridY

`func (o *PointJsonLd) GetGridY() int32`

GetGridY returns the GridY field if non-nil, zero value otherwise.

### GetGridYOk

`func (o *PointJsonLd) GetGridYOk() (*int32, bool)`

GetGridYOk returns a tuple with the GridY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridY

`func (o *PointJsonLd) SetGridY(v int32)`

SetGridY sets GridY field to given value.

### HasGridY

`func (o *PointJsonLd) HasGridY() bool`

HasGridY returns a boolean if a field has been set.

### GetForecast

`func (o *PointJsonLd) GetForecast() string`

GetForecast returns the Forecast field if non-nil, zero value otherwise.

### GetForecastOk

`func (o *PointJsonLd) GetForecastOk() (*string, bool)`

GetForecastOk returns a tuple with the Forecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecast

`func (o *PointJsonLd) SetForecast(v string)`

SetForecast sets Forecast field to given value.

### HasForecast

`func (o *PointJsonLd) HasForecast() bool`

HasForecast returns a boolean if a field has been set.

### GetForecastHourly

`func (o *PointJsonLd) GetForecastHourly() string`

GetForecastHourly returns the ForecastHourly field if non-nil, zero value otherwise.

### GetForecastHourlyOk

`func (o *PointJsonLd) GetForecastHourlyOk() (*string, bool)`

GetForecastHourlyOk returns a tuple with the ForecastHourly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastHourly

`func (o *PointJsonLd) SetForecastHourly(v string)`

SetForecastHourly sets ForecastHourly field to given value.

### HasForecastHourly

`func (o *PointJsonLd) HasForecastHourly() bool`

HasForecastHourly returns a boolean if a field has been set.

### GetForecastGridData

`func (o *PointJsonLd) GetForecastGridData() string`

GetForecastGridData returns the ForecastGridData field if non-nil, zero value otherwise.

### GetForecastGridDataOk

`func (o *PointJsonLd) GetForecastGridDataOk() (*string, bool)`

GetForecastGridDataOk returns a tuple with the ForecastGridData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastGridData

`func (o *PointJsonLd) SetForecastGridData(v string)`

SetForecastGridData sets ForecastGridData field to given value.

### HasForecastGridData

`func (o *PointJsonLd) HasForecastGridData() bool`

HasForecastGridData returns a boolean if a field has been set.

### GetObservationStations

`func (o *PointJsonLd) GetObservationStations() string`

GetObservationStations returns the ObservationStations field if non-nil, zero value otherwise.

### GetObservationStationsOk

`func (o *PointJsonLd) GetObservationStationsOk() (*string, bool)`

GetObservationStationsOk returns a tuple with the ObservationStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationStations

`func (o *PointJsonLd) SetObservationStations(v string)`

SetObservationStations sets ObservationStations field to given value.

### HasObservationStations

`func (o *PointJsonLd) HasObservationStations() bool`

HasObservationStations returns a boolean if a field has been set.

### GetRelativeLocation

`func (o *PointJsonLd) GetRelativeLocation() PointRelativeLocation`

GetRelativeLocation returns the RelativeLocation field if non-nil, zero value otherwise.

### GetRelativeLocationOk

`func (o *PointJsonLd) GetRelativeLocationOk() (*PointRelativeLocation, bool)`

GetRelativeLocationOk returns a tuple with the RelativeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeLocation

`func (o *PointJsonLd) SetRelativeLocation(v PointRelativeLocation)`

SetRelativeLocation sets RelativeLocation field to given value.

### HasRelativeLocation

`func (o *PointJsonLd) HasRelativeLocation() bool`

HasRelativeLocation returns a boolean if a field has been set.

### GetForecastZone

`func (o *PointJsonLd) GetForecastZone() string`

GetForecastZone returns the ForecastZone field if non-nil, zero value otherwise.

### GetForecastZoneOk

`func (o *PointJsonLd) GetForecastZoneOk() (*string, bool)`

GetForecastZoneOk returns a tuple with the ForecastZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastZone

`func (o *PointJsonLd) SetForecastZone(v string)`

SetForecastZone sets ForecastZone field to given value.

### HasForecastZone

`func (o *PointJsonLd) HasForecastZone() bool`

HasForecastZone returns a boolean if a field has been set.

### GetCounty

`func (o *PointJsonLd) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *PointJsonLd) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *PointJsonLd) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *PointJsonLd) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### GetFireWeatherZone

`func (o *PointJsonLd) GetFireWeatherZone() string`

GetFireWeatherZone returns the FireWeatherZone field if non-nil, zero value otherwise.

### GetFireWeatherZoneOk

`func (o *PointJsonLd) GetFireWeatherZoneOk() (*string, bool)`

GetFireWeatherZoneOk returns a tuple with the FireWeatherZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFireWeatherZone

`func (o *PointJsonLd) SetFireWeatherZone(v string)`

SetFireWeatherZone sets FireWeatherZone field to given value.

### HasFireWeatherZone

`func (o *PointJsonLd) HasFireWeatherZone() bool`

HasFireWeatherZone returns a boolean if a field has been set.

### GetTimeZone

`func (o *PointJsonLd) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PointJsonLd) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PointJsonLd) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PointJsonLd) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetRadarStation

`func (o *PointJsonLd) GetRadarStation() string`

GetRadarStation returns the RadarStation field if non-nil, zero value otherwise.

### GetRadarStationOk

`func (o *PointJsonLd) GetRadarStationOk() (*string, bool)`

GetRadarStationOk returns a tuple with the RadarStation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadarStation

`func (o *PointJsonLd) SetRadarStation(v string)`

SetRadarStation sets RadarStation field to given value.

### HasRadarStation

`func (o *PointJsonLd) HasRadarStation() bool`

HasRadarStation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


