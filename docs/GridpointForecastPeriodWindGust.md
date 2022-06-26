# GridpointForecastPeriodWindGust

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableFloat32** | A measured value | [optional] 
**MaxValue** | Pointer to **float32** | The maximum value of a range of measured values | [optional] 
**MinValue** | Pointer to **float32** | The minimum value of a range of measured values | [optional] 
**UnitCode** | Pointer to **string** | A string denoting a unit of measure, expressed in the format \&quot;{unit}\&quot; or \&quot;{namespace}:{unit}\&quot;. Units with the namespace \&quot;wmo\&quot; or \&quot;wmoUnit\&quot; are defined in the World Meteorological Organization Codes Registry at http://codes.wmo.int/common/unit and should be canonically resolvable to http://codes.wmo.int/common/unit/{unit}. Units with the namespace \&quot;nwsUnit\&quot; are currently custom and do not align to any standard. Units with no namespace or the namespace \&quot;uc\&quot; are compliant with the Unified Code for Units of Measure syntax defined at https://unitsofmeasure.org/. This also aligns with recent versions of the Geographic Markup Language (GML) standard, the IWXXM standard, and OGC Observations and Measurements v2.0 (ISO/DIS 19156). Namespaced units are considered deprecated. We will be aligning API to use the same standards as GML/IWXXM in the future.  | [optional] 
**QualityControl** | Pointer to **string** | For values in observation records, the quality control flag from the MADIS system. The definitions of these flags can be found at https://madis.ncep.noaa.gov/madis_sfc_qc_notes.shtml  | [optional] 

## Methods

### NewGridpointForecastPeriodWindGust

`func NewGridpointForecastPeriodWindGust() *GridpointForecastPeriodWindGust`

NewGridpointForecastPeriodWindGust instantiates a new GridpointForecastPeriodWindGust object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridpointForecastPeriodWindGustWithDefaults

`func NewGridpointForecastPeriodWindGustWithDefaults() *GridpointForecastPeriodWindGust`

NewGridpointForecastPeriodWindGustWithDefaults instantiates a new GridpointForecastPeriodWindGust object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *GridpointForecastPeriodWindGust) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GridpointForecastPeriodWindGust) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GridpointForecastPeriodWindGust) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *GridpointForecastPeriodWindGust) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *GridpointForecastPeriodWindGust) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *GridpointForecastPeriodWindGust) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMaxValue

`func (o *GridpointForecastPeriodWindGust) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *GridpointForecastPeriodWindGust) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *GridpointForecastPeriodWindGust) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *GridpointForecastPeriodWindGust) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetMinValue

`func (o *GridpointForecastPeriodWindGust) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *GridpointForecastPeriodWindGust) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *GridpointForecastPeriodWindGust) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *GridpointForecastPeriodWindGust) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetUnitCode

`func (o *GridpointForecastPeriodWindGust) GetUnitCode() string`

GetUnitCode returns the UnitCode field if non-nil, zero value otherwise.

### GetUnitCodeOk

`func (o *GridpointForecastPeriodWindGust) GetUnitCodeOk() (*string, bool)`

GetUnitCodeOk returns a tuple with the UnitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCode

`func (o *GridpointForecastPeriodWindGust) SetUnitCode(v string)`

SetUnitCode sets UnitCode field to given value.

### HasUnitCode

`func (o *GridpointForecastPeriodWindGust) HasUnitCode() bool`

HasUnitCode returns a boolean if a field has been set.

### GetQualityControl

`func (o *GridpointForecastPeriodWindGust) GetQualityControl() string`

GetQualityControl returns the QualityControl field if non-nil, zero value otherwise.

### GetQualityControlOk

`func (o *GridpointForecastPeriodWindGust) GetQualityControlOk() (*string, bool)`

GetQualityControlOk returns a tuple with the QualityControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityControl

`func (o *GridpointForecastPeriodWindGust) SetQualityControl(v string)`

SetQualityControl sets QualityControl field to given value.

### HasQualityControl

`func (o *GridpointForecastPeriodWindGust) HasQualityControl() bool`

HasQualityControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


