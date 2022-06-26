# GridpointForecastGeoJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Geometry** | [**NullableGeoJsonGeometry**](GeoJsonGeometry.md) |  | 
**Properties** | [**GridpointForecast**](GridpointForecast.md) |  | 

## Methods

### NewGridpointForecastGeoJson

`func NewGridpointForecastGeoJson(type_ string, geometry NullableGeoJsonGeometry, properties GridpointForecast, ) *GridpointForecastGeoJson`

NewGridpointForecastGeoJson instantiates a new GridpointForecastGeoJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridpointForecastGeoJsonWithDefaults

`func NewGridpointForecastGeoJsonWithDefaults() *GridpointForecastGeoJson`

NewGridpointForecastGeoJsonWithDefaults instantiates a new GridpointForecastGeoJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *GridpointForecastGeoJson) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GridpointForecastGeoJson) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GridpointForecastGeoJson) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *GridpointForecastGeoJson) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *GridpointForecastGeoJson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GridpointForecastGeoJson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GridpointForecastGeoJson) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GridpointForecastGeoJson) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GridpointForecastGeoJson) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GridpointForecastGeoJson) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GridpointForecastGeoJson) SetType(v string)`

SetType sets Type field to given value.


### GetGeometry

`func (o *GridpointForecastGeoJson) GetGeometry() GeoJsonGeometry`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *GridpointForecastGeoJson) GetGeometryOk() (*GeoJsonGeometry, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *GridpointForecastGeoJson) SetGeometry(v GeoJsonGeometry)`

SetGeometry sets Geometry field to given value.


### SetGeometryNil

`func (o *GridpointForecastGeoJson) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *GridpointForecastGeoJson) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetProperties

`func (o *GridpointForecastGeoJson) GetProperties() GridpointForecast`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GridpointForecastGeoJson) GetPropertiesOk() (*GridpointForecast, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GridpointForecastGeoJson) SetProperties(v GridpointForecast)`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


