# MetarPhenomenon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Intensity** | **NullableString** |  | 
**Modifier** | **NullableString** |  | 
**Weather** | **string** |  | 
**RawString** | **string** |  | 
**InVicinity** | Pointer to **bool** |  | [optional] 

## Methods

### NewMetarPhenomenon

`func NewMetarPhenomenon(intensity NullableString, modifier NullableString, weather string, rawString string, ) *MetarPhenomenon`

NewMetarPhenomenon instantiates a new MetarPhenomenon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetarPhenomenonWithDefaults

`func NewMetarPhenomenonWithDefaults() *MetarPhenomenon`

NewMetarPhenomenonWithDefaults instantiates a new MetarPhenomenon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntensity

`func (o *MetarPhenomenon) GetIntensity() string`

GetIntensity returns the Intensity field if non-nil, zero value otherwise.

### GetIntensityOk

`func (o *MetarPhenomenon) GetIntensityOk() (*string, bool)`

GetIntensityOk returns a tuple with the Intensity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntensity

`func (o *MetarPhenomenon) SetIntensity(v string)`

SetIntensity sets Intensity field to given value.


### SetIntensityNil

`func (o *MetarPhenomenon) SetIntensityNil(b bool)`

 SetIntensityNil sets the value for Intensity to be an explicit nil

### UnsetIntensity
`func (o *MetarPhenomenon) UnsetIntensity()`

UnsetIntensity ensures that no value is present for Intensity, not even an explicit nil
### GetModifier

`func (o *MetarPhenomenon) GetModifier() string`

GetModifier returns the Modifier field if non-nil, zero value otherwise.

### GetModifierOk

`func (o *MetarPhenomenon) GetModifierOk() (*string, bool)`

GetModifierOk returns a tuple with the Modifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifier

`func (o *MetarPhenomenon) SetModifier(v string)`

SetModifier sets Modifier field to given value.


### SetModifierNil

`func (o *MetarPhenomenon) SetModifierNil(b bool)`

 SetModifierNil sets the value for Modifier to be an explicit nil

### UnsetModifier
`func (o *MetarPhenomenon) UnsetModifier()`

UnsetModifier ensures that no value is present for Modifier, not even an explicit nil
### GetWeather

`func (o *MetarPhenomenon) GetWeather() string`

GetWeather returns the Weather field if non-nil, zero value otherwise.

### GetWeatherOk

`func (o *MetarPhenomenon) GetWeatherOk() (*string, bool)`

GetWeatherOk returns a tuple with the Weather field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeather

`func (o *MetarPhenomenon) SetWeather(v string)`

SetWeather sets Weather field to given value.


### GetRawString

`func (o *MetarPhenomenon) GetRawString() string`

GetRawString returns the RawString field if non-nil, zero value otherwise.

### GetRawStringOk

`func (o *MetarPhenomenon) GetRawStringOk() (*string, bool)`

GetRawStringOk returns a tuple with the RawString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawString

`func (o *MetarPhenomenon) SetRawString(v string)`

SetRawString sets RawString field to given value.


### GetInVicinity

`func (o *MetarPhenomenon) GetInVicinity() bool`

GetInVicinity returns the InVicinity field if non-nil, zero value otherwise.

### GetInVicinityOk

`func (o *MetarPhenomenon) GetInVicinityOk() (*bool, bool)`

GetInVicinityOk returns a tuple with the InVicinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInVicinity

`func (o *MetarPhenomenon) SetInVicinity(v bool)`

SetInVicinity sets InVicinity field to given value.

### HasInVicinity

`func (o *MetarPhenomenon) HasInVicinity() bool`

HasInVicinity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


