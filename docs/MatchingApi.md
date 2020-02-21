# \MatchingApi

All URIs are relative to *https://eu1.locationiq.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Matching**](MatchingApi.md#Matching) | **Get** /matching/driving/{coordinates} | Matching Service



## Matching

> DirectionsMatching Matching(ctx, coordinates, optional)

Matching Service

Matching API matches or snaps given GPS points to the road network in the most plausible way.  Please note the request might result multiple sub-traces.  Large jumps in the timestamps (> 60s) or improbable transitions lead to trace splits if a complete matching could not be found. The algorithm might not be able to match all points. Outliers are removed if they can not be matched successfully.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**coordinates** | **string**| String of format {longitude},{latitude};{longitude},{latitude}[;{longitude},{latitude} ...] or polyline({polyline}) or polyline6({polyline6}). polyline follows Google&#39;s polyline format with precision 5 | 
 **optional** | ***MatchingOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MatchingOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateHints** | **optional.String**| Adds a Hint to the response which can be used in subsequent requests, see hints parameter. Input Value - true (default), false Format - Base64 String | 
 **approaches** | **optional.String**| Keep waypoints on curb side. Input Value - {approach};{approach}[;{approach} ...] Format - curb or unrestricted (default) | 
 **exclude** | **optional.String**| Additive list of classes to avoid, order does not matter. input Value - {class}[,{class}] Format - A class name determined by the profile or none. | 
 **bearings** | **optional.String**| Limits the search to segments with given bearing in degrees towards true north in clockwise direction. List of positive integer pairs separated by semi-colon and bearings array should be equal to length of coordinate array. Input Value :- {bearing};{bearing}[;{bearing} ...] Bearing follows the following format : bearing {value},{range} integer 0 .. 360,integer 0 .. 180 | 
 **radiuses** | **optional.String**| Limits the search to given radius in meters Radiuses array length should be same as coordinates array, eaach value separated by semi-colon. Input Value - {radius};{radius}[;{radius} ...] Radius has following format :- double &gt;&#x3D; 0 or unlimited (default) | 
 **steps** | **optional.String**| Returned route steps for each route leg [ true, false (default) ] | 
 **annotations** | **optional.String**| Returns additional metadata for each coordinate along the route geometry.  [ true, false (default), nodes, distance, duration, datasources, weight, speed ] | [default to &quot;false&quot;]
 **geometries** | **optional.String**| Returned route geometry format (influences overview and per step) [ polyline (default), polyline6, geojson ] | [default to &quot;polyline&quot;]
 **overview** | **optional.String**| Add overview geometry either full, simplified according to highest zoom level it could be display on, or not at all. [ simplified (default), full, false ] | [default to &quot;simplified&quot;]
 **timestamps** | **optional.String**| Timestamps for the input locations in seconds since UNIX epoch. Timestamps need to be monotonically increasing. [ {timestamp};{timestamp}[;{timestamp} ...]  integer seconds since UNIX epoch | 
 **gaps** | **optional.String**| Allows the input track splitting based on huge timestamp gaps between points. [ split (default), ignore ] | [default to &quot;split&quot;]
 **tidy** | **optional.String**| Allows the input track modification to obtain better matching quality for noisy tracks. [ true, false (default) ] | [default to &quot;false&quot;]
 **waypoints** | **optional.String**| Treats input coordinates indicated by given indices as waypoints in returned Match object. Default is to treat all input coordinates as waypoints. [ {index};{index};{index}... ] | 

### Return type

[**DirectionsMatching**](directions-matching.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

