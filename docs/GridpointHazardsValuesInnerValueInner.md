# GridpointHazardsValuesInnerValueInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phenomenon** | **string** | Hazard code. This value will correspond to a P-VTEC phenomenon code as defined in NWS Directive 10-1703.  | 
**Significance** | **string** | Significance code. This value will correspond to a P-VTEC significance code as defined in NWS Directive 10-1703. This will most frequently be \&quot;A\&quot; for a watch or \&quot;Y\&quot; for an advisory.  | 
**EventNumber** | **NullableInt32** | Event number. If this hazard refers to a national or regional center product (such as a Storm Prediction Center convective watch), this value will be the sequence number of that product.  | 

## Methods

### NewGridpointHazardsValuesInnerValueInner

`func NewGridpointHazardsValuesInnerValueInner(phenomenon string, significance string, eventNumber NullableInt32, ) *GridpointHazardsValuesInnerValueInner`

NewGridpointHazardsValuesInnerValueInner instantiates a new GridpointHazardsValuesInnerValueInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridpointHazardsValuesInnerValueInnerWithDefaults

`func NewGridpointHazardsValuesInnerValueInnerWithDefaults() *GridpointHazardsValuesInnerValueInner`

NewGridpointHazardsValuesInnerValueInnerWithDefaults instantiates a new GridpointHazardsValuesInnerValueInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhenomenon

`func (o *GridpointHazardsValuesInnerValueInner) GetPhenomenon() string`

GetPhenomenon returns the Phenomenon field if non-nil, zero value otherwise.

### GetPhenomenonOk

`func (o *GridpointHazardsValuesInnerValueInner) GetPhenomenonOk() (*string, bool)`

GetPhenomenonOk returns a tuple with the Phenomenon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhenomenon

`func (o *GridpointHazardsValuesInnerValueInner) SetPhenomenon(v string)`

SetPhenomenon sets Phenomenon field to given value.


### GetSignificance

`func (o *GridpointHazardsValuesInnerValueInner) GetSignificance() string`

GetSignificance returns the Significance field if non-nil, zero value otherwise.

### GetSignificanceOk

`func (o *GridpointHazardsValuesInnerValueInner) GetSignificanceOk() (*string, bool)`

GetSignificanceOk returns a tuple with the Significance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignificance

`func (o *GridpointHazardsValuesInnerValueInner) SetSignificance(v string)`

SetSignificance sets Significance field to given value.


### GetEventNumber

`func (o *GridpointHazardsValuesInnerValueInner) GetEventNumber() int32`

GetEventNumber returns the EventNumber field if non-nil, zero value otherwise.

### GetEventNumberOk

`func (o *GridpointHazardsValuesInnerValueInner) GetEventNumberOk() (*int32, bool)`

GetEventNumberOk returns a tuple with the EventNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventNumber

`func (o *GridpointHazardsValuesInnerValueInner) SetEventNumber(v int32)`

SetEventNumber sets EventNumber field to given value.


### SetEventNumberNil

`func (o *GridpointHazardsValuesInnerValueInner) SetEventNumberNil(b bool)`

 SetEventNumberNil sets the value for EventNumber to be an explicit nil

### UnsetEventNumber
`func (o *GridpointHazardsValuesInnerValueInner) UnsetEventNumber()`

UnsetEventNumber ensures that no value is present for EventNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


