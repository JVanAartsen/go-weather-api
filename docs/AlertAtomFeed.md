# AlertAtomFeed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Generator** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**AlertAtomFeedAuthor**](AlertAtomFeedAuthor.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Entry** | Pointer to [**[]AlertAtomEntry**](AlertAtomEntry.md) |  | [optional] 

## Methods

### NewAlertAtomFeed

`func NewAlertAtomFeed() *AlertAtomFeed`

NewAlertAtomFeed instantiates a new AlertAtomFeed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertAtomFeedWithDefaults

`func NewAlertAtomFeedWithDefaults() *AlertAtomFeed`

NewAlertAtomFeedWithDefaults instantiates a new AlertAtomFeed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertAtomFeed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertAtomFeed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertAtomFeed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertAtomFeed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGenerator

`func (o *AlertAtomFeed) GetGenerator() string`

GetGenerator returns the Generator field if non-nil, zero value otherwise.

### GetGeneratorOk

`func (o *AlertAtomFeed) GetGeneratorOk() (*string, bool)`

GetGeneratorOk returns a tuple with the Generator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerator

`func (o *AlertAtomFeed) SetGenerator(v string)`

SetGenerator sets Generator field to given value.

### HasGenerator

`func (o *AlertAtomFeed) HasGenerator() bool`

HasGenerator returns a boolean if a field has been set.

### GetUpdated

`func (o *AlertAtomFeed) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AlertAtomFeed) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AlertAtomFeed) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AlertAtomFeed) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetAuthor

`func (o *AlertAtomFeed) GetAuthor() AlertAtomFeedAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AlertAtomFeed) GetAuthorOk() (*AlertAtomFeedAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AlertAtomFeed) SetAuthor(v AlertAtomFeedAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AlertAtomFeed) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetTitle

`func (o *AlertAtomFeed) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AlertAtomFeed) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AlertAtomFeed) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *AlertAtomFeed) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEntry

`func (o *AlertAtomFeed) GetEntry() []AlertAtomEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *AlertAtomFeed) GetEntryOk() (*[]AlertAtomEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *AlertAtomFeed) SetEntry(v []AlertAtomEntry)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *AlertAtomFeed) HasEntry() bool`

HasEntry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


