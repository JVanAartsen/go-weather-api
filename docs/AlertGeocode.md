# AlertGeocode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UGC** | Pointer to **[]string** | A list of NWS public zone or county identifiers. | [optional] 
**SAME** | Pointer to **[]string** | A list of SAME (Specific Area Message Encoding) codes for affected counties. | [optional] 

## Methods

### NewAlertGeocode

`func NewAlertGeocode() *AlertGeocode`

NewAlertGeocode instantiates a new AlertGeocode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertGeocodeWithDefaults

`func NewAlertGeocodeWithDefaults() *AlertGeocode`

NewAlertGeocodeWithDefaults instantiates a new AlertGeocode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUGC

`func (o *AlertGeocode) GetUGC() []string`

GetUGC returns the UGC field if non-nil, zero value otherwise.

### GetUGCOk

`func (o *AlertGeocode) GetUGCOk() (*[]string, bool)`

GetUGCOk returns a tuple with the UGC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUGC

`func (o *AlertGeocode) SetUGC(v []string)`

SetUGC sets UGC field to given value.

### HasUGC

`func (o *AlertGeocode) HasUGC() bool`

HasUGC returns a boolean if a field has been set.

### GetSAME

`func (o *AlertGeocode) GetSAME() []string`

GetSAME returns the SAME field if non-nil, zero value otherwise.

### GetSAMEOk

`func (o *AlertGeocode) GetSAMEOk() (*[]string, bool)`

GetSAMEOk returns a tuple with the SAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAME

`func (o *AlertGeocode) SetSAME(v []string)`

SetSAME sets SAME field to given value.

### HasSAME

`func (o *AlertGeocode) HasSAME() bool`

HasSAME returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


