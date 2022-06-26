# GridpointQuantitativeValueLayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uom** | Pointer to **string** | A string denoting a unit of measure, expressed in the format \&quot;{unit}\&quot; or \&quot;{namespace}:{unit}\&quot;. Units with the namespace \&quot;wmo\&quot; or \&quot;wmoUnit\&quot; are defined in the World Meteorological Organization Codes Registry at http://codes.wmo.int/common/unit and should be canonically resolvable to http://codes.wmo.int/common/unit/{unit}. Units with the namespace \&quot;nwsUnit\&quot; are currently custom and do not align to any standard. Units with no namespace or the namespace \&quot;uc\&quot; are compliant with the Unified Code for Units of Measure syntax defined at https://unitsofmeasure.org/. This also aligns with recent versions of the Geographic Markup Language (GML) standard, the IWXXM standard, and OGC Observations and Measurements v2.0 (ISO/DIS 19156). Namespaced units are considered deprecated. We will be aligning API to use the same standards as GML/IWXXM in the future.  | [optional] 
**Values** | [**[]GridpointQuantitativeValueLayerValuesInner**](GridpointQuantitativeValueLayerValuesInner.md) |  | 

## Methods

### NewGridpointQuantitativeValueLayer

`func NewGridpointQuantitativeValueLayer(values []GridpointQuantitativeValueLayerValuesInner, ) *GridpointQuantitativeValueLayer`

NewGridpointQuantitativeValueLayer instantiates a new GridpointQuantitativeValueLayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridpointQuantitativeValueLayerWithDefaults

`func NewGridpointQuantitativeValueLayerWithDefaults() *GridpointQuantitativeValueLayer`

NewGridpointQuantitativeValueLayerWithDefaults instantiates a new GridpointQuantitativeValueLayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUom

`func (o *GridpointQuantitativeValueLayer) GetUom() string`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *GridpointQuantitativeValueLayer) GetUomOk() (*string, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *GridpointQuantitativeValueLayer) SetUom(v string)`

SetUom sets Uom field to given value.

### HasUom

`func (o *GridpointQuantitativeValueLayer) HasUom() bool`

HasUom returns a boolean if a field has been set.

### GetValues

`func (o *GridpointQuantitativeValueLayer) GetValues() []GridpointQuantitativeValueLayerValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GridpointQuantitativeValueLayer) GetValuesOk() (*[]GridpointQuantitativeValueLayerValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GridpointQuantitativeValueLayer) SetValues(v []GridpointQuantitativeValueLayerValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


