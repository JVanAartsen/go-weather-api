# Observation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JsonLdContext**](JsonLdContext.md) |  | [optional] 
**Geometry** | Pointer to **NullableString** | A geometry represented in Well-Known Text (WKT) format. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Elevation** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**Station** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**RawMessage** | Pointer to **string** |  | [optional] 
**TextDescription** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**PresentWeather** | Pointer to [**[]MetarPhenomenon**](MetarPhenomenon.md) |  | [optional] 
**Temperature** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**Dewpoint** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**WindDirection** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**WindSpeed** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**WindGust** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**BarometricPressure** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**SeaLevelPressure** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**Visibility** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**MaxTemperatureLast24Hours** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**MinTemperatureLast24Hours** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**PrecipitationLastHour** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**PrecipitationLast3Hours** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**PrecipitationLast6Hours** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**RelativeHumidity** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**WindChill** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**HeatIndex** | Pointer to [**QuantitativeValue**](QuantitativeValue.md) |  | [optional] 
**CloudLayers** | Pointer to [**[]ObservationCloudLayersInner**](ObservationCloudLayersInner.md) |  | [optional] 

## Methods

### NewObservation

`func NewObservation() *Observation`

NewObservation instantiates a new Observation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservationWithDefaults

`func NewObservationWithDefaults() *Observation`

NewObservationWithDefaults instantiates a new Observation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *Observation) GetContext() JsonLdContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Observation) GetContextOk() (*JsonLdContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Observation) SetContext(v JsonLdContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *Observation) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetGeometry

`func (o *Observation) GetGeometry() string`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *Observation) GetGeometryOk() (*string, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *Observation) SetGeometry(v string)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *Observation) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### SetGeometryNil

`func (o *Observation) SetGeometryNil(b bool)`

 SetGeometryNil sets the value for Geometry to be an explicit nil

### UnsetGeometry
`func (o *Observation) UnsetGeometry()`

UnsetGeometry ensures that no value is present for Geometry, not even an explicit nil
### GetId

`func (o *Observation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Observation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Observation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Observation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Observation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Observation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Observation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Observation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetElevation

`func (o *Observation) GetElevation() QuantitativeValue`

GetElevation returns the Elevation field if non-nil, zero value otherwise.

### GetElevationOk

`func (o *Observation) GetElevationOk() (*QuantitativeValue, bool)`

GetElevationOk returns a tuple with the Elevation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevation

`func (o *Observation) SetElevation(v QuantitativeValue)`

SetElevation sets Elevation field to given value.

### HasElevation

`func (o *Observation) HasElevation() bool`

HasElevation returns a boolean if a field has been set.

### GetStation

`func (o *Observation) GetStation() string`

GetStation returns the Station field if non-nil, zero value otherwise.

### GetStationOk

`func (o *Observation) GetStationOk() (*string, bool)`

GetStationOk returns a tuple with the Station field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStation

`func (o *Observation) SetStation(v string)`

SetStation sets Station field to given value.

### HasStation

`func (o *Observation) HasStation() bool`

HasStation returns a boolean if a field has been set.

### GetTimestamp

`func (o *Observation) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Observation) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Observation) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Observation) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRawMessage

`func (o *Observation) GetRawMessage() string`

GetRawMessage returns the RawMessage field if non-nil, zero value otherwise.

### GetRawMessageOk

`func (o *Observation) GetRawMessageOk() (*string, bool)`

GetRawMessageOk returns a tuple with the RawMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawMessage

`func (o *Observation) SetRawMessage(v string)`

SetRawMessage sets RawMessage field to given value.

### HasRawMessage

`func (o *Observation) HasRawMessage() bool`

HasRawMessage returns a boolean if a field has been set.

### GetTextDescription

`func (o *Observation) GetTextDescription() string`

GetTextDescription returns the TextDescription field if non-nil, zero value otherwise.

### GetTextDescriptionOk

`func (o *Observation) GetTextDescriptionOk() (*string, bool)`

GetTextDescriptionOk returns a tuple with the TextDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextDescription

`func (o *Observation) SetTextDescription(v string)`

SetTextDescription sets TextDescription field to given value.

### HasTextDescription

`func (o *Observation) HasTextDescription() bool`

HasTextDescription returns a boolean if a field has been set.

### GetIcon

`func (o *Observation) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Observation) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Observation) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Observation) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetPresentWeather

`func (o *Observation) GetPresentWeather() []MetarPhenomenon`

GetPresentWeather returns the PresentWeather field if non-nil, zero value otherwise.

### GetPresentWeatherOk

`func (o *Observation) GetPresentWeatherOk() (*[]MetarPhenomenon, bool)`

GetPresentWeatherOk returns a tuple with the PresentWeather field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentWeather

`func (o *Observation) SetPresentWeather(v []MetarPhenomenon)`

SetPresentWeather sets PresentWeather field to given value.

### HasPresentWeather

`func (o *Observation) HasPresentWeather() bool`

HasPresentWeather returns a boolean if a field has been set.

### GetTemperature

`func (o *Observation) GetTemperature() QuantitativeValue`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *Observation) GetTemperatureOk() (*QuantitativeValue, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *Observation) SetTemperature(v QuantitativeValue)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *Observation) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetDewpoint

`func (o *Observation) GetDewpoint() QuantitativeValue`

GetDewpoint returns the Dewpoint field if non-nil, zero value otherwise.

### GetDewpointOk

`func (o *Observation) GetDewpointOk() (*QuantitativeValue, bool)`

GetDewpointOk returns a tuple with the Dewpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDewpoint

`func (o *Observation) SetDewpoint(v QuantitativeValue)`

SetDewpoint sets Dewpoint field to given value.

### HasDewpoint

`func (o *Observation) HasDewpoint() bool`

HasDewpoint returns a boolean if a field has been set.

### GetWindDirection

`func (o *Observation) GetWindDirection() QuantitativeValue`

GetWindDirection returns the WindDirection field if non-nil, zero value otherwise.

### GetWindDirectionOk

`func (o *Observation) GetWindDirectionOk() (*QuantitativeValue, bool)`

GetWindDirectionOk returns a tuple with the WindDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindDirection

`func (o *Observation) SetWindDirection(v QuantitativeValue)`

SetWindDirection sets WindDirection field to given value.

### HasWindDirection

`func (o *Observation) HasWindDirection() bool`

HasWindDirection returns a boolean if a field has been set.

### GetWindSpeed

`func (o *Observation) GetWindSpeed() QuantitativeValue`

GetWindSpeed returns the WindSpeed field if non-nil, zero value otherwise.

### GetWindSpeedOk

`func (o *Observation) GetWindSpeedOk() (*QuantitativeValue, bool)`

GetWindSpeedOk returns a tuple with the WindSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindSpeed

`func (o *Observation) SetWindSpeed(v QuantitativeValue)`

SetWindSpeed sets WindSpeed field to given value.

### HasWindSpeed

`func (o *Observation) HasWindSpeed() bool`

HasWindSpeed returns a boolean if a field has been set.

### GetWindGust

`func (o *Observation) GetWindGust() QuantitativeValue`

GetWindGust returns the WindGust field if non-nil, zero value otherwise.

### GetWindGustOk

`func (o *Observation) GetWindGustOk() (*QuantitativeValue, bool)`

GetWindGustOk returns a tuple with the WindGust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindGust

`func (o *Observation) SetWindGust(v QuantitativeValue)`

SetWindGust sets WindGust field to given value.

### HasWindGust

`func (o *Observation) HasWindGust() bool`

HasWindGust returns a boolean if a field has been set.

### GetBarometricPressure

`func (o *Observation) GetBarometricPressure() QuantitativeValue`

GetBarometricPressure returns the BarometricPressure field if non-nil, zero value otherwise.

### GetBarometricPressureOk

`func (o *Observation) GetBarometricPressureOk() (*QuantitativeValue, bool)`

GetBarometricPressureOk returns a tuple with the BarometricPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarometricPressure

`func (o *Observation) SetBarometricPressure(v QuantitativeValue)`

SetBarometricPressure sets BarometricPressure field to given value.

### HasBarometricPressure

`func (o *Observation) HasBarometricPressure() bool`

HasBarometricPressure returns a boolean if a field has been set.

### GetSeaLevelPressure

`func (o *Observation) GetSeaLevelPressure() QuantitativeValue`

GetSeaLevelPressure returns the SeaLevelPressure field if non-nil, zero value otherwise.

### GetSeaLevelPressureOk

`func (o *Observation) GetSeaLevelPressureOk() (*QuantitativeValue, bool)`

GetSeaLevelPressureOk returns a tuple with the SeaLevelPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeaLevelPressure

`func (o *Observation) SetSeaLevelPressure(v QuantitativeValue)`

SetSeaLevelPressure sets SeaLevelPressure field to given value.

### HasSeaLevelPressure

`func (o *Observation) HasSeaLevelPressure() bool`

HasSeaLevelPressure returns a boolean if a field has been set.

### GetVisibility

`func (o *Observation) GetVisibility() QuantitativeValue`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Observation) GetVisibilityOk() (*QuantitativeValue, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Observation) SetVisibility(v QuantitativeValue)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Observation) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMaxTemperatureLast24Hours

`func (o *Observation) GetMaxTemperatureLast24Hours() QuantitativeValue`

GetMaxTemperatureLast24Hours returns the MaxTemperatureLast24Hours field if non-nil, zero value otherwise.

### GetMaxTemperatureLast24HoursOk

`func (o *Observation) GetMaxTemperatureLast24HoursOk() (*QuantitativeValue, bool)`

GetMaxTemperatureLast24HoursOk returns a tuple with the MaxTemperatureLast24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTemperatureLast24Hours

`func (o *Observation) SetMaxTemperatureLast24Hours(v QuantitativeValue)`

SetMaxTemperatureLast24Hours sets MaxTemperatureLast24Hours field to given value.

### HasMaxTemperatureLast24Hours

`func (o *Observation) HasMaxTemperatureLast24Hours() bool`

HasMaxTemperatureLast24Hours returns a boolean if a field has been set.

### GetMinTemperatureLast24Hours

`func (o *Observation) GetMinTemperatureLast24Hours() QuantitativeValue`

GetMinTemperatureLast24Hours returns the MinTemperatureLast24Hours field if non-nil, zero value otherwise.

### GetMinTemperatureLast24HoursOk

`func (o *Observation) GetMinTemperatureLast24HoursOk() (*QuantitativeValue, bool)`

GetMinTemperatureLast24HoursOk returns a tuple with the MinTemperatureLast24Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTemperatureLast24Hours

`func (o *Observation) SetMinTemperatureLast24Hours(v QuantitativeValue)`

SetMinTemperatureLast24Hours sets MinTemperatureLast24Hours field to given value.

### HasMinTemperatureLast24Hours

`func (o *Observation) HasMinTemperatureLast24Hours() bool`

HasMinTemperatureLast24Hours returns a boolean if a field has been set.

### GetPrecipitationLastHour

`func (o *Observation) GetPrecipitationLastHour() QuantitativeValue`

GetPrecipitationLastHour returns the PrecipitationLastHour field if non-nil, zero value otherwise.

### GetPrecipitationLastHourOk

`func (o *Observation) GetPrecipitationLastHourOk() (*QuantitativeValue, bool)`

GetPrecipitationLastHourOk returns a tuple with the PrecipitationLastHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecipitationLastHour

`func (o *Observation) SetPrecipitationLastHour(v QuantitativeValue)`

SetPrecipitationLastHour sets PrecipitationLastHour field to given value.

### HasPrecipitationLastHour

`func (o *Observation) HasPrecipitationLastHour() bool`

HasPrecipitationLastHour returns a boolean if a field has been set.

### GetPrecipitationLast3Hours

`func (o *Observation) GetPrecipitationLast3Hours() QuantitativeValue`

GetPrecipitationLast3Hours returns the PrecipitationLast3Hours field if non-nil, zero value otherwise.

### GetPrecipitationLast3HoursOk

`func (o *Observation) GetPrecipitationLast3HoursOk() (*QuantitativeValue, bool)`

GetPrecipitationLast3HoursOk returns a tuple with the PrecipitationLast3Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecipitationLast3Hours

`func (o *Observation) SetPrecipitationLast3Hours(v QuantitativeValue)`

SetPrecipitationLast3Hours sets PrecipitationLast3Hours field to given value.

### HasPrecipitationLast3Hours

`func (o *Observation) HasPrecipitationLast3Hours() bool`

HasPrecipitationLast3Hours returns a boolean if a field has been set.

### GetPrecipitationLast6Hours

`func (o *Observation) GetPrecipitationLast6Hours() QuantitativeValue`

GetPrecipitationLast6Hours returns the PrecipitationLast6Hours field if non-nil, zero value otherwise.

### GetPrecipitationLast6HoursOk

`func (o *Observation) GetPrecipitationLast6HoursOk() (*QuantitativeValue, bool)`

GetPrecipitationLast6HoursOk returns a tuple with the PrecipitationLast6Hours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecipitationLast6Hours

`func (o *Observation) SetPrecipitationLast6Hours(v QuantitativeValue)`

SetPrecipitationLast6Hours sets PrecipitationLast6Hours field to given value.

### HasPrecipitationLast6Hours

`func (o *Observation) HasPrecipitationLast6Hours() bool`

HasPrecipitationLast6Hours returns a boolean if a field has been set.

### GetRelativeHumidity

`func (o *Observation) GetRelativeHumidity() QuantitativeValue`

GetRelativeHumidity returns the RelativeHumidity field if non-nil, zero value otherwise.

### GetRelativeHumidityOk

`func (o *Observation) GetRelativeHumidityOk() (*QuantitativeValue, bool)`

GetRelativeHumidityOk returns a tuple with the RelativeHumidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeHumidity

`func (o *Observation) SetRelativeHumidity(v QuantitativeValue)`

SetRelativeHumidity sets RelativeHumidity field to given value.

### HasRelativeHumidity

`func (o *Observation) HasRelativeHumidity() bool`

HasRelativeHumidity returns a boolean if a field has been set.

### GetWindChill

`func (o *Observation) GetWindChill() QuantitativeValue`

GetWindChill returns the WindChill field if non-nil, zero value otherwise.

### GetWindChillOk

`func (o *Observation) GetWindChillOk() (*QuantitativeValue, bool)`

GetWindChillOk returns a tuple with the WindChill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindChill

`func (o *Observation) SetWindChill(v QuantitativeValue)`

SetWindChill sets WindChill field to given value.

### HasWindChill

`func (o *Observation) HasWindChill() bool`

HasWindChill returns a boolean if a field has been set.

### GetHeatIndex

`func (o *Observation) GetHeatIndex() QuantitativeValue`

GetHeatIndex returns the HeatIndex field if non-nil, zero value otherwise.

### GetHeatIndexOk

`func (o *Observation) GetHeatIndexOk() (*QuantitativeValue, bool)`

GetHeatIndexOk returns a tuple with the HeatIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatIndex

`func (o *Observation) SetHeatIndex(v QuantitativeValue)`

SetHeatIndex sets HeatIndex field to given value.

### HasHeatIndex

`func (o *Observation) HasHeatIndex() bool`

HasHeatIndex returns a boolean if a field has been set.

### GetCloudLayers

`func (o *Observation) GetCloudLayers() []ObservationCloudLayersInner`

GetCloudLayers returns the CloudLayers field if non-nil, zero value otherwise.

### GetCloudLayersOk

`func (o *Observation) GetCloudLayersOk() (*[]ObservationCloudLayersInner, bool)`

GetCloudLayersOk returns a tuple with the CloudLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudLayers

`func (o *Observation) SetCloudLayers(v []ObservationCloudLayersInner)`

SetCloudLayers sets CloudLayers field to given value.

### HasCloudLayers

`func (o *Observation) HasCloudLayers() bool`

HasCloudLayers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


