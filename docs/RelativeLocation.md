# RelativeLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Distance** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**Bearing** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 

## Methods

### NewRelativeLocation

`func NewRelativeLocation() *RelativeLocation`

NewRelativeLocation instantiates a new RelativeLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelativeLocationWithDefaults

`func NewRelativeLocationWithDefaults() *RelativeLocation`

NewRelativeLocationWithDefaults instantiates a new RelativeLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *RelativeLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *RelativeLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *RelativeLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *RelativeLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *RelativeLocation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RelativeLocation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RelativeLocation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RelativeLocation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetDistance

`func (o *RelativeLocation) GetDistance() QuantitativeValue`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *RelativeLocation) GetDistanceOk() (*QuantitativeValue, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *RelativeLocation) SetDistance(v QuantitativeValue)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *RelativeLocation) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetBearing

`func (o *RelativeLocation) GetBearing() QuantitativeValue`

GetBearing returns the Bearing field if non-nil, zero value otherwise.

### GetBearingOk

`func (o *RelativeLocation) GetBearingOk() (*QuantitativeValue, bool)`

GetBearingOk returns a tuple with the Bearing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearing

`func (o *RelativeLocation) SetBearing(v QuantitativeValue)`

SetBearing sets Bearing field to given value.

### HasBearing

`func (o *RelativeLocation) HasBearing() bool`

HasBearing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


