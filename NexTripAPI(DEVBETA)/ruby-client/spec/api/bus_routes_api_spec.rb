=begin
#MetroTransit NexTrip API (BETA)

#***    - BETA release of an OAuth protected iteration of the metrotransit Bus Routes API   - @iryanb (For Dev and Test Use Only, All Rights Reserved (c) 10.23.2017)  ***  ## MetroTransit NexTrip API (BETA)  ***    - Currently this API permits 2 requests per minute for anonymous API clients    - In a future release this API will offer rate limits for OAuth clients   - OAuth 2.0 AccessCode Tokens Scopes           - routes:all                  Grants read access to all routes       - routes:express              Grants read access to only express routes       - routes:local                Grants read access to local routes    # GET /NexTrip/{ROUTE}/{DIRECTION}/{STOP}    - GET TIMEPOINT DEPARTURES   - To make a valid request specify the integer values of the Route ID, Direction, and the Bus Stop ID as path parameters   - Without an Accept Header the default response schema is XML, for JSON set the Accept header to application/json       # GET /NexTrip/{STOPID}      - GET DEPARTURES   - This operation is used to return a list of departures scheduled for any given bus stop.    - A StopID is an integer value identifying any one of the many thousands of bus stops in the metro.    - Stop information can be derived from the GTFS schedule data updated weekly for public use.    - datafinder.org/metadata/transit_schedule_google_feed.html    # GET /NexTrip/Directions/{ROUTE}    - GET DIRECTIONS   - Returns the two directions that are valid for a given route. Either North/South or East/West.    - The result includes text/value pair with the direction name and an ID.    - Directions are identified with an ID value. 1 = South, 2 = East, 3 = West, 4 = North.     # GET /NexTrip/Providers    - GET PROVIDERS   - Returns a list of area Transit providers.    - Providers are identified in the list of Routes allowing routes to be selected for a single provider.     # GET /NexTrip/Routes    - GET ROUTES    - Returns a list of Transit routes that are in service on the current day.    # GET /NexTrip/Stops/{ROUTE}/{DIRECTION}    - GET STOPS   - Returns a list of Timepoint stops for the given Route/Direction.    - The result includes text/value pairs with the stop description and a 4 character stop (or node) identifier.     # GET /NexTrip/VehicleLocations/{ROUTE}    - GET VEHICLE LOCATIONS   - This operation returns a list of vehicles currently in service that have recently (within 5 minutes)    - reported their locations. A route paramter is used to return results for the given route.    - Use \"0\" for the route parameter to return a list of all vehicles in service.     ***    # POST /NexTrip/update/Routes (Future API Enhancement)       - This API operation requires an admin_AccessCode Token with one of the following scopes           - admin_updateroutes    - The POST message body must be a properly formatted JSON object with the following fields         ```                    {               \"Route\":\"integerRouteID\",                \"ProviderID\":\"integerProviderID\",                \"duration\":\"integerMinutes\",                \"stops\":[{\"Text\":\"stringStopName\", \"Value\":\"stringStopID\"},{\"Text\":\"stringStopName\", \"Value\":\"stringStopID\"}]               \"numberofstops\":\"integerStops\",                \"express\":\"True or False\",                \"Description\":\"Details about the route\"             }                    ```  ## Security Definitions        ***   # AccessCode        - Web and Mobile Authenticated Users              - routes:all                  Grants read access to all routes         - routes:express              Grants read access to only express routes         - routes:local                Grants read access to local routes         - readpublic_key              List and view details for public keys             ***   # MobileApp_Implicit     - Implicitly Trusted Mobile App with Anonymous User           - readpublic_key              List and view details for public keys            ***   # admin_AccessCode     - administrative purposes          - admin_updateroutes          Grants write access to routes data         - admin_writepublic_key       Create, list, and view details for public keys         - admin_public_key            Fully manage PKI keys    *** 

OpenAPI spec version: BETA

Generated by: https://github.com/swagger-api/swagger-codegen.git
Swagger Codegen version: 2.2.3

=end

require 'spec_helper'
require 'json'

# Unit tests for SwaggerClient::BusRoutesApi
# Automatically generated by swagger-codegen (github.com/swagger-api/swagger-codegen)
# Please update as you see appropriate
describe 'BusRoutesApi' do
  before do
    # run before each test
    @instance = SwaggerClient::BusRoutesApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of BusRoutesApi' do
    it 'should create an instance of BusRoutesApi' do
      expect(@instance).to be_instance_of(SwaggerClient::BusRoutesApi)
    end
  end

  # unit tests for get_departures
  # 
  # Returns a list of departures scheduled for any given bus stop.
  # @param stopid Specify the value of the Bus Stop ID as an abbreviated string
  # @param [Hash] opts the optional parameters
  # @return [Success]
  describe 'get_departures test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_directions
  # 
  # Returns the two directions that are valid for a given route.
  # @param route Sepcify the Route ID as an integer.
  # @param [Hash] opts the optional parameters
  # @return [Success]
  describe 'get_directions test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_providers
  # 
  # Returns a list of area Transit providers.  Providers are identified in the list of Routes allowing routes to be selected for a single provider. 
  # @param [Hash] opts the optional parameters
  # @return [Success]
  describe 'get_providers test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_routes
  # 
  # Returns a list of Transit routes that are in service on the current day.
  # @param [Hash] opts the optional parameters
  # @return [RouteData]
  describe 'get_routes test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_stops
  # 
  # Returns a list of Timepoint stops for the given Route/Direction.
  # @param route Sepcify the Route ID as an integer.
  # @param direction Specify the direction as an integer.  1 (South), 2 (East), 3 (West), 4 (North)
  # @param [Hash] opts the optional parameters
  # @return [Success]
  describe 'get_stops test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_timepoint_departures
  # 
  # Returns the scheduled departures for a selected route, direction and timepoint stop.
  # @param route Sepcify the Route ID as an integer.
  # @param direction Specify the direction as an integer.  1 (South), 2 (East), 3 (West), 4 (North)
  # @param stop Specify the value of the Bus Stop ID as an abbreviated string
  # @param [Hash] opts the optional parameters
  # @return [TimePoints]
  describe 'get_timepoint_departures test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

  # unit tests for get_vehicle_locations
  # 
  # This operation returns a list of vehicles currently in service that have recently (within 5 minutes)  reported their locations. A route paramter is used to return results for the given route.  Use \&quot;0\&quot; for the route parameter to return a list of all vehicles in service.
  # @param route Sepcify the Route ID as an integer.
  # @param [Hash] opts the optional parameters
  # @return [Success]
  describe 'get_vehicle_locations test' do
    it "should work" do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end