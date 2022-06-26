# ObservationStationCollectionGeoJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Type** | **string** |  | 
**Features** | [**[]ObservationStationGeoJsonAllOf**](ObservationStationGeoJsonAllOf.md) |  | 
**ObservationStations** | Pointer to **[]string** |  | [optional] 

## Methods

### NewObservationStationCollectionGeoJson

`func NewObservationStationCollectionGeoJson(type_ string, features []ObservationStationGeoJsonAllOf, ) *ObservationStationCollectionGeoJson`

NewObservationStationCollectionGeoJson instantiates a new ObservationStationCollectionGeoJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationStationCollectionGeoJsonWithDefaults

`func NewObservationStationCollectionGeoJsonWithDefaults() *ObservationStationCollectionGeoJson`

NewObservationStationCollectionGeoJsonWithDefaults instantiates a new ObservationStationCollectionGeoJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ObservationStationCollectionGeoJson) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ObservationStationCollectionGeoJson) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ObservationStationCollectionGeoJson) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *ObservationStationCollectionGeoJson) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetType

`func (o *ObservationStationCollectionGeoJson) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ObservationStationCollectionGeoJson) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ObservationStationCollectionGeoJson) SetType(v string)`

SetType sets Type field to given value.


### GetFeatures

`func (o *ObservationStationCollectionGeoJson) GetFeatures() []ObservationStationGeoJsonAllOf`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ObservationStationCollectionGeoJson) GetFeaturesOk() (*[]ObservationStationGeoJsonAllOf, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ObservationStationCollectionGeoJson) SetFeatures(v []ObservationStationGeoJsonAllOf)`

SetFeatures sets Features field to given value.


### GetObservationStations

`func (o *ObservationStationCollectionGeoJson) GetObservationStations() []string`

GetObservationStations returns the ObservationStations field if non-nil, zero value otherwise.

### GetObservationStationsOk

`func (o *ObservationStationCollectionGeoJson) GetObservationStationsOk() (*[]string, bool)`

GetObservationStationsOk returns a tuple with the ObservationStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationStations

`func (o *ObservationStationCollectionGeoJson) SetObservationStations(v []string)`

SetObservationStations sets ObservationStations field to given value.

### HasObservationStations

`func (o *ObservationStationCollectionGeoJson) HasObservationStations() bool`

HasObservationStations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


