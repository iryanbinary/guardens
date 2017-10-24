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
	"bytes"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"net/url"
	"io/ioutil"
	"gopkg.in/go-resty/resty.v0"
)

type APIClient struct {
	config *Configuration
}

func (c *APIClient) SelectHeaderContentType(contentTypes []string) string {

	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

func (c *APIClient) SelectHeaderAccept(accepts []string) string {

	if len(accepts) == 0 {
		return ""
	}
	if contains(accepts, "application/json") {
		return "application/json"
	}
	return strings.Join(accepts, ",")
}

func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

func (c *APIClient) CallAPI(path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	fileName string,
	fileBytes []byte) (*resty.Response, error) {

	rClient := c.prepareClient()
	request := c.prepareRequest(rClient, postBody, headerParams, queryParams, formParams, fileName, fileBytes)

	switch strings.ToUpper(method) {
	case "GET":
		response, err := request.Get(path)
		return response, err
	case "POST":
		response, err := request.Post(path)
		return response, err
	case "PUT":
		response, err := request.Put(path)
		return response, err
	case "PATCH":
		response, err := request.Patch(path)
		return response, err
	case "DELETE":
		response, err := request.Delete(path)
		return response, err
	}

	return nil, fmt.Errorf("invalid method %v", method)
}

func (c *APIClient) ParameterToString(obj interface{}, collectionFormat string) string {
	delimiter := ""
	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

func (c *APIClient) prepareClient() *resty.Client {

	rClient := resty.New()

	rClient.SetDebug(c.config.Debug)
	if c.config.Transport != nil {
		rClient.SetTransport(c.config.Transport)
	}

	if c.config.Timeout != nil {
		rClient.SetTimeout(*c.config.Timeout)
	}
	rClient.SetLogger(ioutil.Discard)
	return rClient
}

func (c *APIClient) prepareRequest(
	rClient *resty.Client,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams map[string]string,
	fileName string,
	fileBytes []byte) *resty.Request {


	request := rClient.R()
	request.SetBody(postBody)

	if c.config.UserAgent != "" {
		request.SetHeader("User-Agent", c.config.UserAgent)
	}

	// add header parameter, if any
	if len(headerParams) > 0 {
		request.SetHeaders(headerParams)
	}

	// add query parameter, if any
	if len(queryParams) > 0 {
		request.SetMultiValueQueryParams(queryParams)
	}

	// add form parameter, if any
	if len(formParams) > 0 {
		request.SetFormData(formParams)
	}

	if len(fileBytes) > 0 && fileName != "" {
		_, fileNm := filepath.Split(fileName)
		request.SetFileReader("file", fileNm, bytes.NewReader(fileBytes))
	}
	return request
}
