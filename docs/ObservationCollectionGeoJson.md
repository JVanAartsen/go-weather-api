# ObservationCollectionGeoJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Type** | **string** |  | 
**Features** | [**[]ObservationGeoJsonAllOf**](ObservationGeoJsonAllOf.md) |  | 

## Methods

### NewObservationCollectionGeoJson

`func NewObservationCollectionGeoJson(type_ string, features []ObservationGeoJsonAllOf, ) *ObservationCollectionGeoJson`

NewObservationCollectionGeoJson instantiates a new ObservationCollectionGeoJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationCollectionGeoJsonWithDefaults

`func NewObservationCollectionGeoJsonWithDefaults() *ObservationCollectionGeoJson`

NewObservationCollectionGeoJsonWithDefaults instantiates a new ObservationCollectionGeoJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ObservationCollectionGeoJson) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ObservationCollectionGeoJson) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ObservationCollectionGeoJson) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ObservationCollectionGeoJson) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetType

`func (o *ObservationCollectionGeoJson) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObservationCollectionGeoJson) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObservationCollectionGeoJson) SetType(v string)`

SetType sets Type field to given value.


### GetFeatures

`func (o *ObservationCollectionGeoJson) GetFeatures() []ObservationGeoJsonAllOf`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ObservationCollectionGeoJson) GetFeaturesOk() (*[]ObservationGeoJsonAllOf, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ObservationCollectionGeoJson) SetFeatures(v []ObservationGeoJsonAllOf)`

SetFeatures sets Features field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


