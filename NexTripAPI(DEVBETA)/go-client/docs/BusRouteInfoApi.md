# \BusRouteInfoApi

All URIs are relative to *https://svc.metrotransit.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Updateroute**](BusRouteInfoApi.md#Updateroute) | **Post** /NexTrip/update/{ROUTE} | 


# **Updateroute**
> RouteData Updateroute($rOUTE)



Update a bus route


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rOUTE** | **int32**| Sepcify the Route ID as an integer. | 

### Return type

[**RouteData**](RouteData.md)

### Authorization

[admin_AccessCode](../README.md#admin_AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

