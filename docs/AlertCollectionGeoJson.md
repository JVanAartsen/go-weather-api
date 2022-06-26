# AlertCollectionGeoJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Type** | **string** |  | 
**Features** | [**[]AlertCollectionGeoJsonAllOfFeatures**](AlertCollectionGeoJsonAllOfFeatures.md) |  | 
**Title** | Pointer to **string** | A title describing the alert collection | [optional] 
**Updated** | Pointer to **time.Time** | The last time a change occurred to this collection | [optional] 
**Pagination** | Pointer to [**AlertCollectionPagination**](AlertCollectionPagination.md) |  | [optional] 

## Methods

### NewAlertCollectionGeoJson

`func NewAlertCollectionGeoJson(type_ string, features []AlertCollectionGeoJsonAllOfFeatures, ) *AlertCollectionGeoJson`

NewAlertCollectionGeoJson instantiates a new AlertCollectionGeoJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertCollectionGeoJsonWithDefaults

`func NewAlertCollectionGeoJsonWithDefaults() *AlertCollectionGeoJson`

NewAlertCollectionGeoJsonWithDefaults instantiates a new AlertCollectionGeoJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *AlertCollectionGeoJson) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AlertCollectionGeoJson) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AlertCollectionGeoJson) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *AlertCollectionGeoJson) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetType

`func (o *AlertCollectionGeoJson) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertCollectionGeoJson) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertCollectionGeoJson) SetType(v string)`

SetType sets Type field to given value.


### GetFeatures

`func (o *AlertCollectionGeoJson) GetFeatures() []AlertCollectionGeoJsonAllOfFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AlertCollectionGeoJson) GetFeaturesOk() (*[]AlertCollectionGeoJsonAllOfFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AlertCollectionGeoJson) SetFeatures(v []AlertCollectionGeoJsonAllOfFeatures)`

SetFeatures sets Features field to given value.


### GetTitle

`func (o *AlertCollectionGeoJson) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertCollectionGeoJson) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlertCollectionGeoJson) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlertCollectionGeoJson) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *AlertCollectionGeoJson) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AlertCollectionGeoJson) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AlertCollectionGeoJson) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AlertCollectionGeoJson) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetPagination

`func (o *AlertCollectionGeoJson) GetPagination() AlertCollectionPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AlertCollectionGeoJson) GetPaginationOk() (*AlertCollectionPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AlertCollectionGeoJson) SetPagination(v AlertCollectionPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AlertCollectionGeoJson) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


