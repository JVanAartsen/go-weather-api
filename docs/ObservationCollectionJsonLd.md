# ObservationCollectionJsonLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Graph** | Pointer to [**[]Observation**](Observation.md) |  | [optional] 

## Methods

### NewObservationCollectionJsonLd

`func NewObservationCollectionJsonLd() *ObservationCollectionJsonLd`

NewObservationCollectionJsonLd instantiates a new ObservationCollectionJsonLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationCollectionJsonLdWithDefaults

`func NewObservationCollectionJsonLdWithDefaults() *ObservationCollectionJsonLd`

NewObservationCollectionJsonLdWithDefaults instantiates a new ObservationCollectionJsonLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ObservationCollectionJsonLd) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ObservationCollectionJsonLd) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ObservationCollectionJsonLd) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ObservationCollectionJsonLd) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetGraph

`func (o *ObservationCollectionJsonLd) GetGraph() []Observation`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *ObservationCollectionJsonLd) GetGraphOk() (*[]Observation, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *ObservationCollectionJsonLd) SetGraph(v []Observation)`

SetGraph sets Graph field to given value.

### HasGraph

`func (o *ObservationCollectionJsonLd) HasGraph() bool`

HasGraph returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


