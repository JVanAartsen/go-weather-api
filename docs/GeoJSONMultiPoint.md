# GeoJSONMultiPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Coordinates** | **[][]float32** |  | 
**Bbox** | Pointer to **[]float32** | A GeoJSON bounding box. Please refer to IETF RFC 7946 for information on the GeoJSON format. | [optional] 

## Methods

### NewGeoJSONMultiPoint

`func NewGeoJSONMultiPoint(type_ string, coordinates [][]float32, ) *GeoJSONMultiPoint`

NewGeoJSONMultiPoint instantiates a new GeoJSONMultiPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoJSONMultiPointWithDefaults

`func NewGeoJSONMultiPointWithDefaults() *GeoJSONMultiPoint`

NewGeoJSONMultiPointWithDefaults instantiates a new GeoJSONMultiPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GeoJSONMultiPoint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoJSONMultiPoint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoJSONMultiPoint) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *GeoJSONMultiPoint) GetCoordinates() [][]float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *GeoJSONMultiPoint) GetCoordinatesOk() (*[][]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *GeoJSONMultiPoint) SetCoordinates(v [][]float32)`

SetCoordinates sets Coordinates field to given value.


### GetBbox

`func (o *GeoJSONMultiPoint) GetBbox() []float32`

GetBbox returns the Bbox field if non-nil, zero value otherwise.

### GetBboxOk

`func (o *GeoJSONMultiPoint) GetBboxOk() (*[]float32, bool)`

GetBboxOk returns a tuple with the Bbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBbox

`func (o *GeoJSONMultiPoint) SetBbox(v []float32)`

SetBbox sets Bbox field to given value.

### HasBbox

`func (o *GeoJSONMultiPoint) HasBbox() bool`

HasBbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


