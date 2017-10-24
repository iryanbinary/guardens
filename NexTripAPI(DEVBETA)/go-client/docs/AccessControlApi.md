# \AccessControlApi

All URIs are relative to *https://svc.metrotransit.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuthCode**](AccessControlApi.md#GetAuthCode) | **Get** /NexTrip/oauth20/authorize | 
[**GetTokenRequest**](AccessControlApi.md#GetTokenRequest) | **Get** /NexTrip/oauth20/token | 
[**PostTokenRequest**](AccessControlApi.md#PostTokenRequest) | **Post** /NexTrip/oauth20/token | 


# **GetAuthCode**
> Success GetAuthCode($grantType, $clientId, $redirectUri)



Request a temporary code for the desired API Access Token Scope(s)


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string**| value &#x3D; authorization_code | 
 **clientId** | **string**| a valid OAuth2 client id registered to the app authorized to use the API is required (this changes with each app store version iteration) | 
 **redirectUri** | **string**| App Callback URI | 

### Return type

[**Success**](success.md)

### Authorization

[AccessCode](../README.md#AccessCode), [MobileApp_Implicit](../README.md#MobileApp_Implicit), [admin_AccessCode](../README.md#admin_AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokenRequest**
> OAuthToken GetTokenRequest($grantType, $clientId, $clientSecret)



Applications request an implicit token with a client id and secret prior to user authentication


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string**| value &#x3D; implicit | 
 **clientId** | **string**| a valid  OAuth2 client id registered to the app authorized to use the API is required (this changes with each app store version iteration) | 
 **clientSecret** | **string**| the client secret associated to the app client id (this changes with each app store version iteration) | 

### Return type

[**OAuthToken**](OAuth_Token.md)

### Authorization

[AccessCode](../README.md#AccessCode), [MobileApp_Implicit](../README.md#MobileApp_Implicit), [admin_AccessCode](../README.md#admin_AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostTokenRequest**
> OAuthToken PostTokenRequest()



Authorization Code grant types require a POST to the token endpoint after a GET for Authorization code.


### Parameters
This endpoint does not need any parameter.

### Return type

[**OAuthToken**](OAuth_Token.md)

### Authorization

[AccessCode](../README.md#AccessCode), [MobileApp_Implicit](../README.md#MobileApp_Implicit), [admin_AccessCode](../README.md#admin_AccessCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

