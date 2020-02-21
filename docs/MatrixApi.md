# \MatrixApi

All URIs are relative to *https://eu1.locationiq.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Matrix**](MatrixApi.md#Matrix) | **Get** /matrix/driving/{coordinates} | Matrix Service



## Matrix

> DirectionsMatrix Matrix(ctx, coordinates, optional)

Matrix Service

Computes duration of the fastest route between all pairs of supplied coordinates. Returns the durations or distances or both between the coordinate pairs. Note that the distances are not the shortest distance between two coordinates, but rather the distances of the fastest routes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**coordinates** | **string**| String of format {longitude},{latitude};{longitude},{latitude}[;{longitude},{latitude} ...] or polyline({polyline}) or polyline6({polyline6}). polyline follows Google&#39;s polyline format with precision 5 | 
 **optional** | ***MatrixOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MatrixOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bearings** | **optional.String**| Limits the search to segments with given bearing in degrees towards true north in clockwise direction. List of positive integer pairs separated by semi-colon and bearings array should be equal to length of coordinate array. Input Value :- {bearing};{bearing}[;{bearing} ...] Bearing follows the following format : bearing {value},{range} integer 0 .. 360,integer 0 .. 180 | 
 **radiuses** | **optional.String**| Limits the search to given radius in meters Radiuses array length should be same as coordinates array, eaach value separated by semi-colon. Input Value - {radius};{radius}[;{radius} ...] Radius has following format :- double &gt;&#x3D; 0 or unlimited (default) | 
 **generateHints** | **optional.String**| Adds a Hint to the response which can be used in subsequent requests, see hints parameter. Input Value - true (default), false Format - Base64 String | 
 **approaches** | **optional.String**| Keep waypoints on curb side. Input Value - {approach};{approach}[;{approach} ...] Format - curb or unrestricted (default) | 
 **exclude** | **optional.String**| Additive list of classes to avoid, order does not matter. input Value - {class}[,{class}] Format - A class name determined by the profile or none. | 
 **annotations** | **optional.String**| Returns additional metadata for each coordinate along the route geometry.  [ true, false (default), nodes, distance, duration, datasources, weight, speed ] | 
 **sources** | **optional.Int32**| Use location with given index as source. [ {index};{index}[;{index} ...] or all (default) ] &#x3D;&gt; index  0 &lt;&#x3D; integer &lt; #locations | 
 **destinations** | **optional.Int32**| Use location with given index as destination. [ {index};{index}[;{index} ...] or all (default) ] | 
 **fallbackSpeed** | **optional.Float32**| If no route found between a source/destination pair, calculate the as-the-crow-flies distance,  then use this speed to estimate duration. double &gt; 0 | 
 **fallbackCoordinate** | **optional.String**| When using a fallback_speed, use the user-supplied coordinate (input), or the snapped location (snapped) for calculating distances. [ input (default), or snapped ] | [default to &quot;input&quot;]

### Return type

[**DirectionsMatrix**](directions-matrix.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

