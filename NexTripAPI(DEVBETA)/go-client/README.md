# Go API client for swagger

***    - BETA release of an OAuth protected iteration of the metrotransit Bus Routes API   - @iryanb (For Dev and Test Use Only, All Rights Reserved (c) 10.23.2017)  ***  ## MetroTransit NexTrip API (BETA)  ***    - Currently this API permits 2 requests per minute for anonymous API clients    - In a future release this API will offer rate limits for OAuth clients   - OAuth 2.0 AccessCode Tokens Scopes           - routes:all                  Grants read access to all routes       - routes:express              Grants read access to only express routes       - routes:local                Grants read access to local routes    # GET /NexTrip/{ROUTE}/{DIRECTION}/{STOP}    - GET TIMEPOINT DEPARTURES   - To make a valid request specify the integer values of the Route ID, Direction, and the Bus Stop ID as path parameters   - Without an Accept Header the default response schema is XML, for JSON set the Accept header to application/json       # GET /NexTrip/{STOPID}      - GET DEPARTURES   - This operation is used to return a list of departures scheduled for any given bus stop.    - A StopID is an integer value identifying any one of the many thousands of bus stops in the metro.    - Stop information can be derived from the GTFS schedule data updated weekly for public use.    - datafinder.org/metadata/transit_schedule_google_feed.html    # GET /NexTrip/Directions/{ROUTE}    - GET DIRECTIONS   - Returns the two directions that are valid for a given route. Either North/South or East/West.    - The result includes text/value pair with the direction name and an ID.    - Directions are identified with an ID value. 1 = South, 2 = East, 3 = West, 4 = North.     # GET /NexTrip/Providers    - GET PROVIDERS   - Returns a list of area Transit providers.    - Providers are identified in the list of Routes allowing routes to be selected for a single provider.     # GET /NexTrip/Routes    - GET ROUTES    - Returns a list of Transit routes that are in service on the current day.    # GET /NexTrip/Stops/{ROUTE}/{DIRECTION}    - GET STOPS   - Returns a list of Timepoint stops for the given Route/Direction.    - The result includes text/value pairs with the stop description and a 4 character stop (or node) identifier.     # GET /NexTrip/VehicleLocations/{ROUTE}    - GET VEHICLE LOCATIONS   - This operation returns a list of vehicles currently in service that have recently (within 5 minutes)    - reported their locations. A route paramter is used to return results for the given route.    - Use \"0\" for the route parameter to return a list of all vehicles in service.     ***    # POST /NexTrip/update/Routes (Future API Enhancement)       - This API operation requires an admin_AccessCode Token with one of the following scopes           - admin_updateroutes    - The POST message body must be a properly formatted JSON object with the following fields         ```                    {               \"Route\":\"integerRouteID\",                \"ProviderID\":\"integerProviderID\",                \"duration\":\"integerMinutes\",                \"stops\":[{\"Text\":\"stringStopName\", \"Value\":\"stringStopID\"},{\"Text\":\"stringStopName\", \"Value\":\"stringStopID\"}]               \"numberofstops\":\"integerStops\",                \"express\":\"True or False\",                \"Description\":\"Details about the route\"             }                    ```  ## Security Definitions        ***   # AccessCode        - Web and Mobile Authenticated Users              - routes:all                  Grants read access to all routes         - routes:express              Grants read access to only express routes         - routes:local                Grants read access to local routes         - readpublic_key              List and view details for public keys             ***   # MobileApp_Implicit     - Implicitly Trusted Mobile App with Anonymous User           - readpublic_key              List and view details for public keys            ***   # admin_AccessCode     - administrative purposes          - admin_updateroutes          Grants write access to routes data         - admin_writepublic_key       Create, list, and view details for public keys         - admin_public_key            Fully manage PKI keys    *** 

## Overview

- API version: BETA
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://svc.metrotransit.org*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccessControlApi* | [**GetAuthCode**](docs/AccessControlApi.md#getauthcode) | **Get** /NexTrip/oauth20/authorize | 
*AccessControlApi* | [**GetTokenRequest**](docs/AccessControlApi.md#gettokenrequest) | **Get** /NexTrip/oauth20/token | 
*AccessControlApi* | [**PostTokenRequest**](docs/AccessControlApi.md#posttokenrequest) | **Post** /NexTrip/oauth20/token | 
*BusRouteInfoApi* | [**Updateroute**](docs/BusRouteInfoApi.md#updateroute) | **Post** /NexTrip/update/{ROUTE} | 
*BusRoutesApi* | [**GetDepartures**](docs/BusRoutesApi.md#getdepartures) | **Get** /NexTrip/{STOPID} | 
*BusRoutesApi* | [**GetDirections**](docs/BusRoutesApi.md#getdirections) | **Get** /NexTrip/Directions/{ROUTE} | 
*BusRoutesApi* | [**GetProviders**](docs/BusRoutesApi.md#getproviders) | **Get** /NexTrip/Providers | 
*BusRoutesApi* | [**GetRoutes**](docs/BusRoutesApi.md#getroutes) | **Get** /NexTrip/Routes | 
*BusRoutesApi* | [**GetStops**](docs/BusRoutesApi.md#getstops) | **Get** /NexTrip/Stops/{ROUTE}/{DIRECTION} | 
*BusRoutesApi* | [**GetTimepointDepartures**](docs/BusRoutesApi.md#gettimepointdepartures) | **Get** /NexTrip/{ROUTE}/{DIRECTION}/{STOP} | 
*BusRoutesApi* | [**GetVehicleLocations**](docs/BusRoutesApi.md#getvehiclelocations) | **Get** /NexTrip/VehicleLocations/{ROUTE} | 


## Documentation For Models

 - [Err](docs/Err.md)
 - [OAuthToken](docs/OAuthToken.md)
 - [RouteData](docs/RouteData.md)
 - [Success](docs/Success.md)
 - [TimePoints](docs/TimePoints.md)


## Documentation For Authorization


## AccessCode

- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://svc.metrotransit.org/oauth/authorize
- **Scopes**: 
 - **routes:all**: Grants read access to all routes
 - **routes:express**: Grants read access to only express routes
 - **routes:local**: Grants read access to local routes
 - **readpublic_key**: List and view details for public keys

## MobileApp_Implicit

- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://svc.metrotransit.org/oauth/authorize
- **Scopes**: 
 - **readpublic_key**: List and view details for public keys

## admin_AccessCode

- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://svc.metrotransit.org/oauth/authorize
- **Scopes**: 
 - **admin_updateroutes**: Grants write access to routes data
 - **admin_writepublic_key**: Create, list, and view details for public keys
 - **admin_public_key**: Fully manage PKI keys


## Author



