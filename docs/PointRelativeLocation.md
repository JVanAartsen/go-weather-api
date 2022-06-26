# PointRelativeLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Geometry** | **NullableString** | A geometry represented in Well-Known Text (WKT) format. | 
**Properties** | [**RelativeLocation**](RelativeLocation.md) |  | 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**Bearing** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 

## Methods

### NewPointRelativeLocation

`func NewPointRelativeLocation(type_ string, geometry NullableString, properties RelativeLocation, ) *PointRelativeLocation`

NewPointRelativeLocation instantiates a new PointRelativeLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointRelativeLocationWithDefaults

`func NewPointRelativeLocationWithDefaults() *PointRelativeLocation`

NewPointRelativeLocationWithDefaults instantiates a new PointRelativeLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *PointRelativeLocation) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PointRelativeLocation) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PointRelativeLocation) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *PointRelativeLocation) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *PointRelativeLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PointRelativeLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PointRelativeLocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PointRelativeLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PointRelativeLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PointRelativeLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PointRelativeLocation) SetType(v string)`

SetType sets Type field to given value.


### GetGeometry

`func (o *PointRelativeLocation) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *PointRelativeLocation) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *PointRelativeLocation) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.


### SetGeometryNil

`func (o *PointRelativeLocation) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *PointRelativeLocation) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetProperties

`func (o *PointRelativeLocation) GetProperties() RelativeLocation`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PointRelativeLocation) GetPropertiesOk() (*RelativeLocation, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PointRelativeLocation) SetProperties(v RelativeLocation)`

SetProperties sets Properties field to given value.


### GetCity

`func (o *PointRelativeLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PointRelativeLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PointRelativeLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *PointRelativeLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *PointRelativeLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PointRelativeLocation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PointRelativeLocation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PointRelativeLocation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDistance

`func (o *PointRelativeLocation) GetDistance() QuantitativeValue`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *PointRelativeLocation) GetDistanceOk() (*QuantitativeValue, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *PointRelativeLocation) SetDistance(v QuantitativeValue)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *PointRelativeLocation) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetBearing

`func (o *PointRelativeLocation) GetBearing() QuantitativeValue`

GetBearing returns the Bearing field if non-nil, zero value otherwise.

### GetBearingOk

`func (o *PointRelativeLocation) GetBearingOk() (*QuantitativeValue, bool)`

GetBearingOk returns a tuple with the Bearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearing

`func (o *PointRelativeLocation) SetBearing(v QuantitativeValue)`

SetBearing sets Bearing field to given value.

### HasBearing

`func (o *PointRelativeLocation) HasBearing() bool`

HasBearing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


