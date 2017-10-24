# SwaggerClient::BusRouteInfoApi

All URIs are relative to *https://svc.metrotransit.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**updateroute**](BusRouteInfoApi.md#updateroute) | **POST** /NexTrip/update/{ROUTE} | 


# **updateroute**
> RouteData updateroute(route)



Update a bus route

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: admin_AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::BusRouteInfoApi.new

route = 56 # Integer | Sepcify the Route ID as an integer.


begin
  result = api_instance.updateroute(route)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling BusRouteInfoApi->updateroute: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **route** | **Integer**| Sepcify the Route ID as an integer. | 

### Return type

[**RouteData**](RouteData.md)

### Authorization

[admin_AccessCode](../README.md#admin_AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



