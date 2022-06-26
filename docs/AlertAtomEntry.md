# AlertAtomEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 
**Published** | Pointer to **string** |  | [optional] 
**Author** | Pointer to [**AlertAtomEntryAuthor**](AlertAtomEntryAuthor.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Sent** | Pointer to **string** |  | [optional] 
**Effective** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MsgType** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Urgency** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Certainty** | Pointer to **string** |  | [optional] 
**AreaDesc** | Pointer to **string** |  | [optional] 
**Polygon** | Pointer to **string** |  | [optional] 
**Geocode** | Pointer to [**[]AlertXMLParameter**](AlertXMLParameter.md) |  | [optional] 
**Parameter** | Pointer to [**[]AlertXMLParameter**](AlertXMLParameter.md) |  | [optional] 

## Methods

### NewAlertAtomEntry

`func NewAlertAtomEntry() *AlertAtomEntry`

NewAlertAtomEntry instantiates a new AlertAtomEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertAtomEntryWithDefaults

`func NewAlertAtomEntryWithDefaults() *AlertAtomEntry`

NewAlertAtomEntryWithDefaults instantiates a new AlertAtomEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertAtomEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertAtomEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertAtomEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AlertAtomEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdated

`func (o *AlertAtomEntry) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *AlertAtomEntry) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *AlertAtomEntry) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *AlertAtomEntry) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetPublished

`func (o *AlertAtomEntry) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *AlertAtomEntry) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *AlertAtomEntry) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *AlertAtomEntry) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetAuthor

`func (o *AlertAtomEntry) GetAuthor() AlertAtomEntryAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AlertAtomEntry) GetAuthorOk() (*AlertAtomEntryAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AlertAtomEntry) SetAuthor(v AlertAtomEntryAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AlertAtomEntry) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetSummary

`func (o *AlertAtomEntry) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AlertAtomEntry) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AlertAtomEntry) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *AlertAtomEntry) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetEvent

`func (o *AlertAtomEntry) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AlertAtomEntry) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AlertAtomEntry) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *AlertAtomEntry) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetSent

`func (o *AlertAtomEntry) GetSent() string`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *AlertAtomEntry) GetSentOk() (*string, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *AlertAtomEntry) SetSent(v string)`

SetSent sets Sent field to given value.

### HasSent

`func (o *AlertAtomEntry) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetEffective

`func (o *AlertAtomEntry) GetEffective() string`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *AlertAtomEntry) GetEffectiveOk() (*string, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *AlertAtomEntry) SetEffective(v string)`

SetEffective sets Effective field to given value.

### HasEffective

`func (o *AlertAtomEntry) HasEffective() bool`

HasEffective returns a boolean if a field has been set.

### GetExpires

`func (o *AlertAtomEntry) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AlertAtomEntry) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AlertAtomEntry) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AlertAtomEntry) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetStatus

`func (o *AlertAtomEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertAtomEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertAtomEntry) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertAtomEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsgType

`func (o *AlertAtomEntry) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *AlertAtomEntry) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *AlertAtomEntry) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *AlertAtomEntry) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetCategory

`func (o *AlertAtomEntry) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AlertAtomEntry) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AlertAtomEntry) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AlertAtomEntry) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetUrgency

`func (o *AlertAtomEntry) GetUrgency() string`

GetUrgency returns the Urgency field if non-nil, zero value otherwise.

### GetUrgencyOk

`func (o *AlertAtomEntry) GetUrgencyOk() (*string, bool)`

GetUrgencyOk returns a tuple with the Urgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgency

`func (o *AlertAtomEntry) SetUrgency(v string)`

SetUrgency sets Urgency field to given value.

### HasUrgency

`func (o *AlertAtomEntry) HasUrgency() bool`

HasUrgency returns a boolean if a field has been set.

### GetSeverity

`func (o *AlertAtomEntry) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AlertAtomEntry) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AlertAtomEntry) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AlertAtomEntry) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCertainty

`func (o *AlertAtomEntry) GetCertainty() string`

GetCertainty returns the Certainty field if non-nil, zero value otherwise.

### GetCertaintyOk

`func (o *AlertAtomEntry) GetCertaintyOk() (*string, bool)`

GetCertaintyOk returns a tuple with the Certainty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertainty

`func (o *AlertAtomEntry) SetCertainty(v string)`

SetCertainty sets Certainty field to given value.

### HasCertainty

`func (o *AlertAtomEntry) HasCertainty() bool`

HasCertainty returns a boolean if a field has been set.

### GetAreaDesc

`func (o *AlertAtomEntry) GetAreaDesc() string`

GetAreaDesc returns the AreaDesc field if non-nil, zero value otherwise.

### GetAreaDescOk

`func (o *AlertAtomEntry) GetAreaDescOk() (*string, bool)`

GetAreaDescOk returns a tuple with the AreaDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaDesc

`func (o *AlertAtomEntry) SetAreaDesc(v string)`

SetAreaDesc sets AreaDesc field to given value.

### HasAreaDesc

`func (o *AlertAtomEntry) HasAreaDesc() bool`

HasAreaDesc returns a boolean if a field has been set.

### GetPolygon

`func (o *AlertAtomEntry) GetPolygon() string`

GetPolygon returns the Polygon field if non-nil, zero value otherwise.

### GetPolygonOk

`func (o *AlertAtomEntry) GetPolygonOk() (*string, bool)`

GetPolygonOk returns a tuple with the Polygon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolygon

`func (o *AlertAtomEntry) SetPolygon(v string)`

SetPolygon sets Polygon field to given value.

### HasPolygon

`func (o *AlertAtomEntry) HasPolygon() bool`

HasPolygon returns a boolean if a field has been set.

### GetGeocode

`func (o *AlertAtomEntry) GetGeocode() []AlertXMLParameter`

GetGeocode returns the Geocode field if non-nil, zero value otherwise.

### GetGeocodeOk

`func (o *AlertAtomEntry) GetGeocodeOk() (*[]AlertXMLParameter, bool)`

GetGeocodeOk returns a tuple with the Geocode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocode

`func (o *AlertAtomEntry) SetGeocode(v []AlertXMLParameter)`

SetGeocode sets Geocode field to given value.

### HasGeocode

`func (o *AlertAtomEntry) HasGeocode() bool`

HasGeocode returns a boolean if a field has been set.

### GetParameter

`func (o *AlertAtomEntry) GetParameter() []AlertXMLParameter`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *AlertAtomEntry) GetParameterOk() (*[]AlertXMLParameter, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *AlertAtomEntry) SetParameter(v []AlertXMLParameter)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *AlertAtomEntry) HasParameter() bool`

HasParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


