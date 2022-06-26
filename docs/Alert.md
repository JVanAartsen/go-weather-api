# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the alert message. | [optional] 
**AreaDesc** | Pointer to **string** | A textual description of the area affected by the alert. | [optional] 
**Geocode** | Pointer to [**AlertGeocode**](AlertGeocode.md) |  | [optional] 
**AffectedZones** | Pointer to **[]string** | An array of API links for zones affected by the alert. This is an API-specific extension field and is not part of the CAP specification.  | [optional] 
**References** | Pointer to [**[]AlertReferencesInner**](AlertReferencesInner.md) | A list of prior alerts that this alert updates or replaces. | [optional] 
**Sent** | Pointer to **time.Time** | The time of the origination of the alert message. | [optional] 
**Effective** | Pointer to **time.Time** | The effective time of the information of the alert message. | [optional] 
**Onset** | Pointer to **NullableTime** | The expected time of the beginning of the subject event of the alert message. | [optional] 
**Expires** | Pointer to **time.Time** | The expiry time of the information of the alert message. | [optional] 
**Ends** | Pointer to **NullableTime** | The expected end time of the subject event of the alert message. | [optional] 
**Status** | Pointer to [**AlertStatus**](AlertStatus.md) |  | [optional] 
**MessageType** | Pointer to [**AlertMessageType**](AlertMessageType.md) |  | [optional] 
**Category** | Pointer to **string** | The code denoting the category of the subject event of the alert message. | [optional] 
**Severity** | Pointer to [**AlertSeverity**](AlertSeverity.md) |  | [optional] 
**Certainty** | Pointer to [**AlertCertainty**](AlertCertainty.md) |  | [optional] 
**Urgency** | Pointer to [**AlertUrgency**](AlertUrgency.md) |  | [optional] 
**Event** | Pointer to **string** | The text denoting the type of the subject event of the alert message. | [optional] 
**Sender** | Pointer to **string** | Email address of the NWS webmaster. | [optional] 
**SenderName** | Pointer to **string** | The text naming the originator of the alert message. | [optional] 
**Headline** | Pointer to **NullableString** | The text headline of the alert message. | [optional] 
**Description** | Pointer to **string** | The text describing the subject event of the alert message. | [optional] 
**Instruction** | Pointer to **NullableString** | The text describing the recommended action to be taken by recipients of the alert message.  | [optional] 
**Response** | Pointer to **string** | The code denoting the type of action recommended for the target audience. This corresponds to responseType in the CAP specification.  | [optional] 
**Parameters** | Pointer to [**map[string]Array**](array.md) | System-specific additional parameters associated with the alert message. The keys in this object correspond to parameter definitions in the NWS CAP specification.  | [optional] 

## Methods

### NewAlert

`func NewAlert() *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Alert) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alert) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alert) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Alert) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAreaDesc

`func (o *Alert) GetAreaDesc() string`

GetAreaDesc returns the AreaDesc field if non-nil, zero value otherwise.

### GetAreaDescOk

`func (o *Alert) GetAreaDescOk() (*string, bool)`

GetAreaDescOk returns a tuple with the AreaDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaDesc

`func (o *Alert) SetAreaDesc(v string)`

SetAreaDesc sets AreaDesc field to given value.

### HasAreaDesc

`func (o *Alert) HasAreaDesc() bool`

HasAreaDesc returns a boolean if a field has been set.

### GetGeocode

`func (o *Alert) GetGeocode() AlertGeocode`

GetGeocode returns the Geocode field if non-nil, zero value otherwise.

### GetGeocodeOk

`func (o *Alert) GetGeocodeOk() (*AlertGeocode, bool)`

GetGeocodeOk returns a tuple with the Geocode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeocode

`func (o *Alert) SetGeocode(v AlertGeocode)`

SetGeocode sets Geocode field to given value.

### HasGeocode

`func (o *Alert) HasGeocode() bool`

HasGeocode returns a boolean if a field has been set.

### GetAffectedZones

`func (o *Alert) GetAffectedZones() []string`

GetAffectedZones returns the AffectedZones field if non-nil, zero value otherwise.

### GetAffectedZonesOk

`func (o *Alert) GetAffectedZonesOk() (*[]string, bool)`

GetAffectedZonesOk returns a tuple with the AffectedZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedZones

`func (o *Alert) SetAffectedZones(v []string)`

SetAffectedZones sets AffectedZones field to given value.

### HasAffectedZones

`func (o *Alert) HasAffectedZones() bool`

HasAffectedZones returns a boolean if a field has been set.

### GetReferences

`func (o *Alert) GetReferences() []AlertReferencesInner`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *Alert) GetReferencesOk() (*[]AlertReferencesInner, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *Alert) SetReferences(v []AlertReferencesInner)`

SetReferences sets References field to given value.

### HasReferences

`func (o *Alert) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetSent

`func (o *Alert) GetSent() time.Time`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *Alert) GetSentOk() (*time.Time, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *Alert) SetSent(v time.Time)`

SetSent sets Sent field to given value.

### HasSent

`func (o *Alert) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetEffective

`func (o *Alert) GetEffective() time.Time`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *Alert) GetEffectiveOk() (*time.Time, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *Alert) SetEffective(v time.Time)`

SetEffective sets Effective field to given value.

### HasEffective

`func (o *Alert) HasEffective() bool`

HasEffective returns a boolean if a field has been set.

### GetOnset

`func (o *Alert) GetOnset() time.Time`

GetOnset returns the Onset field if non-nil, zero value otherwise.

### GetOnsetOk

`func (o *Alert) GetOnsetOk() (*time.Time, bool)`

GetOnsetOk returns a tuple with the Onset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnset

`func (o *Alert) SetOnset(v time.Time)`

SetOnset sets Onset field to given value.

### HasOnset

`func (o *Alert) HasOnset() bool`

HasOnset returns a boolean if a field has been set.

### SetOnsetNil

`func (o *Alert) SetOnsetNil(b bool)`

 SetOnsetNil sets the value for Onset to be an explicit nil

### UnsetOnset
`func (o *Alert) UnsetOnset()`

UnsetOnset ensures that no value is present for Onset, not even an explicit nil
### GetExpires

`func (o *Alert) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *Alert) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *Alert) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *Alert) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetEnds

`func (o *Alert) GetEnds() time.Time`

GetEnds returns the Ends field if non-nil, zero value otherwise.

### GetEndsOk

`func (o *Alert) GetEndsOk() (*time.Time, bool)`

GetEndsOk returns a tuple with the Ends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnds

`func (o *Alert) SetEnds(v time.Time)`

SetEnds sets Ends field to given value.

### HasEnds

`func (o *Alert) HasEnds() bool`

HasEnds returns a boolean if a field has been set.

### SetEndsNil

`func (o *Alert) SetEndsNil(b bool)`

 SetEndsNil sets the value for Ends to be an explicit nil

### UnsetEnds
`func (o *Alert) UnsetEnds()`

UnsetEnds ensures that no value is present for Ends, not even an explicit nil
### GetStatus

`func (o *Alert) GetStatus() AlertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Alert) GetStatusOk() (*AlertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Alert) SetStatus(v AlertStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Alert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessageType

`func (o *Alert) GetMessageType() AlertMessageType`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *Alert) GetMessageTypeOk() (*AlertMessageType, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *Alert) SetMessageType(v AlertMessageType)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *Alert) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetCategory

`func (o *Alert) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Alert) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Alert) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Alert) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSeverity

`func (o *Alert) GetSeverity() AlertSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Alert) GetSeverityOk() (*AlertSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Alert) SetSeverity(v AlertSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Alert) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCertainty

`func (o *Alert) GetCertainty() AlertCertainty`

GetCertainty returns the Certainty field if non-nil, zero value otherwise.

### GetCertaintyOk

`func (o *Alert) GetCertaintyOk() (*AlertCertainty, bool)`

GetCertaintyOk returns a tuple with the Certainty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertainty

`func (o *Alert) SetCertainty(v AlertCertainty)`

SetCertainty sets Certainty field to given value.

### HasCertainty

`func (o *Alert) HasCertainty() bool`

HasCertainty returns a boolean if a field has been set.

### GetUrgency

`func (o *Alert) GetUrgency() AlertUrgency`

GetUrgency returns the Urgency field if non-nil, zero value otherwise.

### GetUrgencyOk

`func (o *Alert) GetUrgencyOk() (*AlertUrgency, bool)`

GetUrgencyOk returns a tuple with the Urgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrgency

`func (o *Alert) SetUrgency(v AlertUrgency)`

SetUrgency sets Urgency field to given value.

### HasUrgency

`func (o *Alert) HasUrgency() bool`

HasUrgency returns a boolean if a field has been set.

### GetEvent

`func (o *Alert) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *Alert) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *Alert) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *Alert) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetSender

`func (o *Alert) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Alert) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Alert) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *Alert) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetSenderName

`func (o *Alert) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *Alert) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *Alert) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *Alert) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### GetHeadline

`func (o *Alert) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *Alert) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *Alert) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *Alert) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### SetHeadlineNil

`func (o *Alert) SetHeadlineNil(b bool)`

 SetHeadlineNil sets the value for Headline to be an explicit nil

### UnsetHeadline
`func (o *Alert) UnsetHeadline()`

UnsetHeadline ensures that no value is present for Headline, not even an explicit nil
### GetDescription

`func (o *Alert) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Alert) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Alert) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Alert) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstruction

`func (o *Alert) GetInstruction() string`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *Alert) GetInstructionOk() (*string, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *Alert) SetInstruction(v string)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *Alert) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### SetInstructionNil

`func (o *Alert) SetInstructionNil(b bool)`

 SetInstructionNil sets the value for Instruction to be an explicit nil

### UnsetInstruction
`func (o *Alert) UnsetInstruction()`

UnsetInstruction ensures that no value is present for Instruction, not even an explicit nil
### GetResponse

`func (o *Alert) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Alert) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Alert) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Alert) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetParameters

`func (o *Alert) GetParameters() map[string]Array`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Alert) GetParametersOk() (*map[string]Array, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Alert) SetParameters(v map[string]Array)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Alert) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


