# \BusRoutesApi

All URIs are relative to *https://svc.metrotransit.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDepartures**](BusRoutesApi.md#GetDepartures) | **Get** /NexTrip/{STOPID} | 
[**GetDirections**](BusRoutesApi.md#GetDirections) | **Get** /NexTrip/Directions/{ROUTE} | 
[**GetProviders**](BusRoutesApi.md#GetProviders) | **Get** /NexTrip/Providers | 
[**GetRoutes**](BusRoutesApi.md#GetRoutes) | **Get** /NexTrip/Routes | 
[**GetStops**](BusRoutesApi.md#GetStops) | **Get** /NexTrip/Stops/{ROUTE}/{DIRECTION} | 
[**GetTimepointDepartures**](BusRoutesApi.md#GetTimepointDepartures) | **Get** /NexTrip/{ROUTE}/{DIRECTION}/{STOP} | 
[**GetVehicleLocations**](BusRoutesApi.md#GetVehicleLocations) | **Get** /NexTrip/VehicleLocations/{ROUTE} | 


# **GetDepartures**
> Success GetDepartures($sTOPID)



Returns a list of departures scheduled for any given bus stop.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sTOPID** | **string**| Specify the value of the Bus Stop ID as an abbreviated string | 

### Return type

[**Success**](success.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDirections**
> Success GetDirections($rOUTE)



Returns the two directions that are valid for a given route.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rOUTE** | **int32**| Sepcify the Route ID as an integer. | 

### Return type

[**Success**](success.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProviders**
> Success GetProviders()



Returns a list of area Transit providers.  Providers are identified in the list of Routes allowing routes to be selected for a single provider. 


### Parameters
This endpoint does not need any parameter.

### Return type

[**Success**](success.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRoutes**
> RouteData GetRoutes()



Returns a list of Transit routes that are in service on the current day.


### Parameters
This endpoint does not need any parameter.

### Return type

[**RouteData**](RouteData.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStops**
> Success GetStops($rOUTE, $dIRECTION)



Returns a list of Timepoint stops for the given Route/Direction.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rOUTE** | **int32**| Sepcify the Route ID as an integer. | 
 **dIRECTION** | **int32**| Specify the direction as an integer.  1 (South), 2 (East), 3 (West), 4 (North) | 

### Return type

[**Success**](success.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimepointDepartures**
> TimePoints GetTimepointDepartures($rOUTE, $dIRECTION, $sTOP)



Returns the scheduled departures for a selected route, direction and timepoint stop.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rOUTE** | **int32**| Sepcify the Route ID as an integer. | 
 **dIRECTION** | **int32**| Specify the direction as an integer.  1 (South), 2 (East), 3 (West), 4 (North) | 
 **sTOP** | **string**| Specify the value of the Bus Stop ID as an abbreviated string | 

### Return type

[**TimePoints**](TimePoints.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVehicleLocations**
> Success GetVehicleLocations($rOUTE)



This operation returns a list of vehicles currently in service that have recently (within 5 minutes)  reported their locations. A route paramter is used to return results for the given route.  Use \"0\" for the route parameter to return a list of all vehicles in service.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rOUTE** | **int32**| Sepcify the Route ID as an integer. | 

### Return type

[**Success**](success.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

