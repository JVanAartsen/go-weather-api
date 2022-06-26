# Office

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**OfficeAddress**](OfficeAddress.md) |  | [optional] 
**Telephone** | Pointer to **string** |  | [optional] 
**FaxNumber** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**SameAs** | Pointer to **string** |  | [optional] 
**NwsRegion** | Pointer to **string** |  | [optional] 
**ParentOrganization** | Pointer to **string** |  | [optional] 
**ResponsibleCounties** | Pointer to **[]string** |  | [optional] 
**ResponsibleForecastZones** | Pointer to **[]string** |  | [optional] 
**ResponsibleFireZones** | Pointer to **[]string** |  | [optional] 
**ApprovedObservationStations** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOffice

`func NewOffice() *Office`

NewOffice instantiates a new Office object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfficeWithDefaults

`func NewOfficeWithDefaults() *Office`

NewOfficeWithDefaults instantiates a new Office object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Office) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Office) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Office) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *Office) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetType

`func (o *Office) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Office) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Office) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Office) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Office) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Office) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Office) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Office) HasId() bool`

HasId returns a boolean if a field has been set.

### GetId

`func (o *Office) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Office) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Office) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Office) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Office) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Office) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Office) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Office) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *Office) GetAddress() OfficeAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Office) GetAddressOk() (*OfficeAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Office) SetAddress(v OfficeAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Office) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTelephone

`func (o *Office) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *Office) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *Office) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephone

`func (o *Office) HasTelephone() bool`

HasTelephone returns a boolean if a field has been set.

### GetFaxNumber

`func (o *Office) GetFaxNumber() string`

GetFaxNumber returns the FaxNumber field if non-nil, zero value otherwise.

### GetFaxNumberOk

`func (o *Office) GetFaxNumberOk() (*string, bool)`

GetFaxNumberOk returns a tuple with the FaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNumber

`func (o *Office) SetFaxNumber(v string)`

SetFaxNumber sets FaxNumber field to given value.

### HasFaxNumber

`func (o *Office) HasFaxNumber() bool`

HasFaxNumber returns a boolean if a field has been set.

### GetEmail

`func (o *Office) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Office) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Office) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Office) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSameAs

`func (o *Office) GetSameAs() string`

GetSameAs returns the SameAs field if non-nil, zero value otherwise.

### GetSameAsOk

`func (o *Office) GetSameAsOk() (*string, bool)`

GetSameAsOk returns a tuple with the SameAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameAs

`func (o *Office) SetSameAs(v string)`

SetSameAs sets SameAs field to given value.

### HasSameAs

`func (o *Office) HasSameAs() bool`

HasSameAs returns a boolean if a field has been set.

### GetNwsRegion

`func (o *Office) GetNwsRegion() string`

GetNwsRegion returns the NwsRegion field if non-nil, zero value otherwise.

### GetNwsRegionOk

`func (o *Office) GetNwsRegionOk() (*string, bool)`

GetNwsRegionOk returns a tuple with the NwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNwsRegion

`func (o *Office) SetNwsRegion(v string)`

SetNwsRegion sets NwsRegion field to given value.

### HasNwsRegion

`func (o *Office) HasNwsRegion() bool`

HasNwsRegion returns a boolean if a field has been set.

### GetParentOrganization

`func (o *Office) GetParentOrganization() string`

GetParentOrganization returns the ParentOrganization field if non-nil, zero value otherwise.

### GetParentOrganizationOk

`func (o *Office) GetParentOrganizationOk() (*string, bool)`

GetParentOrganizationOk returns a tuple with the ParentOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrganization

`func (o *Office) SetParentOrganization(v string)`

SetParentOrganization sets ParentOrganization field to given value.

### HasParentOrganization

`func (o *Office) HasParentOrganization() bool`

HasParentOrganization returns a boolean if a field has been set.

### GetResponsibleCounties

`func (o *Office) GetResponsibleCounties() []string`

GetResponsibleCounties returns the ResponsibleCounties field if non-nil, zero value otherwise.

### GetResponsibleCountiesOk

`func (o *Office) GetResponsibleCountiesOk() (*[]string, bool)`

GetResponsibleCountiesOk returns a tuple with the ResponsibleCounties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleCounties

`func (o *Office) SetResponsibleCounties(v []string)`

SetResponsibleCounties sets ResponsibleCounties field to given value.

### HasResponsibleCounties

`func (o *Office) HasResponsibleCounties() bool`

HasResponsibleCounties returns a boolean if a field has been set.

### GetResponsibleForecastZones

`func (o *Office) GetResponsibleForecastZones() []string`

GetResponsibleForecastZones returns the ResponsibleForecastZones field if non-nil, zero value otherwise.

### GetResponsibleForecastZonesOk

`func (o *Office) GetResponsibleForecastZonesOk() (*[]string, bool)`

GetResponsibleForecastZonesOk returns a tuple with the ResponsibleForecastZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleForecastZones

`func (o *Office) SetResponsibleForecastZones(v []string)`

SetResponsibleForecastZones sets ResponsibleForecastZones field to given value.

### HasResponsibleForecastZones

`func (o *Office) HasResponsibleForecastZones() bool`

HasResponsibleForecastZones returns a boolean if a field has been set.

### GetResponsibleFireZones

`func (o *Office) GetResponsibleFireZones() []string`

GetResponsibleFireZones returns the ResponsibleFireZones field if non-nil, zero value otherwise.

### GetResponsibleFireZonesOk

`func (o *Office) GetResponsibleFireZonesOk() (*[]string, bool)`

GetResponsibleFireZonesOk returns a tuple with the ResponsibleFireZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibleFireZones

`func (o *Office) SetResponsibleFireZones(v []string)`

SetResponsibleFireZones sets ResponsibleFireZones field to given value.

### HasResponsibleFireZones

`func (o *Office) HasResponsibleFireZones() bool`

HasResponsibleFireZones returns a boolean if a field has been set.

### GetApprovedObservationStations

`func (o *Office) GetApprovedObservationStations() []string`

GetApprovedObservationStations returns the ApprovedObservationStations field if non-nil, zero value otherwise.

### GetApprovedObservationStationsOk

`func (o *Office) GetApprovedObservationStationsOk() (*[]string, bool)`

GetApprovedObservationStationsOk returns a tuple with the ApprovedObservationStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedObservationStations

`func (o *Office) SetApprovedObservationStations(v []string)`

SetApprovedObservationStations sets ApprovedObservationStations field to given value.

### HasApprovedObservationStations

`func (o *Office) HasApprovedObservationStations() bool`

HasApprovedObservationStations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


