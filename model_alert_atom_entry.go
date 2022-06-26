/*
weather.gov API

weather.gov API

API version: 1.8.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package weatherApi

import (
	"encoding/json"
)

// AlertAtomEntry An alert entry in an Atom feed
type AlertAtomEntry struct {
	Id *string `json:"id,omitempty"`
	Updated *string `json:"updated,omitempty"`
	Published *string `json:"published,omitempty"`
	Author *AlertAtomEntryAuthor `json:"author,omitempty"`
	Summary *string `json:"summary,omitempty"`
	Event *string `json:"event,omitempty"`
	Sent *string `json:"sent,omitempty"`
	Effective *string `json:"effective,omitempty"`
	Expires *string `json:"expires,omitempty"`
	Status *string `json:"status,omitempty"`
	MsgType *string `json:"msgType,omitempty"`
	Category *string `json:"category,omitempty"`
	Urgency *string `json:"urgency,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Certainty *string `json:"certainty,omitempty"`
	AreaDesc *string `json:"areaDesc,omitempty"`
	Polygon *string `json:"polygon,omitempty"`
	Geocode []AlertXMLParameter `json:"geocode,omitempty"`
	Parameter []AlertXMLParameter `json:"parameter,omitempty"`
}

// NewAlertAtomEntry instantiates a new AlertAtomEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertAtomEntry() *AlertAtomEntry {
	this := AlertAtomEntry{}
	return &this
}

// NewAlertAtomEntryWithDefaults instantiates a new AlertAtomEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertAtomEntryWithDefaults() *AlertAtomEntry {
	this := AlertAtomEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AlertAtomEntry) SetId(v string) {
	o.Id = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetUpdated() string {
	if o == nil || o.Updated == nil {
		var ret string
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetUpdatedOk() (*string, bool) {
	if o == nil || o.Updated == nil {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasUpdated() bool {
	if o != nil && o.Updated != nil {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given string and assigns it to the Updated field.
func (o *AlertAtomEntry) SetUpdated(v string) {
	o.Updated = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetPublished() string {
	if o == nil || o.Published == nil {
		var ret string
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetPublishedOk() (*string, bool) {
	if o == nil || o.Published == nil {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasPublished() bool {
	if o != nil && o.Published != nil {
		return true
	}

	return false
}

// SetPublished gets a reference to the given string and assigns it to the Published field.
func (o *AlertAtomEntry) SetPublished(v string) {
	o.Published = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetAuthor() AlertAtomEntryAuthor {
	if o == nil || o.Author == nil {
		var ret AlertAtomEntryAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetAuthorOk() (*AlertAtomEntryAuthor, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasAuthor() bool {
	if o != nil && o.Author != nil {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given AlertAtomEntryAuthor and assigns it to the Author field.
func (o *AlertAtomEntry) SetAuthor(v AlertAtomEntryAuthor) {
	o.Author = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *AlertAtomEntry) SetSummary(v string) {
	o.Summary = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *AlertAtomEntry) SetEvent(v string) {
	o.Event = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetSent() string {
	if o == nil || o.Sent == nil {
		var ret string
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetSentOk() (*string, bool) {
	if o == nil || o.Sent == nil {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasSent() bool {
	if o != nil && o.Sent != nil {
		return true
	}

	return false
}

// SetSent gets a reference to the given string and assigns it to the Sent field.
func (o *AlertAtomEntry) SetSent(v string) {
	o.Sent = &v
}

// GetEffective returns the Effective field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetEffective() string {
	if o == nil || o.Effective == nil {
		var ret string
		return ret
	}
	return *o.Effective
}

// GetEffectiveOk returns a tuple with the Effective field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetEffectiveOk() (*string, bool) {
	if o == nil || o.Effective == nil {
		return nil, false
	}
	return o.Effective, true
}

// HasEffective returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasEffective() bool {
	if o != nil && o.Effective != nil {
		return true
	}

	return false
}

// SetEffective gets a reference to the given string and assigns it to the Effective field.
func (o *AlertAtomEntry) SetEffective(v string) {
	o.Effective = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetExpires() string {
	if o == nil || o.Expires == nil {
		var ret string
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetExpiresOk() (*string, bool) {
	if o == nil || o.Expires == nil {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasExpires() bool {
	if o != nil && o.Expires != nil {
		return true
	}

	return false
}

// SetExpires gets a reference to the given string and assigns it to the Expires field.
func (o *AlertAtomEntry) SetExpires(v string) {
	o.Expires = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlertAtomEntry) SetStatus(v string) {
	o.Status = &v
}

// GetMsgType returns the MsgType field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetMsgType() string {
	if o == nil || o.MsgType == nil {
		var ret string
		return ret
	}
	return *o.MsgType
}

// GetMsgTypeOk returns a tuple with the MsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetMsgTypeOk() (*string, bool) {
	if o == nil || o.MsgType == nil {
		return nil, false
	}
	return o.MsgType, true
}

// HasMsgType returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasMsgType() bool {
	if o != nil && o.MsgType != nil {
		return true
	}

	return false
}

// SetMsgType gets a reference to the given string and assigns it to the MsgType field.
func (o *AlertAtomEntry) SetMsgType(v string) {
	o.MsgType = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *AlertAtomEntry) SetCategory(v string) {
	o.Category = &v
}

// GetUrgency returns the Urgency field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetUrgency() string {
	if o == nil || o.Urgency == nil {
		var ret string
		return ret
	}
	return *o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetUrgencyOk() (*string, bool) {
	if o == nil || o.Urgency == nil {
		return nil, false
	}
	return o.Urgency, true
}

// HasUrgency returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasUrgency() bool {
	if o != nil && o.Urgency != nil {
		return true
	}

	return false
}

// SetUrgency gets a reference to the given string and assigns it to the Urgency field.
func (o *AlertAtomEntry) SetUrgency(v string) {
	o.Urgency = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *AlertAtomEntry) SetSeverity(v string) {
	o.Severity = &v
}

// GetCertainty returns the Certainty field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetCertainty() string {
	if o == nil || o.Certainty == nil {
		var ret string
		return ret
	}
	return *o.Certainty
}

// GetCertaintyOk returns a tuple with the Certainty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetCertaintyOk() (*string, bool) {
	if o == nil || o.Certainty == nil {
		return nil, false
	}
	return o.Certainty, true
}

// HasCertainty returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasCertainty() bool {
	if o != nil && o.Certainty != nil {
		return true
	}

	return false
}

// SetCertainty gets a reference to the given string and assigns it to the Certainty field.
func (o *AlertAtomEntry) SetCertainty(v string) {
	o.Certainty = &v
}

// GetAreaDesc returns the AreaDesc field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetAreaDesc() string {
	if o == nil || o.AreaDesc == nil {
		var ret string
		return ret
	}
	return *o.AreaDesc
}

// GetAreaDescOk returns a tuple with the AreaDesc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetAreaDescOk() (*string, bool) {
	if o == nil || o.AreaDesc == nil {
		return nil, false
	}
	return o.AreaDesc, true
}

// HasAreaDesc returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasAreaDesc() bool {
	if o != nil && o.AreaDesc != nil {
		return true
	}

	return false
}

// SetAreaDesc gets a reference to the given string and assigns it to the AreaDesc field.
func (o *AlertAtomEntry) SetAreaDesc(v string) {
	o.AreaDesc = &v
}

// GetPolygon returns the Polygon field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetPolygon() string {
	if o == nil || o.Polygon == nil {
		var ret string
		return ret
	}
	return *o.Polygon
}

// GetPolygonOk returns a tuple with the Polygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetPolygonOk() (*string, bool) {
	if o == nil || o.Polygon == nil {
		return nil, false
	}
	return o.Polygon, true
}

// HasPolygon returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasPolygon() bool {
	if o != nil && o.Polygon != nil {
		return true
	}

	return false
}

// SetPolygon gets a reference to the given string and assigns it to the Polygon field.
func (o *AlertAtomEntry) SetPolygon(v string) {
	o.Polygon = &v
}

// GetGeocode returns the Geocode field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetGeocode() []AlertXMLParameter {
	if o == nil || o.Geocode == nil {
		var ret []AlertXMLParameter
		return ret
	}
	return o.Geocode
}

// GetGeocodeOk returns a tuple with the Geocode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetGeocodeOk() ([]AlertXMLParameter, bool) {
	if o == nil || o.Geocode == nil {
		return nil, false
	}
	return o.Geocode, true
}

// HasGeocode returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasGeocode() bool {
	if o != nil && o.Geocode != nil {
		return true
	}

	return false
}

// SetGeocode gets a reference to the given []AlertXMLParameter and assigns it to the Geocode field.
func (o *AlertAtomEntry) SetGeocode(v []AlertXMLParameter) {
	o.Geocode = v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *AlertAtomEntry) GetParameter() []AlertXMLParameter {
	if o == nil || o.Parameter == nil {
		var ret []AlertXMLParameter
		return ret
	}
	return o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertAtomEntry) GetParameterOk() ([]AlertXMLParameter, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *AlertAtomEntry) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given []AlertXMLParameter and assigns it to the Parameter field.
func (o *AlertAtomEntry) SetParameter(v []AlertXMLParameter) {
	o.Parameter = v
}

func (o AlertAtomEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Updated != nil {
		toSerialize["updated"] = o.Updated
	}
	if o.Published != nil {
		toSerialize["published"] = o.Published
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Sent != nil {
		toSerialize["sent"] = o.Sent
	}
	if o.Effective != nil {
		toSerialize["effective"] = o.Effective
	}
	if o.Expires != nil {
		toSerialize["expires"] = o.Expires
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.MsgType != nil {
		toSerialize["msgType"] = o.MsgType
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Urgency != nil {
		toSerialize["urgency"] = o.Urgency
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Certainty != nil {
		toSerialize["certainty"] = o.Certainty
	}
	if o.AreaDesc != nil {
		toSerialize["areaDesc"] = o.AreaDesc
	}
	if o.Polygon != nil {
		toSerialize["polygon"] = o.Polygon
	}
	if o.Geocode != nil {
		toSerialize["geocode"] = o.Geocode
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	return json.Marshal(toSerialize)
}

type NullableAlertAtomEntry struct {
	value *AlertAtomEntry
	isSet bool
}

func (v NullableAlertAtomEntry) Get() *AlertAtomEntry {
	return v.value
}

func (v *NullableAlertAtomEntry) Set(val *AlertAtomEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertAtomEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertAtomEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertAtomEntry(val *AlertAtomEntry) *NullableAlertAtomEntry {
	return &NullableAlertAtomEntry{value: val, isSet: true}
}

func (v NullableAlertAtomEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertAtomEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


