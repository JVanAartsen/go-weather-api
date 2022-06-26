# IconsSummary200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Icons** | [**map[string]IconsSummary200ResponseIconsValue**](IconsSummary200ResponseIconsValue.md) |  | 

## Methods

### NewIconsSummary200Response

`func NewIconsSummary200Response(icons map[string]IconsSummary200ResponseIconsValue, ) *IconsSummary200Response`

NewIconsSummary200Response instantiates a new IconsSummary200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIconsSummary200ResponseWithDefaults

`func NewIconsSummary200ResponseWithDefaults() *IconsSummary200Response`

NewIconsSummary200ResponseWithDefaults instantiates a new IconsSummary200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *IconsSummary200Response) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *IconsSummary200Response) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *IconsSummary200Response) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *IconsSummary200Response) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetIcons

`func (o *IconsSummary200Response) GetIcons() map[string]IconsSummary200ResponseIconsValue`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *IconsSummary200Response) GetIconsOk() (*map[string]IconsSummary200ResponseIconsValue, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *IconsSummary200Response) SetIcons(v map[string]IconsSummary200ResponseIconsValue)`

SetIcons sets Icons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


