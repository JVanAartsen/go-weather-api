# TextProductLocationCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Locations** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTextProductLocationCollection

`func NewTextProductLocationCollection() *TextProductLocationCollection`

NewTextProductLocationCollection instantiates a new TextProductLocationCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTextProductLocationCollectionWithDefaults

`func NewTextProductLocationCollectionWithDefaults() *TextProductLocationCollection`

NewTextProductLocationCollectionWithDefaults instantiates a new TextProductLocationCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *TextProductLocationCollection) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TextProductLocationCollection) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TextProductLocationCollection) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *TextProductLocationCollection) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLocations

`func (o *TextProductLocationCollection) GetLocations() map[string]string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *TextProductLocationCollection) GetLocationsOk() (*map[string]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *TextProductLocationCollection) SetLocations(v map[string]string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *TextProductLocationCollection) HasLocations() bool`

HasLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


