# RelativeLocationJsonLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**Bearing** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**Geometry** | **NullableString** | A geometry represented in Well-Known Text (WKT) format. | 

## Methods

### NewRelativeLocationJsonLd

`func NewRelativeLocationJsonLd(geometry NullableString, ) *RelativeLocationJsonLd`

NewRelativeLocationJsonLd instantiates a new RelativeLocationJsonLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelativeLocationJsonLdWithDefaults

`func NewRelativeLocationJsonLdWithDefaults() *RelativeLocationJsonLd`

NewRelativeLocationJsonLdWithDefaults instantiates a new RelativeLocationJsonLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *RelativeLocationJsonLd) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RelativeLocationJsonLd) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RelativeLocationJsonLd) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *RelativeLocationJsonLd) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *RelativeLocationJsonLd) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RelativeLocationJsonLd) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RelativeLocationJsonLd) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RelativeLocationJsonLd) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDistance

`func (o *RelativeLocationJsonLd) GetDistance() QuantitativeValue`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *RelativeLocationJsonLd) GetDistanceOk() (*QuantitativeValue, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *RelativeLocationJsonLd) SetDistance(v QuantitativeValue)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *RelativeLocationJsonLd) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetBearing

`func (o *RelativeLocationJsonLd) GetBearing() QuantitativeValue`

GetBearing returns the Bearing field if non-nil, zero value otherwise.

### GetBearingOk

`func (o *RelativeLocationJsonLd) GetBearingOk() (*QuantitativeValue, bool)`

GetBearingOk returns a tuple with the Bearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearing

`func (o *RelativeLocationJsonLd) SetBearing(v QuantitativeValue)`

SetBearing sets Bearing field to given value.

### HasBearing

`func (o *RelativeLocationJsonLd) HasBearing() bool`

HasBearing returns a boolean if a field has been set.

### GetGeometry

`func (o *RelativeLocationJsonLd) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *RelativeLocationJsonLd) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *RelativeLocationJsonLd) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.


### SetGeometryNil

`func (o *RelativeLocationJsonLd) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *RelativeLocationJsonLd) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


