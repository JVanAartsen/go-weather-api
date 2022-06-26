# ZoneForecastPeriodsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **int32** | A sequential identifier number. | 
**Name** | **string** | A textual description of the period. | 
**DetailedForecast** | **string** | A detailed textual forecast for the period. | 

## Methods

### NewZoneForecastPeriodsInner

`func NewZoneForecastPeriodsInner(number int32, name string, detailedForecast string, ) *ZoneForecastPeriodsInner`

NewZoneForecastPeriodsInner instantiates a new ZoneForecastPeriodsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneForecastPeriodsInnerWithDefaults

`func NewZoneForecastPeriodsInnerWithDefaults() *ZoneForecastPeriodsInner`

NewZoneForecastPeriodsInnerWithDefaults instantiates a new ZoneForecastPeriodsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *ZoneForecastPeriodsInner) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ZoneForecastPeriodsInner) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ZoneForecastPeriodsInner) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetName

`func (o *ZoneForecastPeriodsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ZoneForecastPeriodsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ZoneForecastPeriodsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDetailedForecast

`func (o *ZoneForecastPeriodsInner) GetDetailedForecast() string`

GetDetailedForecast returns the DetailedForecast field if non-nil, zero value otherwise.

### GetDetailedForecastOk

`func (o *ZoneForecastPeriodsInner) GetDetailedForecastOk() (*string, bool)`

GetDetailedForecastOk returns a tuple with the DetailedForecast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedForecast

`func (o *ZoneForecastPeriodsInner) SetDetailedForecast(v string)`

SetDetailedForecast sets DetailedForecast field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


