# Go API client for weatherApi

weather.gov API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.8.3
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import weatherApi "github.com/jvanaartsen/go-weather-api"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), weatherApi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), weatherApi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), weatherApi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), weatherApi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.weather.gov*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**AlertsActive**](docs/DefaultApi.md#alertsactive) | **Get** /alerts/active | 
*DefaultApi* | [**AlertsActiveArea**](docs/DefaultApi.md#alertsactivearea) | **Get** /alerts/active/area/{area} | 
*DefaultApi* | [**AlertsActiveCount**](docs/DefaultApi.md#alertsactivecount) | **Get** /alerts/active/count | 
*DefaultApi* | [**AlertsActiveRegion**](docs/DefaultApi.md#alertsactiveregion) | **Get** /alerts/active/region/{region} | 
*DefaultApi* | [**AlertsActiveZone**](docs/DefaultApi.md#alertsactivezone) | **Get** /alerts/active/zone/{zoneId} | 
*DefaultApi* | [**AlertsQuery**](docs/DefaultApi.md#alertsquery) | **Get** /alerts | 
*DefaultApi* | [**AlertsSingle**](docs/DefaultApi.md#alertssingle) | **Get** /alerts/{id} | 
*DefaultApi* | [**AlertsTypes**](docs/DefaultApi.md#alertstypes) | **Get** /alerts/types | 
*DefaultApi* | [**Glossary**](docs/DefaultApi.md#glossary) | **Get** /glossary | 
*DefaultApi* | [**Gridpoint**](docs/DefaultApi.md#gridpoint) | **Get** /gridpoints/{wfo}/{x},{y} | 
*DefaultApi* | [**GridpointForecast**](docs/DefaultApi.md#gridpointforecast) | **Get** /gridpoints/{wfo}/{x},{y}/forecast | 
*DefaultApi* | [**GridpointForecastHourly**](docs/DefaultApi.md#gridpointforecasthourly) | **Get** /gridpoints/{wfo}/{x},{y}/forecast/hourly | 
*DefaultApi* | [**GridpointStations**](docs/DefaultApi.md#gridpointstations) | **Get** /gridpoints/{wfo}/{x},{y}/stations | 
*DefaultApi* | [**Icons**](docs/DefaultApi.md#icons) | **Get** /icons/{set}/{timeOfDay}/{first} | 
*DefaultApi* | [**IconsDualCondition**](docs/DefaultApi.md#iconsdualcondition) | **Get** /icons/{set}/{timeOfDay}/{first}/{second} | 
*DefaultApi* | [**IconsSummary**](docs/DefaultApi.md#iconssummary) | **Get** /icons | 
*DefaultApi* | [**LocationProducts**](docs/DefaultApi.md#locationproducts) | **Get** /products/locations/{locationId}/types | 
*DefaultApi* | [**ObsStation**](docs/DefaultApi.md#obsstation) | **Get** /stations/{stationId} | 
*DefaultApi* | [**ObsStations**](docs/DefaultApi.md#obsstations) | **Get** /stations | 
*DefaultApi* | [**Office**](docs/DefaultApi.md#office) | **Get** /offices/{officeId} | 
*DefaultApi* | [**OfficeHeadline**](docs/DefaultApi.md#officeheadline) | **Get** /offices/{officeId}/headlines/{headlineId} | 
*DefaultApi* | [**OfficeHeadlines**](docs/DefaultApi.md#officeheadlines) | **Get** /offices/{officeId}/headlines | 
*DefaultApi* | [**Point**](docs/DefaultApi.md#point) | **Get** /points/{point} | 
*DefaultApi* | [**PointStations**](docs/DefaultApi.md#pointstations) | **Get** /points/{point}/stations | 
*DefaultApi* | [**Product**](docs/DefaultApi.md#product) | **Get** /products/{productId} | 
*DefaultApi* | [**ProductLocations**](docs/DefaultApi.md#productlocations) | **Get** /products/locations | 
*DefaultApi* | [**ProductTypes**](docs/DefaultApi.md#producttypes) | **Get** /products/types | 
*DefaultApi* | [**ProductsQuery**](docs/DefaultApi.md#productsquery) | **Get** /products | 
*DefaultApi* | [**ProductsType**](docs/DefaultApi.md#productstype) | **Get** /products/types/{typeId} | 
*DefaultApi* | [**ProductsTypeLocation**](docs/DefaultApi.md#productstypelocation) | **Get** /products/types/{typeId}/locations/{locationId} | 
*DefaultApi* | [**ProductsTypeLocations**](docs/DefaultApi.md#productstypelocations) | **Get** /products/types/{typeId}/locations | 
*DefaultApi* | [**RadarProfiler**](docs/DefaultApi.md#radarprofiler) | **Get** /radar/profilers/{stationId} | 
*DefaultApi* | [**RadarQueue**](docs/DefaultApi.md#radarqueue) | **Get** /radar/queues/{host} | 
*DefaultApi* | [**RadarServer**](docs/DefaultApi.md#radarserver) | **Get** /radar/servers/{id} | 
*DefaultApi* | [**RadarServers**](docs/DefaultApi.md#radarservers) | **Get** /radar/servers | 
*DefaultApi* | [**RadarStation**](docs/DefaultApi.md#radarstation) | **Get** /radar/stations/{stationId} | 
*DefaultApi* | [**RadarStationAlarms**](docs/DefaultApi.md#radarstationalarms) | **Get** /radar/stations/{stationId}/alarms | 
*DefaultApi* | [**RadarStations**](docs/DefaultApi.md#radarstations) | **Get** /radar/stations | 
*DefaultApi* | [**SatelliteThumbnails**](docs/DefaultApi.md#satellitethumbnails) | **Get** /thumbnails/satellite/{area} | 
*DefaultApi* | [**StationObservationLatest**](docs/DefaultApi.md#stationobservationlatest) | **Get** /stations/{stationId}/observations/latest | 
*DefaultApi* | [**StationObservationList**](docs/DefaultApi.md#stationobservationlist) | **Get** /stations/{stationId}/observations | 
*DefaultApi* | [**StationObservationTime**](docs/DefaultApi.md#stationobservationtime) | **Get** /stations/{stationId}/observations/{time} | 
*DefaultApi* | [**Zone**](docs/DefaultApi.md#zone) | **Get** /zones/{type}/{zoneId} | 
*DefaultApi* | [**ZoneForecast**](docs/DefaultApi.md#zoneforecast) | **Get** /zones/{type}/{zoneId}/forecast | 
*DefaultApi* | [**ZoneList**](docs/DefaultApi.md#zonelist) | **Get** /zones | 
*DefaultApi* | [**ZoneListType**](docs/DefaultApi.md#zonelisttype) | **Get** /zones/{type} | 
*DefaultApi* | [**ZoneObs**](docs/DefaultApi.md#zoneobs) | **Get** /zones/forecast/{zoneId}/observations | 
*DefaultApi* | [**ZoneStations**](docs/DefaultApi.md#zonestations) | **Get** /zones/forecast/{zoneId}/stations | 


## Documentation For Models

 - [Alert](docs/Alert.md)
 - [AlertAtomEntry](docs/AlertAtomEntry.md)
 - [AlertAtomEntryAuthor](docs/AlertAtomEntryAuthor.md)
 - [AlertAtomFeed](docs/AlertAtomFeed.md)
 - [AlertAtomFeedAuthor](docs/AlertAtomFeedAuthor.md)
 - [AlertCertainty](docs/AlertCertainty.md)
 - [AlertCollection](docs/AlertCollection.md)
 - [AlertCollectionGeoJson](docs/AlertCollectionGeoJson.md)
 - [AlertCollectionGeoJsonAllOf](docs/AlertCollectionGeoJsonAllOf.md)
 - [AlertCollectionGeoJsonAllOfFeatures](docs/AlertCollectionGeoJsonAllOfFeatures.md)
 - [AlertCollectionJsonLd](docs/AlertCollectionJsonLd.md)
 - [AlertCollectionJsonLdAllOf](docs/AlertCollectionJsonLdAllOf.md)
 - [AlertCollectionPagination](docs/AlertCollectionPagination.md)
 - [AlertGeoJson](docs/AlertGeoJson.md)
 - [AlertGeocode](docs/AlertGeocode.md)
 - [AlertJsonLd](docs/AlertJsonLd.md)
 - [AlertMessageType](docs/AlertMessageType.md)
 - [AlertReferencesInner](docs/AlertReferencesInner.md)
 - [AlertSeverity](docs/AlertSeverity.md)
 - [AlertStatus](docs/AlertStatus.md)
 - [AlertUrgency](docs/AlertUrgency.md)
 - [AlertXMLParameter](docs/AlertXMLParameter.md)
 - [AlertsActiveCount200Response](docs/AlertsActiveCount200Response.md)
 - [AlertsTypes200Response](docs/AlertsTypes200Response.md)
 - [AreaCode](docs/AreaCode.md)
 - [GeoJSONLineString](docs/GeoJSONLineString.md)
 - [GeoJSONMultiLineString](docs/GeoJSONMultiLineString.md)
 - [GeoJSONMultiPoint](docs/GeoJSONMultiPoint.md)
 - [GeoJSONMultiPolygon](docs/GeoJSONMultiPolygon.md)
 - [GeoJSONPoint](docs/GeoJSONPoint.md)
 - [GeoJSONPolygon](docs/GeoJSONPolygon.md)
 - [GeoJsonFeature](docs/GeoJsonFeature.md)
 - [GeoJsonFeatureCollection](docs/GeoJsonFeatureCollection.md)
 - [GeoJsonGeometry](docs/GeoJsonGeometry.md)
 - [Glossary200Response](docs/Glossary200Response.md)
 - [Glossary200ResponseGlossaryInner](docs/Glossary200ResponseGlossaryInner.md)
 - [Gridpoint](docs/Gridpoint.md)
 - [GridpointForecast](docs/GridpointForecast.md)
 - [GridpointForecastGeoJson](docs/GridpointForecastGeoJson.md)
 - [GridpointForecastGeoJsonAllOf](docs/GridpointForecastGeoJsonAllOf.md)
 - [GridpointForecastJsonLd](docs/GridpointForecastJsonLd.md)
 - [GridpointForecastJsonLdAllOf](docs/GridpointForecastJsonLdAllOf.md)
 - [GridpointForecastPeriod](docs/GridpointForecastPeriod.md)
 - [GridpointForecastPeriodTemperature](docs/GridpointForecastPeriodTemperature.md)
 - [GridpointForecastPeriodWindGust](docs/GridpointForecastPeriodWindGust.md)
 - [GridpointForecastPeriodWindSpeed](docs/GridpointForecastPeriodWindSpeed.md)
 - [GridpointForecastUnits](docs/GridpointForecastUnits.md)
 - [GridpointGeoJson](docs/GridpointGeoJson.md)
 - [GridpointGeoJsonAllOf](docs/GridpointGeoJsonAllOf.md)
 - [GridpointHazards](docs/GridpointHazards.md)
 - [GridpointHazardsValuesInner](docs/GridpointHazardsValuesInner.md)
 - [GridpointHazardsValuesInnerValueInner](docs/GridpointHazardsValuesInnerValueInner.md)
 - [GridpointQuantitativeValueLayer](docs/GridpointQuantitativeValueLayer.md)
 - [GridpointQuantitativeValueLayerValuesInner](docs/GridpointQuantitativeValueLayerValuesInner.md)
 - [GridpointWeather](docs/GridpointWeather.md)
 - [GridpointWeatherValuesInner](docs/GridpointWeatherValuesInner.md)
 - [GridpointWeatherValuesInnerValueInner](docs/GridpointWeatherValuesInnerValueInner.md)
 - [ISO8601Interval](docs/ISO8601Interval.md)
 - [IconsSizeParameter](docs/IconsSizeParameter.md)
 - [IconsSizeParameterAnyOf](docs/IconsSizeParameterAnyOf.md)
 - [IconsSummary200Response](docs/IconsSummary200Response.md)
 - [IconsSummary200ResponseIconsValue](docs/IconsSummary200ResponseIconsValue.md)
 - [JsonLdContext](docs/JsonLdContext.md)
 - [LandRegionCode](docs/LandRegionCode.md)
 - [MarineAreaCode](docs/MarineAreaCode.md)
 - [MarineRegionCode](docs/MarineRegionCode.md)
 - [MetarPhenomenon](docs/MetarPhenomenon.md)
 - [MetarSkyCoverage](docs/MetarSkyCoverage.md)
 - [NWSForecastOfficeId](docs/NWSForecastOfficeId.md)
 - [NWSZoneType](docs/NWSZoneType.md)
 - [Observation](docs/Observation.md)
 - [ObservationCloudLayersInner](docs/ObservationCloudLayersInner.md)
 - [ObservationCollectionGeoJson](docs/ObservationCollectionGeoJson.md)
 - [ObservationCollectionGeoJsonAllOf](docs/ObservationCollectionGeoJsonAllOf.md)
 - [ObservationCollectionJsonLd](docs/ObservationCollectionJsonLd.md)
 - [ObservationGeoJson](docs/ObservationGeoJson.md)
 - [ObservationGeoJsonAllOf](docs/ObservationGeoJsonAllOf.md)
 - [ObservationStation](docs/ObservationStation.md)
 - [ObservationStationCollectionGeoJson](docs/ObservationStationCollectionGeoJson.md)
 - [ObservationStationCollectionGeoJsonAllOf](docs/ObservationStationCollectionGeoJsonAllOf.md)
 - [ObservationStationCollectionJsonLd](docs/ObservationStationCollectionJsonLd.md)
 - [ObservationStationGeoJson](docs/ObservationStationGeoJson.md)
 - [ObservationStationGeoJsonAllOf](docs/ObservationStationGeoJsonAllOf.md)
 - [ObservationStationJsonLd](docs/ObservationStationJsonLd.md)
 - [Office](docs/Office.md)
 - [OfficeAddress](docs/OfficeAddress.md)
 - [OfficeHeadline](docs/OfficeHeadline.md)
 - [OfficeHeadlineCollection](docs/OfficeHeadlineCollection.md)
 - [Point](docs/Point.md)
 - [PointGeoJson](docs/PointGeoJson.md)
 - [PointGeoJsonAllOf](docs/PointGeoJsonAllOf.md)
 - [PointJsonLd](docs/PointJsonLd.md)
 - [PointRelativeLocation](docs/PointRelativeLocation.md)
 - [ProblemDetail](docs/ProblemDetail.md)
 - [QuantitativeValue](docs/QuantitativeValue.md)
 - [RegionCode](docs/RegionCode.md)
 - [RelativeLocation](docs/RelativeLocation.md)
 - [RelativeLocationGeoJson](docs/RelativeLocationGeoJson.md)
 - [RelativeLocationGeoJsonAllOf](docs/RelativeLocationGeoJsonAllOf.md)
 - [RelativeLocationJsonLd](docs/RelativeLocationJsonLd.md)
 - [RelativeLocationJsonLdAllOf](docs/RelativeLocationJsonLdAllOf.md)
 - [StateTerritoryCode](docs/StateTerritoryCode.md)
 - [TextProduct](docs/TextProduct.md)
 - [TextProductCollection](docs/TextProductCollection.md)
 - [TextProductLocationCollection](docs/TextProductLocationCollection.md)
 - [TextProductTypeCollection](docs/TextProductTypeCollection.md)
 - [TextProductTypeCollectionGraphInner](docs/TextProductTypeCollectionGraphInner.md)
 - [Zone](docs/Zone.md)
 - [ZoneCollectionGeoJson](docs/ZoneCollectionGeoJson.md)
 - [ZoneCollectionGeoJsonAllOf](docs/ZoneCollectionGeoJsonAllOf.md)
 - [ZoneCollectionJsonLd](docs/ZoneCollectionJsonLd.md)
 - [ZoneForecast](docs/ZoneForecast.md)
 - [ZoneForecastGeoJson](docs/ZoneForecastGeoJson.md)
 - [ZoneForecastGeoJsonAllOf](docs/ZoneForecastGeoJsonAllOf.md)
 - [ZoneForecastPeriodsInner](docs/ZoneForecastPeriodsInner.md)
 - [ZoneGeoJson](docs/ZoneGeoJson.md)
 - [ZoneGeoJsonAllOf](docs/ZoneGeoJsonAllOf.md)
 - [ZoneState](docs/ZoneState.md)


## Documentation For Authorization



### userAgent

- **Type**: API key
- **API key parameter name**: User-Agent
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: User-Agent and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



