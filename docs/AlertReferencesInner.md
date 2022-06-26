# AlertReferencesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | An API link to the prior alert. | [optional] 
**Identifier** | Pointer to **string** | The identifier of the alert message. | [optional] 
**Sender** | Pointer to **string** | The sender of the prior alert. | [optional] 
**Sent** | Pointer to **time.Time** | The time the prior alert was sent. | [optional] 

## Methods

### NewAlertReferencesInner

`func NewAlertReferencesInner() *AlertReferencesInner`

NewAlertReferencesInner instantiates a new AlertReferencesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertReferencesInnerWithDefaults

`func NewAlertReferencesInnerWithDefaults() *AlertReferencesInner`

NewAlertReferencesInnerWithDefaults instantiates a new AlertReferencesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertReferencesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertReferencesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertReferencesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertReferencesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *AlertReferencesInner) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AlertReferencesInner) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AlertReferencesInner) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *AlertReferencesInner) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetSender

`func (o *AlertReferencesInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *AlertReferencesInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *AlertReferencesInner) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *AlertReferencesInner) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetSent

`func (o *AlertReferencesInner) GetSent() time.Time`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *AlertReferencesInner) GetSentOk() (*time.Time, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *AlertReferencesInner) SetSent(v time.Time)`

SetSent sets Sent field to given value.

### HasSent

`func (o *AlertReferencesInner) HasSent() bool`

HasSent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


