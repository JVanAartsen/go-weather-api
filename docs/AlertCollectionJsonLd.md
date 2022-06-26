# AlertCollectionJsonLd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A title describing the alert collection | [optional] 
**Updated** | Pointer to **time.Time** | The last time a change occurred to this collection | [optional] 
**Pagination** | Pointer to [**AlertCollectionPagination**](AlertCollectionPagination.md) |  | [optional] 
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Graph** | Pointer to [**[]Alert**](Alert.md) |  | [optional] 

## Methods

### NewAlertCollectionJsonLd

`func NewAlertCollectionJsonLd() *AlertCollectionJsonLd`

NewAlertCollectionJsonLd instantiates a new AlertCollectionJsonLd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertCollectionJsonLdWithDefaults

`func NewAlertCollectionJsonLdWithDefaults() *AlertCollectionJsonLd`

NewAlertCollectionJsonLdWithDefaults instantiates a new AlertCollectionJsonLd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AlertCollectionJsonLd) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertCollectionJsonLd) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlertCollectionJsonLd) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlertCollectionJsonLd) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *AlertCollectionJsonLd) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AlertCollectionJsonLd) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AlertCollectionJsonLd) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AlertCollectionJsonLd) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetPagination

`func (o *AlertCollectionJsonLd) GetPagination() AlertCollectionPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AlertCollectionJsonLd) GetPaginationOk() (*AlertCollectionPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AlertCollectionJsonLd) SetPagination(v AlertCollectionPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AlertCollectionJsonLd) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetContext

`func (o *AlertCollectionJsonLd) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *AlertCollectionJsonLd) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *AlertCollectionJsonLd) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *AlertCollectionJsonLd) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetGraph

`func (o *AlertCollectionJsonLd) GetGraph() []Alert`

GetGraph returns the Graph field if non-nil, zero value otherwise.

### GetGraphOk

`func (o *AlertCollectionJsonLd) GetGraphOk() (*[]Alert, bool)`

GetGraphOk returns a tuple with the Graph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraph

`func (o *AlertCollectionJsonLd) SetGraph(v []Alert)`

SetGraph sets Graph field to given value.

### HasGraph

`func (o *AlertCollectionJsonLd) HasGraph() bool`

HasGraph returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


