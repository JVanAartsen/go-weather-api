# AlertsActiveCount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | The total number of active alerts | [optional] 
**Land** | Pointer to **int32** | The total number of active alerts affecting land zones | [optional] 
**Marine** | Pointer to **int32** | The total number of active alerts affecting marine zones | [optional] 
**Regions** | Pointer to **map[string]int32** | Active alerts by marine region | [optional] 
**Areas** | Pointer to **map[string]int32** | Active alerts by area (state/territory) | [optional] 
**Zones** | Pointer to **map[string]int32** | Active alerts by NWS public zone or county code | [optional] 

## Methods

### NewAlertsActiveCount200Response

`func NewAlertsActiveCount200Response() *AlertsActiveCount200Response`

NewAlertsActiveCount200Response instantiates a new AlertsActiveCount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsActiveCount200ResponseWithDefaults

`func NewAlertsActiveCount200ResponseWithDefaults() *AlertsActiveCount200Response`

NewAlertsActiveCount200ResponseWithDefaults instantiates a new AlertsActiveCount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AlertsActiveCount200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AlertsActiveCount200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AlertsActiveCount200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AlertsActiveCount200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLand

`func (o *AlertsActiveCount200Response) GetLand() int32`

GetLand returns the Land field if non-nil, zero value otherwise.

### GetLandOk

`func (o *AlertsActiveCount200Response) GetLandOk() (*int32, bool)`

GetLandOk returns a tuple with the Land field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLand

`func (o *AlertsActiveCount200Response) SetLand(v int32)`

SetLand sets Land field to given value.

### HasLand

`func (o *AlertsActiveCount200Response) HasLand() bool`

HasLand returns a boolean if a field has been set.

### GetMarine

`func (o *AlertsActiveCount200Response) GetMarine() int32`

GetMarine returns the Marine field if non-nil, zero value otherwise.

### GetMarineOk

`func (o *AlertsActiveCount200Response) GetMarineOk() (*int32, bool)`

GetMarineOk returns a tuple with the Marine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarine

`func (o *AlertsActiveCount200Response) SetMarine(v int32)`

SetMarine sets Marine field to given value.

### HasMarine

`func (o *AlertsActiveCount200Response) HasMarine() bool`

HasMarine returns a boolean if a field has been set.

### GetRegions

`func (o *AlertsActiveCount200Response) GetRegions() map[string]int32`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *AlertsActiveCount200Response) GetRegionsOk() (*map[string]int32, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *AlertsActiveCount200Response) SetRegions(v map[string]int32)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *AlertsActiveCount200Response) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetAreas

`func (o *AlertsActiveCount200Response) GetAreas() map[string]int32`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *AlertsActiveCount200Response) GetAreasOk() (*map[string]int32, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *AlertsActiveCount200Response) SetAreas(v map[string]int32)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *AlertsActiveCount200Response) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### GetZones

`func (o *AlertsActiveCount200Response) GetZones() map[string]int32`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AlertsActiveCount200Response) GetZonesOk() (*map[string]int32, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AlertsActiveCount200Response) SetZones(v map[string]int32)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AlertsActiveCount200Response) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


