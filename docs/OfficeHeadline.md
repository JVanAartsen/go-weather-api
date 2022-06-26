# OfficeHeadline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Office** | Pointer to **string** |  | [optional] 
**Important** | Pointer to **bool** |  | [optional] 
**IssuanceTime** | Pointer to **time.Time** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **NullableString** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 

## Methods

### NewOfficeHeadline

`func NewOfficeHeadline() *OfficeHeadline`

NewOfficeHeadline instantiates a new OfficeHeadline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfficeHeadlineWithDefaults

`func NewOfficeHeadlineWithDefaults() *OfficeHeadline`

NewOfficeHeadlineWithDefaults instantiates a new OfficeHeadline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *OfficeHeadline) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OfficeHeadline) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OfficeHeadline) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *OfficeHeadline) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetId

`func (o *OfficeHeadline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OfficeHeadline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OfficeHeadline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OfficeHeadline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetId

`func (o *OfficeHeadline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OfficeHeadline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OfficeHeadline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OfficeHeadline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOffice

`func (o *OfficeHeadline) GetOffice() string`

GetOffice returns the Office field if non-nil, zero value otherwise.

### GetOfficeOk

`func (o *OfficeHeadline) GetOfficeOk() (*string, bool)`

GetOfficeOk returns a tuple with the Office field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffice

`func (o *OfficeHeadline) SetOffice(v string)`

SetOffice sets Office field to given value.

### HasOffice

`func (o *OfficeHeadline) HasOffice() bool`

HasOffice returns a boolean if a field has been set.

### GetImportant

`func (o *OfficeHeadline) GetImportant() bool`

GetImportant returns the Important field if non-nil, zero value otherwise.

### GetImportantOk

`func (o *OfficeHeadline) GetImportantOk() (*bool, bool)`

GetImportantOk returns a tuple with the Important field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportant

`func (o *OfficeHeadline) SetImportant(v bool)`

SetImportant sets Important field to given value.

### HasImportant

`func (o *OfficeHeadline) HasImportant() bool`

HasImportant returns a boolean if a field has been set.

### GetIssuanceTime

`func (o *OfficeHeadline) GetIssuanceTime() time.Time`

GetIssuanceTime returns the IssuanceTime field if non-nil, zero value otherwise.

### GetIssuanceTimeOk

`func (o *OfficeHeadline) GetIssuanceTimeOk() (*time.Time, bool)`

GetIssuanceTimeOk returns a tuple with the IssuanceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceTime

`func (o *OfficeHeadline) SetIssuanceTime(v time.Time)`

SetIssuanceTime sets IssuanceTime field to given value.

### HasIssuanceTime

`func (o *OfficeHeadline) HasIssuanceTime() bool`

HasIssuanceTime returns a boolean if a field has been set.

### GetLink

`func (o *OfficeHeadline) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *OfficeHeadline) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *OfficeHeadline) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *OfficeHeadline) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetName

`func (o *OfficeHeadline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OfficeHeadline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OfficeHeadline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OfficeHeadline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *OfficeHeadline) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OfficeHeadline) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OfficeHeadline) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OfficeHeadline) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSummary

`func (o *OfficeHeadline) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *OfficeHeadline) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *OfficeHeadline) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *OfficeHeadline) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *OfficeHeadline) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *OfficeHeadline) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetContent

`func (o *OfficeHeadline) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OfficeHeadline) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OfficeHeadline) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OfficeHeadline) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


