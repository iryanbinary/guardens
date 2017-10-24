# SwaggerClient::AccessControlApi

All URIs are relative to *https://svc.metrotransit.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_auth_code**](AccessControlApi.md#get_auth_code) | **GET** /NexTrip/oauth20/authorize | 
[**get_token_request**](AccessControlApi.md#get_token_request) | **GET** /NexTrip/oauth20/token | 
[**post_token_request**](AccessControlApi.md#post_token_request) | **POST** /NexTrip/oauth20/token | 


# **get_auth_code**
> Success get_auth_code(grant_type, client_id, redirect_uri)



Request a temporary code for the desired API Access Token Scope(s)

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'

  # Configure OAuth2 access token for authorization: MobileApp_Implicit
  config.access_token = 'YOUR ACCESS TOKEN'

  # Configure OAuth2 access token for authorization: admin_AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::AccessControlApi.new

grant_type = "grant_type_example" # String | value = authorization_code

client_id = "client_id_example" # String | a valid OAuth2 client id registered to the app authorized to use the API is required (this changes with each app store version iteration)

redirect_uri = "redirect_uri_example" # String | App Callback URI


begin
  result = api_instance.get_auth_code(grant_type, client_id, redirect_uri)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling AccessControlApi->get_auth_code: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grant_type** | **String**| value &#x3D; authorization_code | 
 **client_id** | **String**| a valid OAuth2 client id registered to the app authorized to use the API is required (this changes with each app store version iteration) | 
 **redirect_uri** | **String**| App Callback URI | 

### Return type

[**Success**](Success.md)

### Authorization

[AccessCode](../README.md#AccessCode), [MobileApp_Implicit](../README.md#MobileApp_Implicit), [admin_AccessCode](../README.md#admin_AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



# **get_token_request**
> OAuthToken get_token_request(grant_type, client_id, client_secret)



Applications request an implicit token with a client id and secret prior to user authentication

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'

  # Configure OAuth2 access token for authorization: MobileApp_Implicit
  config.access_token = 'YOUR ACCESS TOKEN'

  # Configure OAuth2 access token for authorization: admin_AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::AccessControlApi.new

grant_type = "grant_type_example" # String | value = implicit

client_id = "client_id_example" # String | a valid  OAuth2 client id registered to the app authorized to use the API is required (this changes with each app store version iteration)

client_secret = "client_secret_example" # String | the client secret associated to the app client id (this changes with each app store version iteration)


begin
  result = api_instance.get_token_request(grant_type, client_id, client_secret)
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling AccessControlApi->get_token_request: #{e}"
end
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grant_type** | **String**| value &#x3D; implicit | 
 **client_id** | **String**| a valid  OAuth2 client id registered to the app authorized to use the API is required (this changes with each app store version iteration) | 
 **client_secret** | **String**| the client secret associated to the app client id (this changes with each app store version iteration) | 

### Return type

[**OAuthToken**](OAuthToken.md)

### Authorization

[AccessCode](../README.md#AccessCode), [MobileApp_Implicit](../README.md#MobileApp_Implicit), [admin_AccessCode](../README.md#admin_AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



# **post_token_request**
> OAuthToken post_token_request



Authorization Code grant types require a POST to the token endpoint after a GET for Authorization code.

### Example
```ruby
# load the gem
require 'swagger_client'
# setup authorization
SwaggerClient.configure do |config|
  # Configure OAuth2 access token for authorization: AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'

  # Configure OAuth2 access token for authorization: MobileApp_Implicit
  config.access_token = 'YOUR ACCESS TOKEN'

  # Configure OAuth2 access token for authorization: admin_AccessCode
  config.access_token = 'YOUR ACCESS TOKEN'
end

api_instance = SwaggerClient::AccessControlApi.new

begin
  result = api_instance.post_token_request
  p result
rescue SwaggerClient::ApiError => e
  puts "Exception when calling AccessControlApi->post_token_request: #{e}"
end
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**OAuthToken**](OAuthToken.md)

### Authorization

[AccessCode](../README.md#AccessCode), [MobileApp_Implicit](../README.md#MobileApp_Implicit), [admin_AccessCode](../README.md#admin_AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml



