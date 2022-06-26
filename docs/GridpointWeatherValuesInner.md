# GridpointWeatherValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidTime** | [**ISO8601Interval**](ISO8601Interval.md) |  | 
**Value** | [**[]GridpointWeatherValuesInnerValueInner**](GridpointWeatherValuesInnerValueInner.md) |  | 

## Methods

### NewGridpointWeatherValuesInner

`func NewGridpointWeatherValuesInner(validTime ISO8601Interval, value []GridpointWeatherValuesInnerValueInner, ) *GridpointWeatherValuesInner`

NewGridpointWeatherValuesInner instantiates a new GridpointWeatherValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridpointWeatherValuesInnerWithDefaults

`func NewGridpointWeatherValuesInnerWithDefaults() *GridpointWeatherValuesInner`

NewGridpointWeatherValuesInnerWithDefaults instantiates a new GridpointWeatherValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidTime

`func (o *GridpointWeatherValuesInner) GetValidTime() ISO8601Interval`

GetValidTime returns the ValidTime field if non-nil, zero value otherwise.

### GetValidTimeOk

`func (o *GridpointWeatherValuesInner) GetValidTimeOk() (*ISO8601Interval, bool)`

GetValidTimeOk returns a tuple with the ValidTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTime

`func (o *GridpointWeatherValuesInner) SetValidTime(v ISO8601Interval)`

SetValidTime sets ValidTime field to given value.


### GetValue

`func (o *GridpointWeatherValuesInner) GetValue() []GridpointWeatherValuesInnerValueInner`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GridpointWeatherValuesInner) GetValueOk() (*[]GridpointWeatherValuesInnerValueInner, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GridpointWeatherValuesInner) SetValue(v []GridpointWeatherValuesInnerValueInner)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


