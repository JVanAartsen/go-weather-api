# \DefaultApi

All URIs are relative to *https://api.weather.gov*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AlertsActive**](DefaultApi.md#AlertsActive) | **Get** /alerts/active | 
[**AlertsActiveArea**](DefaultApi.md#AlertsActiveArea) | **Get** /alerts/active/area/{area} | 
[**AlertsActiveCount**](DefaultApi.md#AlertsActiveCount) | **Get** /alerts/active/count | 
[**AlertsActiveRegion**](DefaultApi.md#AlertsActiveRegion) | **Get** /alerts/active/region/{region} | 
[**AlertsActiveZone**](DefaultApi.md#AlertsActiveZone) | **Get** /alerts/active/zone/{zoneId} | 
[**AlertsQuery**](DefaultApi.md#AlertsQuery) | **Get** /alerts | 
[**AlertsSingle**](DefaultApi.md#AlertsSingle) | **Get** /alerts/{id} | 
[**AlertsTypes**](DefaultApi.md#AlertsTypes) | **Get** /alerts/types | 
[**Glossary**](DefaultApi.md#Glossary) | **Get** /glossary | 
[**Gridpoint**](DefaultApi.md#Gridpoint) | **Get** /gridpoints/{wfo}/{x},{y} | 
[**GridpointForecast**](DefaultApi.md#GridpointForecast) | **Get** /gridpoints/{wfo}/{x},{y}/forecast | 
[**GridpointForecastHourly**](DefaultApi.md#GridpointForecastHourly) | **Get** /gridpoints/{wfo}/{x},{y}/forecast/hourly | 
[**GridpointStations**](DefaultApi.md#GridpointStations) | **Get** /gridpoints/{wfo}/{x},{y}/stations | 
[**Icons**](DefaultApi.md#Icons) | **Get** /icons/{set}/{timeOfDay}/{first} | 
[**IconsDualCondition**](DefaultApi.md#IconsDualCondition) | **Get** /icons/{set}/{timeOfDay}/{first}/{second} | 
[**IconsSummary**](DefaultApi.md#IconsSummary) | **Get** /icons | 
[**LocationProducts**](DefaultApi.md#LocationProducts) | **Get** /products/locations/{locationId}/types | 
[**ObsStation**](DefaultApi.md#ObsStation) | **Get** /stations/{stationId} | 
[**ObsStations**](DefaultApi.md#ObsStations) | **Get** /stations | 
[**Office**](DefaultApi.md#Office) | **Get** /offices/{officeId} | 
[**OfficeHeadline**](DefaultApi.md#OfficeHeadline) | **Get** /offices/{officeId}/headlines/{headlineId} | 
[**OfficeHeadlines**](DefaultApi.md#OfficeHeadlines) | **Get** /offices/{officeId}/headlines | 
[**Point**](DefaultApi.md#Point) | **Get** /points/{point} | 
[**PointStations**](DefaultApi.md#PointStations) | **Get** /points/{point}/stations | 
[**Product**](DefaultApi.md#Product) | **Get** /products/{productId} | 
[**ProductLocations**](DefaultApi.md#ProductLocations) | **Get** /products/locations | 
[**ProductTypes**](DefaultApi.md#ProductTypes) | **Get** /products/types | 
[**ProductsQuery**](DefaultApi.md#ProductsQuery) | **Get** /products | 
[**ProductsType**](DefaultApi.md#ProductsType) | **Get** /products/types/{typeId} | 
[**ProductsTypeLocation**](DefaultApi.md#ProductsTypeLocation) | **Get** /products/types/{typeId}/locations/{locationId} | 
[**ProductsTypeLocations**](DefaultApi.md#ProductsTypeLocations) | **Get** /products/types/{typeId}/locations | 
[**RadarProfiler**](DefaultApi.md#RadarProfiler) | **Get** /radar/profilers/{stationId} | 
[**RadarQueue**](DefaultApi.md#RadarQueue) | **Get** /radar/queues/{host} | 
[**RadarServer**](DefaultApi.md#RadarServer) | **Get** /radar/servers/{id} | 
[**RadarServers**](DefaultApi.md#RadarServers) | **Get** /radar/servers | 
[**RadarStation**](DefaultApi.md#RadarStation) | **Get** /radar/stations/{stationId} | 
[**RadarStationAlarms**](DefaultApi.md#RadarStationAlarms) | **Get** /radar/stations/{stationId}/alarms | 
[**RadarStations**](DefaultApi.md#RadarStations) | **Get** /radar/stations | 
[**SatelliteThumbnails**](DefaultApi.md#SatelliteThumbnails) | **Get** /thumbnails/satellite/{area} | 
[**StationObservationLatest**](DefaultApi.md#StationObservationLatest) | **Get** /stations/{stationId}/observations/latest | 
[**StationObservationList**](DefaultApi.md#StationObservationList) | **Get** /stations/{stationId}/observations | 
[**StationObservationTime**](DefaultApi.md#StationObservationTime) | **Get** /stations/{stationId}/observations/{time} | 
[**Zone**](DefaultApi.md#Zone) | **Get** /zones/{type}/{zoneId} | 
[**ZoneForecast**](DefaultApi.md#ZoneForecast) | **Get** /zones/{type}/{zoneId}/forecast | 
[**ZoneList**](DefaultApi.md#ZoneList) | **Get** /zones | 
[**ZoneListType**](DefaultApi.md#ZoneListType) | **Get** /zones/{type} | 
[**ZoneObs**](DefaultApi.md#ZoneObs) | **Get** /zones/forecast/{zoneId}/observations | 
[**ZoneStations**](DefaultApi.md#ZoneStations) | **Get** /zones/forecast/{zoneId}/stations | 



## AlertsActive

> AlertCollectionGeoJson AlertsActive(ctx).Status(status).MessageType(messageType).Event(event).Code(code).Area(area).Point(point).Region(region).RegionType(regionType).Zone(zone).Urgency(urgency).Severity(severity).Certainty(certainty).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    status := []string{"Status_example"} // []string | Status (actual, exercise, system, test, draft) (optional)
    messageType := []string{"MessageType_example"} // []string | Message type (alert, update, cancel) (optional)
    event := []string{"Inner_example"} // []string | Event name (optional)
    code := []string{"Inner_example"} // []string | Event code (optional)
    area := []openapiclient.AreaCode{openapiclient.AreaCode{MarineAreaCode: penapiclient.MarineAreaCode("AM")}} // []AreaCode | State/territory code or marine area code This parameter is incompatible with the following parameters: point, region, region_type, zone  (optional)
    point := "point_example" // string | Point (latitude,longitude) This parameter is incompatible with the following parameters: area, region, region_type, zone  (optional)
    region := []openapiclient.MarineRegionCode{openapiclient.MarineRegionCode("AL")} // []MarineRegionCode | Marine region code This parameter is incompatible with the following parameters: area, point, region_type, zone  (optional)
    regionType := "regionType_example" // string | Region type (land or marine) This parameter is incompatible with the following parameters: area, point, region, zone  (optional)
    zone := []string{"Inner_example"} // []string | Zone ID (forecast or county) This parameter is incompatible with the following parameters: area, point, region, region_type  (optional)
    urgency := []openapiclient.AlertUrgency{openapiclient.AlertUrgency("Immediate")} // []AlertUrgency | Urgency (immediate, expected, future, past, unknown) (optional)
    severity := []openapiclient.AlertSeverity{openapiclient.AlertSeverity("Extreme")} // []AlertSeverity | Severity (extreme, severe, moderate, minor, unknown) (optional)
    certainty := []openapiclient.AlertCertainty{openapiclient.AlertCertainty("Observed")} // []AlertCertainty | Certainty (observed, likely, possible, unlikely, unknown) (optional)
    limit := int32(56) // int32 | Limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlertsActive(context.Background()).Status(status).MessageType(messageType).Event(event).Code(code).Area(area).Point(point).Region(region).RegionType(regionType).Zone(zone).Urgency(urgency).Severity(severity).Certainty(certainty).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsActive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsActive`: AlertCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsActive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **[]string** | Status (actual, exercise, system, test, draft) | 
 **messageType** | **[]string** | Message type (alert, update, cancel) | 
 **event** | **[]string** | Event name | 
 **code** | **[]string** | Event code | 
 **area** | [**[]AreaCode**](AreaCode.md) | State/territory code or marine area code This parameter is incompatible with the following parameters: point, region, region_type, zone  | 
 **point** | **string** | Point (latitude,longitude) This parameter is incompatible with the following parameters: area, region, region_type, zone  | 
 **region** | [**[]MarineRegionCode**](MarineRegionCode.md) | Marine region code This parameter is incompatible with the following parameters: area, point, region_type, zone  | 
 **regionType** | **string** | Region type (land or marine) This parameter is incompatible with the following parameters: area, point, region, zone  | 
 **zone** | **[]string** | Zone ID (forecast or county) This parameter is incompatible with the following parameters: area, point, region, region_type  | 
 **urgency** | [**[]AlertUrgency**](AlertUrgency.md) | Urgency (immediate, expected, future, past, unknown) | 
 **severity** | [**[]AlertSeverity**](AlertSeverity.md) | Severity (extreme, severe, moderate, minor, unknown) | 
 **certainty** | [**[]AlertCertainty**](AlertCertainty.md) | Certainty (observed, likely, possible, unlikely, unknown) | 
 **limit** | **int32** | Limit | 

### Return type

[**AlertCollectionGeoJson**](AlertCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/atom+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsActiveArea

> AlertCollectionGeoJson AlertsActiveArea(ctx, area).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    area := openapiclient.AreaCode{MarineAreaCode: penapiclient.MarineAreaCode("AM")} // AreaCode | State/area ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlertsActiveArea(context.Background(), area).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsActiveArea``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsActiveArea`: AlertCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsActiveArea`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**area** | [**AreaCode**](.md) | State/area ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsActiveAreaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertCollectionGeoJson**](AlertCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/atom+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsActiveCount

> AlertsActiveCount200Response AlertsActiveCount(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlertsActiveCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsActiveCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsActiveCount`: AlertsActiveCount200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsActiveCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsActiveCountRequest struct via the builder pattern


### Return type

[**AlertsActiveCount200Response**](AlertsActiveCount200Response.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsActiveRegion

> AlertCollectionGeoJson AlertsActiveRegion(ctx, region).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    region := openapiclient.MarineRegionCode("AL") // MarineRegionCode | Marine region ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlertsActiveRegion(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsActiveRegion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsActiveRegion`: AlertCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsActiveRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | [**MarineRegionCode**](.md) | Marine region ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsActiveRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertCollectionGeoJson**](AlertCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/atom+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsActiveZone

> AlertCollectionGeoJson AlertsActiveZone(ctx, zoneId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    zoneId := "zoneId_example" // string | NWS public zone/county identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlertsActiveZone(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsActiveZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsActiveZone`: AlertCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsActiveZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | NWS public zone/county identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsActiveZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertCollectionGeoJson**](AlertCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/atom+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsQuery

> AlertCollectionGeoJson AlertsQuery(ctx).Active(active).Start(start).End(end).Status(status).MessageType(messageType).Event(event).Code(code).Area(area).Point(point).Region(region).RegionType(regionType).Zone(zone).Urgency(urgency).Severity(severity).Certainty(certainty).Limit(limit).Cursor(cursor).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    active := true // bool | List only active alerts (use /alerts/active endpoints instead) (optional)
    start := time.Now() // time.Time | Start time (optional)
    end := time.Now() // time.Time | End time (optional)
    status := []string{"Status_example"} // []string | Status (actual, exercise, system, test, draft) (optional)
    messageType := []string{"MessageType_example"} // []string | Message type (alert, update, cancel) (optional)
    event := []string{"Inner_example"} // []string | Event name (optional)
    code := []string{"Inner_example"} // []string | Event code (optional)
    area := []openapiclient.AreaCode{openapiclient.AreaCode{MarineAreaCode: penapiclient.MarineAreaCode("AM")}} // []AreaCode | State/territory code or marine area code This parameter is incompatible with the following parameters: point, region, region_type, zone  (optional)
    point := "point_example" // string | Point (latitude,longitude) This parameter is incompatible with the following parameters: area, region, region_type, zone  (optional)
    region := []openapiclient.MarineRegionCode{openapiclient.MarineRegionCode("AL")} // []MarineRegionCode | Marine region code This parameter is incompatible with the following parameters: area, point, region_type, zone  (optional)
    regionType := "regionType_example" // string | Region type (land or marine) This parameter is incompatible with the following parameters: area, point, region, zone  (optional)
    zone := []string{"Inner_example"} // []string | Zone ID (forecast or county) This parameter is incompatible with the following parameters: area, point, region, region_type  (optional)
    urgency := []openapiclient.AlertUrgency{openapiclient.AlertUrgency("Immediate")} // []AlertUrgency | Urgency (immediate, expected, future, past, unknown) (optional)
    severity := []openapiclient.AlertSeverity{openapiclient.AlertSeverity("Extreme")} // []AlertSeverity | Severity (extreme, severe, moderate, minor, unknown) (optional)
    certainty := []openapiclient.AlertCertainty{openapiclient.AlertCertainty("Observed")} // []AlertCertainty | Certainty (observed, likely, possible, unlikely, unknown) (optional)
    limit := int32(56) // int32 | Limit (optional)
    cursor := "cursor_example" // string | Pagination cursor (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlertsQuery(context.Background()).Active(active).Start(start).End(end).Status(status).MessageType(messageType).Event(event).Code(code).Area(area).Point(point).Region(region).RegionType(regionType).Zone(zone).Urgency(urgency).Severity(severity).Certainty(certainty).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsQuery`: AlertCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAlertsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** | List only active alerts (use /alerts/active endpoints instead) | 
 **start** | **time.Time** | Start time | 
 **end** | **time.Time** | End time | 
 **status** | **[]string** | Status (actual, exercise, system, test, draft) | 
 **messageType** | **[]string** | Message type (alert, update, cancel) | 
 **event** | **[]string** | Event name | 
 **code** | **[]string** | Event code | 
 **area** | [**[]AreaCode**](AreaCode.md) | State/territory code or marine area code This parameter is incompatible with the following parameters: point, region, region_type, zone  | 
 **point** | **string** | Point (latitude,longitude) This parameter is incompatible with the following parameters: area, region, region_type, zone  | 
 **region** | [**[]MarineRegionCode**](MarineRegionCode.md) | Marine region code This parameter is incompatible with the following parameters: area, point, region_type, zone  | 
 **regionType** | **string** | Region type (land or marine) This parameter is incompatible with the following parameters: area, point, region, zone  | 
 **zone** | **[]string** | Zone ID (forecast or county) This parameter is incompatible with the following parameters: area, point, region, region_type  | 
 **urgency** | [**[]AlertUrgency**](AlertUrgency.md) | Urgency (immediate, expected, future, past, unknown) | 
 **severity** | [**[]AlertSeverity**](AlertSeverity.md) | Severity (extreme, severe, moderate, minor, unknown) | 
 **certainty** | [**[]AlertCertainty**](AlertCertainty.md) | Certainty (observed, likely, possible, unlikely, unknown) | 
 **limit** | **int32** | Limit | 
 **cursor** | **string** | Pagination cursor | 

### Return type

[**AlertCollectionGeoJson**](AlertCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/atom+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsSingle

> AlertGeoJson AlertsSingle(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Alert identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlertsSingle(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsSingle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsSingle`: AlertGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsSingle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Alert identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertGeoJson**](AlertGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/cap+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AlertsTypes

> AlertsTypes200Response AlertsTypes(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AlertsTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AlertsTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AlertsTypes`: AlertsTypes200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AlertsTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAlertsTypesRequest struct via the builder pattern


### Return type

[**AlertsTypes200Response**](AlertsTypes200Response.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Glossary

> Glossary200Response Glossary(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Glossary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Glossary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Glossary`: Glossary200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Glossary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGlossaryRequest struct via the builder pattern


### Return type

[**Glossary200Response**](Glossary200Response.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Gridpoint

> GridpointGeoJson Gridpoint(ctx, wfo, x, y).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    wfo := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | Forecast office ID
    x := int32(56) // int32 | Forecast grid X coordinate
    y := int32(56) // int32 | Forecast grid Y coordinate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Gridpoint(context.Background(), wfo, x, y).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Gridpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Gridpoint`: GridpointGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Gridpoint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfo** | [**NWSForecastOfficeId**](.md) | Forecast office ID | 
**x** | **int32** | Forecast grid X coordinate | 
**y** | **int32** | Forecast grid Y coordinate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGridpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GridpointGeoJson**](GridpointGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridpointForecast

> GridpointForecastGeoJson GridpointForecast(ctx, wfo, x, y).FeatureFlags(featureFlags).Units(units).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    wfo := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | Forecast office ID
    x := int32(56) // int32 | Forecast grid X coordinate
    y := int32(56) // int32 | Forecast grid Y coordinate
    featureFlags := []string{"FeatureFlags_example"} // []string | Enable future and experimental features (see documentation for more info): * forecast_temperature_qv: Represent temperature as QuantitativeValue * forecast_wind_speed_qv: Represent wind speed as QuantitativeValue  (optional)
    units := openapiclient.GridpointForecastUnits("us") // GridpointForecastUnits | Use US customary or SI (metric) units in textual output (optional) (default to "us")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GridpointForecast(context.Background(), wfo, x, y).FeatureFlags(featureFlags).Units(units).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GridpointForecast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GridpointForecast`: GridpointForecastGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GridpointForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfo** | [**NWSForecastOfficeId**](.md) | Forecast office ID | 
**x** | **int32** | Forecast grid X coordinate | 
**y** | **int32** | Forecast grid Y coordinate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGridpointForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **featureFlags** | **[]string** | Enable future and experimental features (see documentation for more info): * forecast_temperature_qv: Represent temperature as QuantitativeValue * forecast_wind_speed_qv: Represent wind speed as QuantitativeValue  | 
 **units** | [**GridpointForecastUnits**](GridpointForecastUnits.md) | Use US customary or SI (metric) units in textual output | [default to &quot;us&quot;]

### Return type

[**GridpointForecastGeoJson**](GridpointForecastGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/vnd.noaa.dwml+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridpointForecastHourly

> GridpointForecastGeoJson GridpointForecastHourly(ctx, wfo, x, y).FeatureFlags(featureFlags).Units(units).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    wfo := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | Forecast office ID
    x := int32(56) // int32 | Forecast grid X coordinate
    y := int32(56) // int32 | Forecast grid Y coordinate
    featureFlags := []string{"FeatureFlags_example"} // []string | Enable future and experimental features (see documentation for more info): * forecast_temperature_qv: Represent temperature as QuantitativeValue * forecast_wind_speed_qv: Represent wind speed as QuantitativeValue  (optional)
    units := openapiclient.GridpointForecastUnits("us") // GridpointForecastUnits | Use US customary or SI (metric) units in textual output (optional) (default to "us")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GridpointForecastHourly(context.Background(), wfo, x, y).FeatureFlags(featureFlags).Units(units).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GridpointForecastHourly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GridpointForecastHourly`: GridpointForecastGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GridpointForecastHourly`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfo** | [**NWSForecastOfficeId**](.md) | Forecast office ID | 
**x** | **int32** | Forecast grid X coordinate | 
**y** | **int32** | Forecast grid Y coordinate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGridpointForecastHourlyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **featureFlags** | **[]string** | Enable future and experimental features (see documentation for more info): * forecast_temperature_qv: Represent temperature as QuantitativeValue * forecast_wind_speed_qv: Represent wind speed as QuantitativeValue  | 
 **units** | [**GridpointForecastUnits**](GridpointForecastUnits.md) | Use US customary or SI (metric) units in textual output | [default to &quot;us&quot;]

### Return type

[**GridpointForecastGeoJson**](GridpointForecastGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/vnd.noaa.dwml+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridpointStations

> ObservationStationCollectionGeoJson GridpointStations(ctx, wfo, x, y).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    wfo := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | Forecast office ID
    x := int32(56) // int32 | Forecast grid X coordinate
    y := int32(56) // int32 | Forecast grid Y coordinate

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GridpointStations(context.Background(), wfo, x, y).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GridpointStations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GridpointStations`: ObservationStationCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GridpointStations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**wfo** | [**NWSForecastOfficeId**](.md) | Forecast office ID | 
**x** | **int32** | Forecast grid X coordinate | 
**y** | **int32** | Forecast grid Y coordinate | 

### Other Parameters

Other parameters are passed through a pointer to a apiGridpointStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ObservationStationCollectionGeoJson**](ObservationStationCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Icons

> *os.File Icons(ctx, set, timeOfDay, first).Size(size).Fontsize(fontsize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    set := "set_example" // string | .
    timeOfDay := "timeOfDay_example" // string | .
    first := "first_example" // string | .
    size := *openapiclient.NewIconsSizeParameter() // IconsSizeParameter | Font size (optional)
    fontsize := int32(56) // int32 | Font size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Icons(context.Background(), set, timeOfDay, first).Size(size).Fontsize(fontsize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Icons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Icons`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Icons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**set** | **string** | . | 
**timeOfDay** | **string** | . | 
**first** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiIconsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **size** | [**IconsSizeParameter**](IconsSizeParameter.md) | Font size | 
 **fontsize** | **int32** | Font size | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IconsDualCondition

> *os.File IconsDualCondition(ctx, set, timeOfDay, first, second).Size(size).Fontsize(fontsize).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    set := "set_example" // string | .
    timeOfDay := "timeOfDay_example" // string | .
    first := "first_example" // string | .
    second := "second_example" // string | .
    size := *openapiclient.NewIconsSizeParameter() // IconsSizeParameter | Font size (optional)
    fontsize := int32(56) // int32 | Font size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.IconsDualCondition(context.Background(), set, timeOfDay, first, second).Size(size).Fontsize(fontsize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IconsDualCondition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IconsDualCondition`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IconsDualCondition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**set** | **string** | . | 
**timeOfDay** | **string** | . | 
**first** | **string** | . | 
**second** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiIconsDualConditionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **size** | [**IconsSizeParameter**](IconsSizeParameter.md) | Font size | 
 **fontsize** | **int32** | Font size | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IconsSummary

> IconsSummary200Response IconsSummary(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.IconsSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IconsSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IconsSummary`: IconsSummary200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IconsSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIconsSummaryRequest struct via the builder pattern


### Return type

[**IconsSummary200Response**](IconsSummary200Response.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationProducts

> TextProductTypeCollection LocationProducts(ctx, locationId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    locationId := "locationId_example" // string | .

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.LocationProducts(context.Background(), locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LocationProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LocationProducts`: TextProductTypeCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LocationProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiLocationProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TextProductTypeCollection**](TextProductTypeCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObsStation

> ObservationStationGeoJson ObsStation(ctx, stationId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stationId := "stationId_example" // string | Observation station ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ObsStation(context.Background(), stationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ObsStation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ObsStation`: ObservationStationGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ObsStation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Observation station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiObsStationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObservationStationGeoJson**](ObservationStationGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObsStations

> ObservationStationCollectionGeoJson ObsStations(ctx).Id(id).State(state).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := []string{"Inner_example"} // []string | Filter by observation station ID (optional)
    state := []openapiclient.AreaCode{openapiclient.AreaCode{MarineAreaCode: penapiclient.MarineAreaCode("AM")}} // []AreaCode | Filter by state/marine area code (optional)
    limit := int32(56) // int32 | Limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ObsStations(context.Background()).Id(id).State(state).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ObsStations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ObsStations`: ObservationStationCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ObsStations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObsStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Filter by observation station ID | 
 **state** | [**[]AreaCode**](AreaCode.md) | Filter by state/marine area code | 
 **limit** | **int32** | Limit | 

### Return type

[**ObservationStationCollectionGeoJson**](ObservationStationCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Office

> Office Office(ctx, officeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    officeId := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | NWS forecast office ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Office(context.Background(), officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Office``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Office`: Office
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Office`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | [**NWSForecastOfficeId**](.md) | NWS forecast office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Office**](Office.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficeHeadline

> OfficeHeadline OfficeHeadline(ctx, officeId, headlineId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    officeId := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | NWS forecast office ID
    headlineId := "headlineId_example" // string | Headline record ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OfficeHeadline(context.Background(), officeId, headlineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OfficeHeadline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfficeHeadline`: OfficeHeadline
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OfficeHeadline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | [**NWSForecastOfficeId**](.md) | NWS forecast office ID | 
**headlineId** | **string** | Headline record ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficeHeadlineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OfficeHeadline**](OfficeHeadline.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OfficeHeadlines

> OfficeHeadlineCollection OfficeHeadlines(ctx, officeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    officeId := openapiclient.NWSForecastOfficeId("AKQ") // NWSForecastOfficeId | NWS forecast office ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.OfficeHeadlines(context.Background(), officeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OfficeHeadlines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OfficeHeadlines`: OfficeHeadlineCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OfficeHeadlines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | [**NWSForecastOfficeId**](.md) | NWS forecast office ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOfficeHeadlinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OfficeHeadlineCollection**](OfficeHeadlineCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Point

> PointGeoJson Point(ctx, point).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    point := "point_example" // string | Point (latitude, longitude)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Point(context.Background(), point).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Point``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Point`: PointGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Point`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**point** | **string** | Point (latitude, longitude) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PointGeoJson**](PointGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PointStations

> ProblemDetail PointStations(ctx, point).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    point := "point_example" // string | Point (latitude, longitude)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PointStations(context.Background(), point).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PointStations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PointStations`: ProblemDetail
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PointStations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**point** | **string** | Point (latitude, longitude) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPointStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProblemDetail**](ProblemDetail.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Product

> TextProduct Product(ctx, productId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    productId := "productId_example" // string | .

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Product(context.Background(), productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Product``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Product`: TextProduct
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Product`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TextProduct**](TextProduct.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductLocations

> TextProductLocationCollection ProductLocations(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ProductLocations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ProductLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductLocations`: TextProductLocationCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ProductLocations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProductLocationsRequest struct via the builder pattern


### Return type

[**TextProductLocationCollection**](TextProductLocationCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductTypes

> TextProductTypeCollection ProductTypes(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ProductTypes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ProductTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductTypes`: TextProductTypeCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ProductTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProductTypesRequest struct via the builder pattern


### Return type

[**TextProductTypeCollection**](TextProductTypeCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsQuery

> TextProductCollection ProductsQuery(ctx).Location(location).Start(start).End(end).Office(office).Wmoid(wmoid).Type_(type_).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    location := []string{"Inner_example"} // []string | Location id (optional)
    start := time.Now() // time.Time | Start time (optional)
    end := time.Now() // time.Time | End time (optional)
    office := []string{"Inner_example"} // []string | Issuing office (optional)
    wmoid := []string{"Inner_example"} // []string | WMO id code (optional)
    type_ := []string{"Inner_example"} // []string | Product code (optional)
    limit := int32(56) // int32 | Limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ProductsQuery(context.Background()).Location(location).Start(start).End(end).Office(office).Wmoid(wmoid).Type_(type_).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ProductsQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductsQuery`: TextProductCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ProductsQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProductsQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **location** | **[]string** | Location id | 
 **start** | **time.Time** | Start time | 
 **end** | **time.Time** | End time | 
 **office** | **[]string** | Issuing office | 
 **wmoid** | **[]string** | WMO id code | 
 **type_** | **[]string** | Product code | 
 **limit** | **int32** | Limit | 

### Return type

[**TextProductCollection**](TextProductCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsType

> TextProductCollection ProductsType(ctx, typeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    typeId := "typeId_example" // string | .

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ProductsType(context.Background(), typeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ProductsType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductsType`: TextProductCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ProductsType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TextProductCollection**](TextProductCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsTypeLocation

> TextProductCollection ProductsTypeLocation(ctx, typeId, locationId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    typeId := "typeId_example" // string | .
    locationId := "locationId_example" // string | .

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ProductsTypeLocation(context.Background(), typeId, locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ProductsTypeLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductsTypeLocation`: TextProductCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ProductsTypeLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** | . | 
**locationId** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsTypeLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TextProductCollection**](TextProductCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProductsTypeLocations

> TextProductLocationCollection ProductsTypeLocations(ctx, typeId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    typeId := "typeId_example" // string | .

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ProductsTypeLocations(context.Background(), typeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ProductsTypeLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProductsTypeLocations`: TextProductLocationCollection
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ProductsTypeLocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typeId** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiProductsTypeLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TextProductLocationCollection**](TextProductLocationCollection.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarProfiler

> RadarProfiler(ctx, stationId).Time(time).Interval(interval).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stationId := "stationId_example" // string | Profiler station ID
    time := openapiclient.ISO8601Interval{String: new(string)} // ISO8601Interval | Time interval (optional)
    interval := "interval_example" // string | Averaging interval (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RadarProfiler(context.Background(), stationId).Time(time).Interval(interval).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RadarProfiler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Profiler station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadarProfilerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **time** | [**ISO8601Interval**](ISO8601Interval.md) | Time interval | 
 **interval** | **string** | Averaging interval | 

### Return type

 (empty response body)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarQueue

> RadarQueue(ctx, host).Limit(limit).Arrived(arrived).Created(created).Published(published).Station(station).Type_(type_).Feed(feed).Resolution(resolution).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    host := "host_example" // string | LDM host
    limit := int32(56) // int32 | Record limit (optional)
    arrived := openapiclient.ISO8601Interval{String: new(string)} // ISO8601Interval | Range for arrival time (optional)
    created := openapiclient.ISO8601Interval{String: new(string)} // ISO8601Interval | Range for creation time (optional)
    published := openapiclient.ISO8601Interval{String: new(string)} // ISO8601Interval | Range for publish time (optional)
    station := "station_example" // string | Station identifier (optional)
    type_ := "type__example" // string | Record type (optional)
    feed := "feed_example" // string | Originating product feed (optional)
    resolution := int32(56) // int32 | Resolution version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RadarQueue(context.Background(), host).Limit(limit).Arrived(arrived).Created(created).Published(published).Station(station).Type_(type_).Feed(feed).Resolution(resolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RadarQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**host** | **string** | LDM host | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadarQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Record limit | 
 **arrived** | [**ISO8601Interval**](ISO8601Interval.md) | Range for arrival time | 
 **created** | [**ISO8601Interval**](ISO8601Interval.md) | Range for creation time | 
 **published** | [**ISO8601Interval**](ISO8601Interval.md) | Range for publish time | 
 **station** | **string** | Station identifier | 
 **type_** | **string** | Record type | 
 **feed** | **string** | Originating product feed | 
 **resolution** | **int32** | Resolution version | 

### Return type

 (empty response body)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarServer

> RadarServer(ctx, id).ReportingHost(reportingHost).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Server ID
    reportingHost := "reportingHost_example" // string | Show records from specific reporting host (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RadarServer(context.Background(), id).ReportingHost(reportingHost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RadarServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadarServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportingHost** | **string** | Show records from specific reporting host | 

### Return type

 (empty response body)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarServers

> RadarServers(ctx).ReportingHost(reportingHost).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    reportingHost := "reportingHost_example" // string | Show records from specific reporting host (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RadarServers(context.Background()).ReportingHost(reportingHost).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RadarServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRadarServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reportingHost** | **string** | Show records from specific reporting host | 

### Return type

 (empty response body)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarStation

> RadarStation(ctx, stationId).ReportingHost(reportingHost).Host(host).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stationId := "stationId_example" // string | Radar station ID
    reportingHost := "reportingHost_example" // string | Show RDA and latency info from specific reporting host (optional)
    host := "host_example" // string | Show latency info from specific LDM host (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RadarStation(context.Background(), stationId).ReportingHost(reportingHost).Host(host).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RadarStation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Radar station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadarStationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reportingHost** | **string** | Show RDA and latency info from specific reporting host | 
 **host** | **string** | Show latency info from specific LDM host | 

### Return type

 (empty response body)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarStationAlarms

> RadarStationAlarms(ctx, stationId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stationId := "stationId_example" // string | Radar station ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RadarStationAlarms(context.Background(), stationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RadarStationAlarms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Radar station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRadarStationAlarmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadarStations

> RadarStations(ctx).StationType(stationType).ReportingHost(reportingHost).Host(host).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stationType := []string{"Inner_example"} // []string | Limit results to a specific station type or types (optional)
    reportingHost := "reportingHost_example" // string | Show RDA and latency info from specific reporting host (optional)
    host := "host_example" // string | Show latency info from specific LDM host (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RadarStations(context.Background()).StationType(stationType).ReportingHost(reportingHost).Host(host).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RadarStations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRadarStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stationType** | **[]string** | Limit results to a specific station type or types | 
 **reportingHost** | **string** | Show RDA and latency info from specific reporting host | 
 **host** | **string** | Show latency info from specific LDM host | 

### Return type

 (empty response body)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SatelliteThumbnails

> *os.File SatelliteThumbnails(ctx, area).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    area := "area_example" // string | .

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SatelliteThumbnails(context.Background(), area).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SatelliteThumbnails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SatelliteThumbnails`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SatelliteThumbnails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**area** | **string** | . | 

### Other Parameters

Other parameters are passed through a pointer to a apiSatelliteThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/jpeg, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StationObservationLatest

> ObservationGeoJson StationObservationLatest(ctx, stationId).RequireQc(requireQc).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stationId := "stationId_example" // string | Observation station ID
    requireQc := true // bool | Require QC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.StationObservationLatest(context.Background(), stationId).RequireQc(requireQc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StationObservationLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StationObservationLatest`: ObservationGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StationObservationLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Observation station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStationObservationLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requireQc** | **bool** | Require QC | 

### Return type

[**ObservationGeoJson**](ObservationGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/vnd.noaa.obs+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StationObservationList

> ObservationCollectionGeoJson StationObservationList(ctx, stationId).Start(start).End(end).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    stationId := "stationId_example" // string | Observation station ID
    start := time.Now() // time.Time | Start time (optional)
    end := time.Now() // time.Time | End time (optional)
    limit := int32(56) // int32 | Limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.StationObservationList(context.Background(), stationId).Start(start).End(end).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StationObservationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StationObservationList`: ObservationCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StationObservationList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Observation station ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStationObservationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **time.Time** | Start time | 
 **end** | **time.Time** | End time | 
 **limit** | **int32** | Limit | 

### Return type

[**ObservationCollectionGeoJson**](ObservationCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StationObservationTime

> ObservationGeoJson StationObservationTime(ctx, stationId, time).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    stationId := "stationId_example" // string | Observation station ID
    time := time.Now() // time.Time | Timestamp of requested observation

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.StationObservationTime(context.Background(), stationId, time).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StationObservationTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StationObservationTime`: ObservationGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StationObservationTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stationId** | **string** | Observation station ID | 
**time** | **time.Time** | Timestamp of requested observation | 

### Other Parameters

Other parameters are passed through a pointer to a apiStationObservationTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ObservationGeoJson**](ObservationGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/vnd.noaa.obs+xml, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Zone

> ZoneGeoJson Zone(ctx, type_, zoneId).Effective(effective).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    type_ := "type__example" // string | Zone type
    zoneId := "zoneId_example" // string | NWS public zone/county identifier
    effective := time.Now() // time.Time | Effective date/time (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Zone(context.Background(), type_, zoneId).Effective(effective).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Zone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Zone`: ZoneGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Zone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Zone type | 
**zoneId** | **string** | NWS public zone/county identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **effective** | **time.Time** | Effective date/time | 

### Return type

[**ZoneGeoJson**](ZoneGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneForecast

> ZoneForecastGeoJson ZoneForecast(ctx, type_, zoneId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    type_ := "type__example" // string | Zone type
    zoneId := "zoneId_example" // string | NWS public zone/county identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ZoneForecast(context.Background(), type_, zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ZoneForecast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZoneForecast`: ZoneForecastGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ZoneForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Zone type | 
**zoneId** | **string** | NWS public zone/county identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ZoneForecastGeoJson**](ZoneForecastGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneList

> ZoneCollectionGeoJson ZoneList(ctx).Id(id).Area(area).Region(region).Type_(type_).Point(point).IncludeGeometry(includeGeometry).Limit(limit).Effective(effective).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := []string{"Inner_example"} // []string | Zone ID (forecast or county) (optional)
    area := []openapiclient.AreaCode{openapiclient.AreaCode{MarineAreaCode: penapiclient.MarineAreaCode("AM")}} // []AreaCode | State/marine area code (optional)
    region := []openapiclient.RegionCode{openapiclient.RegionCode{LandRegionCode: penapiclient.LandRegionCode("AR")}} // []RegionCode | Region code (optional)
    type_ := []string{"Type_example"} // []string | Zone type (optional)
    point := "point_example" // string | Point (latitude,longitude) (optional)
    includeGeometry := true // bool | Include geometry in results (true/false) (optional)
    limit := int32(56) // int32 | Limit (optional)
    effective := time.Now() // time.Time | Effective date/time (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ZoneList(context.Background()).Id(id).Area(area).Region(region).Type_(type_).Point(point).IncludeGeometry(includeGeometry).Limit(limit).Effective(effective).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ZoneList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZoneList`: ZoneCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ZoneList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiZoneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Zone ID (forecast or county) | 
 **area** | [**[]AreaCode**](AreaCode.md) | State/marine area code | 
 **region** | [**[]RegionCode**](RegionCode.md) | Region code | 
 **type_** | **[]string** | Zone type | 
 **point** | **string** | Point (latitude,longitude) | 
 **includeGeometry** | **bool** | Include geometry in results (true/false) | 
 **limit** | **int32** | Limit | 
 **effective** | **time.Time** | Effective date/time | 

### Return type

[**ZoneCollectionGeoJson**](ZoneCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneListType

> ZoneCollectionGeoJson ZoneListType(ctx, type_).Id(id).Area(area).Region(region).Type_2(type_2).Point(point).IncludeGeometry(includeGeometry).Limit(limit).Effective(effective).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    type_ := "type__example" // string | Zone type
    id := []string{"Inner_example"} // []string | Zone ID (forecast or county) (optional)
    area := []openapiclient.AreaCode{openapiclient.AreaCode{MarineAreaCode: penapiclient.MarineAreaCode("AM")}} // []AreaCode | State/marine area code (optional)
    region := []openapiclient.RegionCode{openapiclient.RegionCode{LandRegionCode: penapiclient.LandRegionCode("AR")}} // []RegionCode | Region code (optional)
    type_2 := []string{"Type_example"} // []string | Zone type (optional)
    point := "point_example" // string | Point (latitude,longitude) (optional)
    includeGeometry := true // bool | Include geometry in results (true/false) (optional)
    limit := int32(56) // int32 | Limit (optional)
    effective := time.Now() // time.Time | Effective date/time (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ZoneListType(context.Background(), type_).Id(id).Area(area).Region(region).Type_2(type_2).Point(point).IncludeGeometry(includeGeometry).Limit(limit).Effective(effective).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ZoneListType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZoneListType`: ZoneCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ZoneListType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Zone type | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneListTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **[]string** | Zone ID (forecast or county) | 
 **area** | [**[]AreaCode**](AreaCode.md) | State/marine area code | 
 **region** | [**[]RegionCode**](RegionCode.md) | Region code | 
 **type_2** | **[]string** | Zone type | 
 **point** | **string** | Point (latitude,longitude) | 
 **includeGeometry** | **bool** | Include geometry in results (true/false) | 
 **limit** | **int32** | Limit | 
 **effective** | **time.Time** | Effective date/time | 

### Return type

[**ZoneCollectionGeoJson**](ZoneCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneObs

> ObservationCollectionGeoJson ZoneObs(ctx, zoneId).Start(start).End(end).Limit(limit).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    zoneId := "zoneId_example" // string | NWS public zone/county identifier
    start := time.Now() // time.Time | Start date/time (optional)
    end := time.Now() // time.Time | End date/time (optional)
    limit := int32(56) // int32 | Limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ZoneObs(context.Background(), zoneId).Start(start).End(end).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ZoneObs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZoneObs`: ObservationCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ZoneObs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | NWS public zone/county identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneObsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **time.Time** | Start date/time | 
 **end** | **time.Time** | End date/time | 
 **limit** | **int32** | Limit | 

### Return type

[**ObservationCollectionGeoJson**](ObservationCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ZoneStations

> ObservationStationCollectionGeoJson ZoneStations(ctx, zoneId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    zoneId := "zoneId_example" // string | NWS public zone/county identifier

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ZoneStations(context.Background(), zoneId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ZoneStations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ZoneStations`: ObservationStationCollectionGeoJson
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ZoneStations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | NWS public zone/county identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiZoneStationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObservationStationCollectionGeoJson**](ObservationStationCollectionGeoJson.md)

### Authorization

[userAgent](../README.md#userAgent)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/geo+json, application/ld+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

