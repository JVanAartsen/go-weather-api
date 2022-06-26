# AlertCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A title describing the alert collection | [optional] 
**Updated** | Pointer to **time.Time** | The last time a change occurred to this collection | [optional] 
**Pagination** | Pointer to [**AlertCollectionPagination**](AlertCollectionPagination.md) |  | [optional] 

## Methods

### NewAlertCollection

`func NewAlertCollection() *AlertCollection`

NewAlertCollection instantiates a new AlertCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertCollectionWithDefaults

`func NewAlertCollectionWithDefaults() *AlertCollection`

NewAlertCollectionWithDefaults instantiates a new AlertCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *AlertCollection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertCollection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlertCollection) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlertCollection) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdated

`func (o *AlertCollection) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AlertCollection) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AlertCollection) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AlertCollection) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetPagination

`func (o *AlertCollection) GetPagination() AlertCollectionPagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AlertCollection) GetPaginationOk() (*AlertCollectionPagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AlertCollection) SetPagination(v AlertCollectionPagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AlertCollection) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


