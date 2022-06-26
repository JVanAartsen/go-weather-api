# GeoJSONPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Coordinates** | **[]float32** | A GeoJSON coordinate. Please refer to IETF RFC 7946 for information on the GeoJSON format. | 
**Bbox** | Pointer to **[]float32** | A GeoJSON bounding box. Please refer to IETF RFC 7946 for information on the GeoJSON format. | [optional] 

## Methods

### NewGeoJSONPoint

`func NewGeoJSONPoint(type_ string, coordinates []float32, ) *GeoJSONPoint`

NewGeoJSONPoint instantiates a new GeoJSONPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoJSONPointWithDefaults

`func NewGeoJSONPointWithDefaults() *GeoJSONPoint`

NewGeoJSONPointWithDefaults instantiates a new GeoJSONPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GeoJSONPoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoJSONPoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoJSONPoint) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *GeoJSONPoint) GetCoordinates() []float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *GeoJSONPoint) GetCoordinatesOk() (*[]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *GeoJSONPoint) SetCoordinates(v []float32)`

SetCoordinates sets Coordinates field to given value.


### GetBbox

`func (o *GeoJSONPoint) GetBbox() []float32`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *GeoJSONPoint) GetBboxOk() (*[]float32, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *GeoJSONPoint) SetBbox(v []float32)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *GeoJSONPoint) HasBbox() bool`

HasBbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

