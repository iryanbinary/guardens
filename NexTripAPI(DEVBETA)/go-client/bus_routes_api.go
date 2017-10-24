/* 
 * MetroTransit NexTrip API (BETA)
 *
 * ***    - BETA release of an OAuth protected iteration of the metrotransit Bus Routes API   - @iryanb (For Dev and Test Use Only, All Rights Reserved (c) 10.23.2017)  ***  ## MetroTransit NexTrip API (BETA)  ***    - Currently this API permits 2 requests per minute for anonymous API clients    - In a future release this API will offer rate limits for OAuth clients   - OAuth 2.0 AccessCode Tokens Scopes           - routes:all                  Grants read access to all routes       - routes:express              Grants read access to only express routes       - routes:local                Grants read access to local routes    # GET /NexTrip/{ROUTE}/{DIRECTION}/{STOP}    - GET TIMEPOINT DEPARTURES   - To make a valid request specify the integer values of the Route ID, Direction, and the Bus Stop ID as path parameters   - Without an Accept Header the default response schema is XML, for JSON set the Accept header to application/json       # GET /NexTrip/{STOPID}      - GET DEPARTURES   - This operation is used to return a list of departures scheduled for any given bus stop.    - A StopID is an integer value identifying any one of the many thousands of bus stops in the metro.    - Stop information can be derived from the GTFS schedule data updated weekly for public use.    - datafinder.org/metadata/transit_schedule_google_feed.html    # GET /NexTrip/Directions/{ROUTE}    - GET DIRECTIONS   - Returns the two directions that are valid for a given route. Either North/South or East/West.    - The result includes text/value pair with the direction name and an ID.    - Directions are identified with an ID value. 1 = South, 2 = East, 3 = West, 4 = North.     # GET /NexTrip/Providers    - GET PROVIDERS   - Returns a list of area Transit providers.    - Providers are identified in the list of Routes allowing routes to be selected for a single provider.     # GET /NexTrip/Routes    - GET ROUTES    - Returns a list of Transit routes that are in service on the current day.    # GET /NexTrip/Stops/{ROUTE}/{DIRECTION}    - GET STOPS   - Returns a list of Timepoint stops for the given Route/Direction.    - The result includes text/value pairs with the stop description and a 4 character stop (or node) identifier.     # GET /NexTrip/VehicleLocations/{ROUTE}    - GET VEHICLE LOCATIONS   - This operation returns a list of vehicles currently in service that have recently (within 5 minutes)    - reported their locations. A route paramter is used to return results for the given route.    - Use \"0\" for the route parameter to return a list of all vehicles in service.     ***    # POST /NexTrip/update/Routes (Future API Enhancement)       - This API operation requires an admin_AccessCode Token with one of the following scopes           - admin_updateroutes    - The POST message body must be a properly formatted JSON object with the following fields         ```                    {               \"Route\":\"integerRouteID\",                \"ProviderID\":\"integerProviderID\",                \"duration\":\"integerMinutes\",                \"stops\":[{\"Text\":\"stringStopName\", \"Value\":\"stringStopID\"},{\"Text\":\"stringStopName\", \"Value\":\"stringStopID\"}]               \"numberofstops\":\"integerStops\",                \"express\":\"True or False\",                \"Description\":\"Details about the route\"             }                    ```  ## Security Definitions        ***   # AccessCode        - Web and Mobile Authenticated Users              - routes:all                  Grants read access to all routes         - routes:express              Grants read access to only express routes         - routes:local                Grants read access to local routes         - readpublic_key              List and view details for public keys             ***   # MobileApp_Implicit     - Implicitly Trusted Mobile App with Anonymous User           - readpublic_key              List and view details for public keys            ***   # admin_AccessCode     - administrative purposes          - admin_updateroutes          Grants write access to routes data         - admin_writepublic_key       Create, list, and view details for public keys         - admin_public_key            Fully manage PKI keys    *** 
 *
 * OpenAPI spec version: BETA
 * 
 */

package swagger

import (
	"net/url"
	"strings"
	"encoding/json"
	"fmt"
)

type BusRoutesApi struct {
	Configuration *Configuration
}

func NewBusRoutesApi() *BusRoutesApi {
	configuration := NewConfiguration()
	return &BusRoutesApi{
		Configuration: configuration,
	}
}

func NewBusRoutesApiWithBasePath(basePath string) *BusRoutesApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &BusRoutesApi{
		Configuration: configuration,
	}
}

/**
 * 
 * Returns a list of departures scheduled for any given bus stop.
 *
 * @param sTOPID Specify the value of the Bus Stop ID as an abbreviated string
 * @return *Success
 */
func (a BusRoutesApi) GetDepartures(sTOPID string) (*Success, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/NexTrip/{STOPID}"
	localVarPath = strings.Replace(localVarPath, "{"+"STOPID"+"}", fmt.Sprintf("%v", sTOPID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(AccessCode)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Success)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetDepartures", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * Returns the two directions that are valid for a given route.
 *
 * @param rOUTE Sepcify the Route ID as an integer.
 * @return *Success
 */
func (a BusRoutesApi) GetDirections(rOUTE int32) (*Success, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/NexTrip/Directions/{ROUTE}"
	localVarPath = strings.Replace(localVarPath, "{"+"ROUTE"+"}", fmt.Sprintf("%v", rOUTE), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(AccessCode)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Success)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetDirections", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * Returns a list of area Transit providers.  Providers are identified in the list of Routes allowing routes to be selected for a single provider. 
 *
 * @return *Success
 */
func (a BusRoutesApi) GetProviders() (*Success, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/NexTrip/Providers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(AccessCode)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Success)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetProviders", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * Returns a list of Transit routes that are in service on the current day.
 *
 * @return *RouteData
 */
func (a BusRoutesApi) GetRoutes() (*RouteData, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/NexTrip/Routes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(AccessCode)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(RouteData)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetRoutes", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * Returns a list of Timepoint stops for the given Route/Direction.
 *
 * @param rOUTE Sepcify the Route ID as an integer.
 * @param dIRECTION Specify the direction as an integer.  1 (South), 2 (East), 3 (West), 4 (North)
 * @return *Success
 */
func (a BusRoutesApi) GetStops(rOUTE int32, dIRECTION int32) (*Success, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/NexTrip/Stops/{ROUTE}/{DIRECTION}"
	localVarPath = strings.Replace(localVarPath, "{"+"ROUTE"+"}", fmt.Sprintf("%v", rOUTE), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"DIRECTION"+"}", fmt.Sprintf("%v", dIRECTION), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(AccessCode)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Success)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetStops", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * Returns the scheduled departures for a selected route, direction and timepoint stop.
 *
 * @param rOUTE Sepcify the Route ID as an integer.
 * @param dIRECTION Specify the direction as an integer.  1 (South), 2 (East), 3 (West), 4 (North)
 * @param sTOP Specify the value of the Bus Stop ID as an abbreviated string
 * @return *TimePoints
 */
func (a BusRoutesApi) GetTimepointDepartures(rOUTE int32, dIRECTION int32, sTOP string) (*TimePoints, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/NexTrip/{ROUTE}/{DIRECTION}/{STOP}"
	localVarPath = strings.Replace(localVarPath, "{"+"ROUTE"+"}", fmt.Sprintf("%v", rOUTE), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"DIRECTION"+"}", fmt.Sprintf("%v", dIRECTION), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"STOP"+"}", fmt.Sprintf("%v", sTOP), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(AccessCode)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(TimePoints)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetTimepointDepartures", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 * 
 * This operation returns a list of vehicles currently in service that have recently (within 5 minutes)  reported their locations. A route paramter is used to return results for the given route.  Use \&quot;0\&quot; for the route parameter to return a list of all vehicles in service.
 *
 * @param rOUTE Sepcify the Route ID as an integer.
 * @return *Success
 */
func (a BusRoutesApi) GetVehicleLocations(rOUTE int32) (*Success, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/NexTrip/VehicleLocations/{ROUTE}"
	localVarPath = strings.Replace(localVarPath, "{"+"ROUTE"+"}", fmt.Sprintf("%v", rOUTE), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(AccessCode)' required
	// oauth required
	if a.Configuration.AccessToken != ""{
		localVarHeaderParams["Authorization"] =  "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Success)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "GetVehicleLocations", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

