# LocationIq\ReverseApi

All URIs are relative to *https://eu1.locationiq.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Reverse**](ReverseApi.md#Reverse) | **Get** /reverse.php | Reverse Geocoding



## Reverse

> Location Reverse(ctx, lat, lon, format, normalizecity, optional)

Reverse Geocoding

Reverse geocoding is the process of converting a coordinate or location (latitude, longitude) to a readable address or place name. This permits the identification of nearby street addresses, places, and/or area subdivisions such as a neighborhood, county, state, or country.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lat** | **float32**| Latitude of the location to generate an address for. | 
**lon** | **float32**| Longitude of the location to generate an address for. | 
**format** | **string**| Format to geocode. Only JSON supported for SDKs | 
**normalizecity** | **int32**| Normalizes village to city level data to city | 
 **optional** | ***ReverseOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ReverseOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **addressdetails** | **optional.Int32**| Include a breakdown of the address into elements. Defaults to 1. | [default to 1]
 **acceptLanguage** | **optional.String**| Preferred language order for showing search results, overrides the value specified in the Accept-Language HTTP header. Defaults to en. To use native language for the response when available, use accept-language&#x3D;native | 
 **namedetails** | **optional.Int32**| Include a list of alternative names in the results. These may include language variants, references, operator and brand. | 
 **extratags** | **optional.Int32**| Include additional information in the result if available, e.g. wikipedia link, opening hours. | 
 **statecode** | **optional.Int32**| Adds state or province code when available to the statecode key inside the address element. Currently supported for addresses in the USA, Canada and Australia. Defaults to 0 | 
 **showdistance** | **optional.Int32**| Returns the straight line distance (meters) between the input location and the result&#39;s location. Value is set in the distance key of the response. Defaults to 0 [0,1] | 
 **postaladdress** | **optional.Int32**| Returns address inside the postaladdress key, that is specifically formatted for each country. Currently supported for addresses in Germany. Defaults to 0 [0,1] | 

### Return type

[**Location**](location.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

