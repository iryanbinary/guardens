# SwaggerClient::BusRoutesApi

All URIs are relative to *https://svc.metrotransit.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_departures**](BusRoutesApi.md#get_departures) | **GET** /NexTrip/{STOPID} | 
[**get_directions**](BusRoutesApi.md#get_directions) | **GET** /NexTrip/Directions/{ROUTE} | 
[**get_providers**](BusRoutesApi.md#get_providers) | **GET** /NexTrip/Providers | 
[**get_routes**](BusRoutesApi.md#get_routes) | **GET** /NexTrip/Routes | 
[**get_stops**](BusRoutesApi.md#get_stops) | **GET** /NexTrip/Stops/{ROUTE}/{DIRECTION} | 
[**get_timepoint_departures**](BusRoutesApi.md#get_timepoint_departures) | **GET** /NexTrip/{ROUTE}/{DIRECTION}/{STOP} | 
[**get_vehicle_locations**](BusRoutesApi.md#get_vehicle_locations) | **GET** /NexTrip/VehicleLocations/{ROUTE} | 


# **get_departures**
> Success get_departures(stopid)



Returns a list of departures scheduled for any given bus stop.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::BusRoutesApi.new

stopid = "stopid_example" # String | Specify the value of the Bus Stop ID as an abbreviated string


begin
  result = api_instance.get_departures(stopid)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling BusRoutesApi->get_departures: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stopid** | **String**| Specify the value of the Bus Stop ID as an abbreviated string | 

### Return type

[**Success**](Success.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



# **get_directions**
> Success get_directions(route)



Returns the two directions that are valid for a given route.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::BusRoutesApi.new

route = 56 # Integer | Sepcify the Route ID as an integer.


begin
  result = api_instance.get_directions(route)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling BusRoutesApi->get_directions: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **route** | **Integer**| Sepcify the Route ID as an integer. | 

### Return type

[**Success**](Success.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



# **get_providers**
> Success get_providers



Returns a list of area Transit providers.  Providers are identified in the list of Routes allowing routes to be selected for a single provider. 

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::BusRoutesApi.new

begin
  result = api_instance.get_providers
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling BusRoutesApi->get_providers: #{e}"
end
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**Success**](Success.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



# **get_routes**
> RouteData get_routes



Returns a list of Transit routes that are in service on the current day.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::BusRoutesApi.new

begin
  result = api_instance.get_routes
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling BusRoutesApi->get_routes: #{e}"
end
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**RouteData**](RouteData.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



# **get_stops**
> Success get_stops(route, direction)



Returns a list of Timepoint stops for the given Route/Direction.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::BusRoutesApi.new

route = 56 # Integer | Sepcify the Route ID as an integer.

direction = 56 # Integer | Specify the direction as an integer.  1 (South), 2 (East), 3 (West), 4 (North)


begin
  result = api_instance.get_stops(route, direction)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling BusRoutesApi->get_stops: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **route** | **Integer**| Sepcify the Route ID as an integer. | 
 **direction** | **Integer**| Specify the direction as an integer.  1 (South), 2 (East), 3 (West), 4 (North) | 

### Return type

[**Success**](Success.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



# **get_timepoint_departures**
> TimePoints get_timepoint_departures(route, direction, stop)



Returns the scheduled departures for a selected route, direction and timepoint stop.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::BusRoutesApi.new

route = 56 # Integer | Sepcify the Route ID as an integer.

direction = 56 # Integer | Specify the direction as an integer.  1 (South), 2 (East), 3 (West), 4 (North)

stop = "stop_example" # String | Specify the value of the Bus Stop ID as an abbreviated string


begin
  result = api_instance.get_timepoint_departures(route, direction, stop)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling BusRoutesApi->get_timepoint_departures: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **route** | **Integer**| Sepcify the Route ID as an integer. | 
 **direction** | **Integer**| Specify the direction as an integer.  1 (South), 2 (East), 3 (West), 4 (North) | 
 **stop** | **String**| Specify the value of the Bus Stop ID as an abbreviated string | 

### Return type

[**TimePoints**](TimePoints.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



# **get_vehicle_locations**
> Success get_vehicle_locations(route)



This operation returns a list of vehicles currently in service that have recently (within 5 minutes)  reported their locations. A route paramter is used to return results for the given route.  Use \"0\" for the route parameter to return a list of all vehicles in service.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::BusRoutesApi.new

route = 56 # Integer | Sepcify the Route ID as an integer.


begin
  result = api_instance.get_vehicle_locations(route)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling BusRoutesApi->get_vehicle_locations: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **route** | **Integer**| Sepcify the Route ID as an integer. | 

### Return type

[**Success**](Success.md)

### Authorization

[AccessCode](../README.md#AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



