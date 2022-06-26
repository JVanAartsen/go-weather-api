# Zone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Geometry** | Pointer to **NullableString** | A geometry represented in Well-Known Text (WKT) format. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | UGC identifier for a NWS forecast zone or county. The first two letters will correspond to either a state code or marine area code (see #/components/schemas/StateTerritoryCode and #/components/schemas/MarineAreaCode for lists of valid letter combinations). The third letter will be Z for public/fire zone or C for county.  | [optional] 
**Type** | Pointer to [**NWSZoneType**](NWSZoneType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**State** | Pointer to [**ZoneState**](ZoneState.md) |  | [optional] 
**Cwa** | Pointer to [**[]NWSForecastOfficeId**](NWSForecastOfficeId.md) |  | [optional] 
**ForecastOffices** | Pointer to **[]string** |  | [optional] 
**TimeZone** | Pointer to **[]string** |  | [optional] 
**ObservationStations** | Pointer to **[]string** |  | [optional] 
**RadarStation** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewZone

`func NewZone() *Zone`

NewZone instantiates a new Zone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneWithDefaults

`func NewZoneWithDefaults() *Zone`

NewZoneWithDefaults instantiates a new Zone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Zone) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Zone) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Zone) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *Zone) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetGeometry

`func (o *Zone) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *Zone) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *Zone) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *Zone) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### SetGeometryNil

`func (o *Zone) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *Zone) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetId

`func (o *Zone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Zone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Zone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Zone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Zone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Zone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Zone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Zone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Zone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Zone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Zone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Zone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Zone) GetType() NWSZoneType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Zone) GetTypeOk() (*NWSZoneType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Zone) SetType(v NWSZoneType)`

SetType sets Type field to given value.

### HasType

`func (o *Zone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Zone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Zone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Zone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Zone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *Zone) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *Zone) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *Zone) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *Zone) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *Zone) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *Zone) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *Zone) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *Zone) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetState

`func (o *Zone) GetState() ZoneState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Zone) GetStateOk() (*ZoneState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Zone) SetState(v ZoneState)`

SetState sets State field to given value.

### HasState

`func (o *Zone) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCwa

`func (o *Zone) GetCwa() []NWSForecastOfficeId`

GetCwa returns the Cwa field if non-nil, zero value otherwise.

### GetCwaOk

`func (o *Zone) GetCwaOk() (*[]NWSForecastOfficeId, bool)`

GetCwaOk returns a tuple with the Cwa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwa

`func (o *Zone) SetCwa(v []NWSForecastOfficeId)`

SetCwa sets Cwa field to given value.

### HasCwa

`func (o *Zone) HasCwa() bool`

HasCwa returns a boolean if a field has been set.

### GetForecastOffices

`func (o *Zone) GetForecastOffices() []string`

GetForecastOffices returns the ForecastOffices field if non-nil, zero value otherwise.

### GetForecastOfficesOk

`func (o *Zone) GetForecastOfficesOk() (*[]string, bool)`

GetForecastOfficesOk returns a tuple with the ForecastOffices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastOffices

`func (o *Zone) SetForecastOffices(v []string)`

SetForecastOffices sets ForecastOffices field to given value.

### HasForecastOffices

`func (o *Zone) HasForecastOffices() bool`

HasForecastOffices returns a boolean if a field has been set.

### GetTimeZone

`func (o *Zone) GetTimeZone() []string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Zone) GetTimeZoneOk() (*[]string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Zone) SetTimeZone(v []string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Zone) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetObservationStations

`func (o *Zone) GetObservationStations() []string`

GetObservationStations returns the ObservationStations field if non-nil, zero value otherwise.

### GetObservationStationsOk

`func (o *Zone) GetObservationStationsOk() (*[]string, bool)`

GetObservationStationsOk returns a tuple with the ObservationStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationStations

`func (o *Zone) SetObservationStations(v []string)`

SetObservationStations sets ObservationStations field to given value.

### HasObservationStations

`func (o *Zone) HasObservationStations() bool`

HasObservationStations returns a boolean if a field has been set.

### GetRadarStation

`func (o *Zone) GetRadarStation() string`

GetRadarStation returns the RadarStation field if non-nil, zero value otherwise.

### GetRadarStationOk

`func (o *Zone) GetRadarStationOk() (*string, bool)`

GetRadarStationOk returns a tuple with the RadarStation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadarStation

`func (o *Zone) SetRadarStation(v string)`

SetRadarStation sets RadarStation field to given value.

### HasRadarStation

`func (o *Zone) HasRadarStation() bool`

HasRadarStation returns a boolean if a field has been set.

### SetRadarStationNil

`func (o *Zone) SetRadarStationNil(b bool)`

 SetRadarStationNil sets the value for RadarStation to be an explicit nil

### UnsetRadarStation
`func (o *Zone) UnsetRadarStation()`

UnsetRadarStation ensures that no value is present for RadarStation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


