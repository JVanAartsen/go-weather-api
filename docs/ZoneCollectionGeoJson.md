# ZoneCollectionGeoJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Type** | **string** |  | 
**Features** | [**[]ZoneGeoJsonAllOf**](ZoneGeoJsonAllOf.md) |  | 

## Methods

### NewZoneCollectionGeoJson

`func NewZoneCollectionGeoJson(type_ string, features []ZoneGeoJsonAllOf, ) *ZoneCollectionGeoJson`

NewZoneCollectionGeoJson instantiates a new ZoneCollectionGeoJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCollectionGeoJsonWithDefaults

`func NewZoneCollectionGeoJsonWithDefaults() *ZoneCollectionGeoJson`

NewZoneCollectionGeoJsonWithDefaults instantiates a new ZoneCollectionGeoJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ZoneCollectionGeoJson) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ZoneCollectionGeoJson) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ZoneCollectionGeoJson) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ZoneCollectionGeoJson) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetType

`func (o *ZoneCollectionGeoJson) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneCollectionGeoJson) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneCollectionGeoJson) SetType(v string)`

SetType sets Type field to given value.


### GetFeatures

`func (o *ZoneCollectionGeoJson) GetFeatures() []ZoneGeoJsonAllOf`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ZoneCollectionGeoJson) GetFeaturesOk() (*[]ZoneGeoJsonAllOf, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ZoneCollectionGeoJson) SetFeatures(v []ZoneGeoJsonAllOf)`

SetFeatures sets Features field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


