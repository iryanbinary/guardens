<?php
require_once __DIR__ . '/vendor/autoload.php';

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Silex\Application;

$app = new Silex\Application();


$app->GET('/NexTrip/oauth20/authorize', function(Application $app, Request $request) {
            $grant_type = $request->get('grant_type');
            $client_id = $request->get('client_id');
            $redirect_uri = $request->get('redirect_uri');
            return new Response('code = 00000403');
            });


$app->GET('/NexTrip/oauth20/token', function(Application $app, Request $request) {
            $grant_type = $request->get('grant_type');
            $client_id = $request->get('client_id');
            $client_secret = $request->get('client_secret');
            return new Response('BEARER TODOINTEGRATEWITHOAUTHSERVER');
            });


$app->POST('/NexTrip/oauth20/token', function(Application $app, Request $request) {
            return new Response('BEARER TODOINTEGRATEWITHOAUTHSERVER');
            });


$app->POST('/NexTrip/update/{rOUTE}', function(Application $app, Request $request, $ROUTE) {
            return new Response('
                                    {"message":"Route Updated."}
                  ');
            });


$app->GET('/NexTrip/{sTOPID}', function(Application $app, Request $request, $STOPID) {
            return new Response('
                        [{"Actual":true,"BlockNumber":1419,"DepartureText":"12 Min","DepartureTime":"\/Date(1508861280000-0500)\/","Description":"Cedar-28Av\/VA Med Ctr\/Via Dupont-57","Gate":"","Route":"22","RouteDirection":"SOUTHBOUND","Terminal":"H","VehicleHeading":0,"VehicleLatitude":44.992100,"VehicleLongitude":-93.288080},{"Actual":true,"BlockNumber":1339,"DepartureText":"15 Min","DepartureTime":"\/Date(1508861460000-0500)\/","Description":"Chicago \/ Mall of Amer","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"E","VehicleHeading":0,"VehicleLatitude":45.001610,"VehicleLongitude":-93.295700},{"Actual":true,"BlockNumber":1121,"DepartureText":"11:21","DepartureTime":"\/Date(1508862060000-0500)\/","Description":"Chicago - 56St","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"B","VehicleHeading":0,"VehicleLatitude":45.005830,"VehicleLongitude":-93.309470},{"Actual":false,"BlockNumber":1491,"DepartureText":"11:29","DepartureTime":"\/Date(1508862540000-0500)\/","Description":"Cedar-28Av\/VA Med Ctr\/Via 51-Penn","Gate":"","Route":"22","RouteDirection":"SOUTHBOUND","Terminal":"H","VehicleHeading":0,"VehicleLatitude":45.046880,"VehicleLongitude":-93.309280},{"Actual":false,"BlockNumber":1127,"DepartureText":"11:30","DepartureTime":"\/Date(1508862600000-0500)\/","Description":"Chicago \/ Mall of Amer","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"E","VehicleHeading":0,"VehicleLatitude":45.060180,"VehicleLongitude":-93.319900},{"Actual":false,"BlockNumber":1131,"DepartureText":"11:40","DepartureTime":"\/Date(1508863200000-0500)\/","Description":"Chicago - 56St","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"B","VehicleHeading":0,"VehicleLatitude":45.059460,"VehicleLongitude":-93.318410},{"Actual":false,"BlockNumber":1501,"DepartureText":"11:47","DepartureTime":"\/Date(1508863620000-0500)\/","Description":"Cedar-28Av\/VA Med Ctr\/Via Logan-57","Gate":"","Route":"22","RouteDirection":"SOUTHBOUND","Terminal":"H","VehicleHeading":0,"VehicleLatitude":45.059670,"VehicleLongitude":-93.318820},{"Actual":false,"BlockNumber":1134,"DepartureText":"11:51","DepartureTime":"\/Date(1508863860000-0500)\/","Description":"Chicago \/ Mall of Amer","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"E","VehicleHeading":0,"VehicleLatitude":44.966640,"VehicleLongitude":-93.262490},{"Actual":false,"BlockNumber":1115,"DepartureText":"12:00","DepartureTime":"\/Date(1508864400000-0500)\/","Description":"Chicago - 56St","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"B","VehicleHeading":0,"VehicleLatitude":45.012470,"VehicleLongitude":-93.294330},{"Actual":false,"BlockNumber":1507,"DepartureText":"12:07","DepartureTime":"\/Date(1508864820000-0500)\/","Description":"Cedar-28Av\/VA Med Ctr\/Via Dupont-57","Gate":"","Route":"22","RouteDirection":"SOUTHBOUND","Terminal":"H","VehicleHeading":0,"VehicleLatitude":0.000000,"VehicleLongitude":0.000000},{"Actual":false,"BlockNumber":1117,"DepartureText":"12:10","DepartureTime":"\/Date(1508865000000-0500)\/","Description":"Chicago \/ Mall of Amer","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"E","VehicleHeading":0,"VehicleLatitude":45.019240,"VehicleLongitude":-93.288000},{"Actual":false,"BlockNumber":1143,"DepartureText":"12:21","DepartureTime":"\/Date(1508865660000-0500)\/","Description":"Chicago - 56St","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"B","VehicleHeading":0,"VehicleLatitude":44.883920,"VehicleLongitude":-93.267740},{"Actual":false,"BlockNumber":2320,"DepartureText":"12:27","DepartureTime":"\/Date(1508866020000-0500)\/","Description":"Cedar-28Av\/VA Med Ctr\/Via 51-Penn","Gate":"","Route":"22","RouteDirection":"SOUTHBOUND","Terminal":"H","VehicleHeading":0,"VehicleLatitude":45.070170,"VehicleLongitude":-93.286130},{"Actual":false,"BlockNumber":1125,"DepartureText":"12:30","DepartureTime":"\/Date(1508866200000-0500)\/","Description":"Chicago \/ Mall of Amer","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"E","VehicleHeading":0,"VehicleLatitude":44.962370,"VehicleLongitude":-93.262500},{"Actual":false,"BlockNumber":1122,"DepartureText":"12:40","DepartureTime":"\/Date(1508866800000-0500)\/","Description":"Chicago - 56St","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"B","VehicleHeading":0,"VehicleLatitude":44.928560,"VehicleLongitude":-93.262460},{"Actual":false,"BlockNumber":1508,"DepartureText":"12:47","DepartureTime":"\/Date(1508867220000-0500)\/","Description":"Cedar-28Av\/VA Med Ctr\/Via Logan-57","Gate":"","Route":"22","RouteDirection":"SOUTHBOUND","Terminal":"H","VehicleHeading":0,"VehicleLatitude":0.000000,"VehicleLongitude":0.000000},{"Actual":false,"BlockNumber":1147,"DepartureText":"12:51","DepartureTime":"\/Date(1508867460000-0500)\/","Description":"Chicago \/ Mall of Amer","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"E","VehicleHeading":0,"VehicleLatitude":0.000000,"VehicleLongitude":0.000000},{"Actual":false,"BlockNumber":1138,"DepartureText":"1:00","DepartureTime":"\/Date(1508868000000-0500)\/","Description":"Chicago - 56St","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"B","VehicleHeading":0,"VehicleLatitude":44.919120,"VehicleLongitude":-93.262610},{"Actual":false,"BlockNumber":1493,"DepartureText":"1:07","DepartureTime":"\/Date(1508868420000-0500)\/","Description":"Cedar-28Av\/VA Med Ctr\/Via Dupont-57","Gate":"","Route":"22","RouteDirection":"SOUTHBOUND","Terminal":"H","VehicleHeading":0,"VehicleLatitude":44.905320,"VehicleLongitude":-93.203010},{"Actual":false,"BlockNumber":1144,"DepartureText":"1:10","DepartureTime":"\/Date(1508868600000-0500)\/","Description":"Chicago \/ Mall of Amer","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"E","VehicleHeading":0,"VehicleLatitude":44.853330,"VehicleLongitude":-93.238390},{"Actual":false,"BlockNumber":1150,"DepartureText":"1:21","DepartureTime":"\/Date(1508869260000-0500)\/","Description":"Chicago - 56St","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"B","VehicleHeading":0,"VehicleLatitude":0.000000,"VehicleLongitude":0.000000},{"Actual":false,"BlockNumber":1509,"DepartureText":"1:24","DepartureTime":"\/Date(1508869440000-0500)\/","Description":"Cedar-28Av\/VA Med Ctr\/Via 51-Penn","Gate":"","Route":"22","RouteDirection":"SOUTHBOUND","Terminal":"H","VehicleHeading":0,"VehicleLatitude":0.000000,"VehicleLongitude":0.000000},{"Actual":false,"BlockNumber":1145,"DepartureText":"1:30","DepartureTime":"\/Date(1508869800000-0500)\/","Description":"Chicago \/ Mall of Amer","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"E","VehicleHeading":0,"VehicleLatitude":0.000000,"VehicleLongitude":0.000000},{"Actual":false,"BlockNumber":1140,"DepartureText":"1:40","DepartureTime":"\/Date(1508870400000-0500)\/","Description":"Chicago - 56St","Gate":"","Route":"5","RouteDirection":"SOUTHBOUND","Terminal":"B","VehicleHeading":0,"VehicleLatitude":0.000000,"VehicleLongitude":0.000000},{"Actual":false,"BlockNumber":2127,"DepartureText":"1:43","DepartureTime":"\/Date(1508870580000-0500)\/","Description":"Cedar-28Av\/VA Med Ctr\/Via Logan-57","Gate":"","Route":"22","RouteDirection":"SOUTHBOUND","Terminal":"H","VehicleHeading":0,"VehicleLatitude":45.094370,"VehicleLongitude":-93.371950}]
                  ');
            });


$app->GET('/NexTrip/Directions/{rOUTE}', function(Application $app, Request $request, $ROUTE) {
            return new Response('
                        [{"Text":"NORTHBOUND","Value":"4"},{"Text":"SOUTHBOUND","Value":"1"}]
                        ');
            });


$app->GET('/NexTrip/Providers', function(Application $app, Request $request) {
            return new Response('
                        [{"Text":"University of Minnesota","Value":"1"},{"Text":"Airport (MAC)","Value":"2"},{"Text":"Other","Value":"3"},{"Text":"Prior Lake","Value":"4"},{"Text":"Scott County","Value":"5"},{"Text":"Metro Transit\/Met Council","Value":"6"},{"Text":"SouthWest Transit","Value":"7"},{"Text":"Metro Transit","Value":"8"},{"Text":"Minnesota Valley","Value":"9"},{"Text":"Plymouth","Value":"10"},{"Text":"Met Council","Value":"11"},{"Text":"Maple Grove","Value":"12"}]
                  ');
            });


$app->GET('/NexTrip/Routes', function(Application $app, Request $request) {
            return new Response('
                                    [{"Description":"METRO Blue Line","ProviderID":"8","Route":"901"},{"Description":"METRO Green Line","ProviderID":"8","Route":"902"},{"Description":"METRO Red Line","ProviderID":"9","Route":"903"},{"Description":"2 - Franklin Av - Riverside Av - U of M - 8th St SE","ProviderID":"8","Route":"2"},{"Description":"3 - U of M - Como Av - Energy Park Dr - Maryland Av","ProviderID":"8","Route":"3"},{"Description":"4 - New Brighton - Johnson St - Bryant Av - Southtown","ProviderID":"8","Route":"4"},{"Description":"5 - Brklyn Center - Fremont - 26th Av - Chicago - MOA","ProviderID":"8","Route":"5"},{"Description":"6 - U of M - Hennepin - Xerxes - France - Southdale","ProviderID":"8","Route":"6"},{"Description":"7 - Plymouth - 27Av - Midtown - 46St LRT - 34Av S","ProviderID":"8","Route":"7"},{"Description":"9 - Glenwood Av - Wayzata Blvd - Cedar Lk Rd -46St LRT","ProviderID":"8","Route":"9"},{"Description":"10 - Central Av - University Av - Northtown","ProviderID":"8","Route":"10"},{"Description":"11 - Columbia Heights - 2nd St NE - 4th Av S","ProviderID":"8","Route":"11"},{"Description":"12 - Uptown - Excelsior Blvd - Hopkins - Opus","ProviderID":"8","Route":"12"},{"Description":"14 - Robbinsdale-West Broadway-Bloomington Av","ProviderID":"8","Route":"14"},{"Description":"16 - U of M - University Av - Midway","ProviderID":"8","Route":"16"},{"Description":"17 - Minnetonka Blvd - Uptown - Washington St NE","ProviderID":"8","Route":"17"},{"Description":"18 - Nicollet Av - South Bloomington","ProviderID":"8","Route":"18"},{"Description":"19 - Olson Memorial Hwy - Penn Av N - Brooklyn Center"}]
                                    ');
            });


$app->GET('/NexTrip/Stops/{rOUTE}/{dIRECTION}', function(Application $app, Request $request, $ROUTE, $DIRECTION) {
            return new Response('
                        [{"Text":"Target North Campus","Value":"TARG"},{"Text":"Wyoming Ave  and 93rd Ave ","Value":"93WY"},{"Text":"W Broadway Ave and 84th Ave ","Value":"BR84"},{"Text":"Starlite Transit Center","Value":"STLI"},{"Text":"63rd Ave  and Zane Ave ","Value":"63ZA"},{"Text":"Brooklyn Center Transit Center","Value":"BCTC"},{"Text":"Osseo Rd and 47th Ave ","Value":"47OS"},{"Text":"Dowling Ave  and Fremont Ave ","Value":"FRDO"},{"Text":"Hennepin Ave and 6th St","Value":"6SHE"},{"Text":"9th Ave and 7th St","Value":"9A7S"}]
                  ');
            });


$app->GET('/NexTrip/{rOUTE}/{dIRECTION}/{sTOP}', function(Application $app, Request $request, $ROUTE, $DIRECTION, $STOP) {
            return new Response('    
                              {
                          "Actual": false,
                          "BlockNumber": 1434,
                          "DepartureText": "12:35",
                          "DepartureTime": "/Date(1508780100000-0500)/",
                          "Description": "Ltd Stop/Brklyn Ctr/Target N Campus",
                          "Gate": "",
                          "Route": "724",
                          "RouteDirection": "NORTHBOUND",
                          "Terminal": "C",
                          "VehicleHeading": 0,
                          "VehicleLatitude": 45.06056,
                          "VehicleLongitude": -93.31913
                              }
    ');
            });


$app->GET('/NexTrip/VehicleLocations/{rOUTE}', function(Application $app, Request $request, $ROUTE) {
            return new Response('
                        [{"Bearing":225,"BlockNumber":1142,"Direction":1,"LocationTime":"\/Date(1508860860000-0500)\/","Odometer":0,"Route":"724","Speed":1.72603109E-07,"Terminal":"D","VehicleLatitude":44.972400,"VehicleLongitude":-93.259340},{"Bearing":315,"BlockNumber":1424,"Direction":4,"LocationTime":"\/Date(1508860895000-0500)\/","Odometer":0,"Route":"724","Speed":2.58904674E-06,"Terminal":"","VehicleLatitude":44.978010,"VehicleLongitude":-93.275220},{"Bearing":225,"BlockNumber":1495,"Direction":1,"LocationTime":"\/Date(1508860950000-0500)\/","Odometer":0,"Route":"724","Speed":0,"Terminal":"D","VehicleLatitude":45.054520,"VehicleLongitude":-93.321720},{"Bearing":270,"BlockNumber":2124,"Direction":4,"LocationTime":"\/Date(1508860935000-0500)\/","Odometer":0,"Route":"724","Speed":5.17809349E-06,"Terminal":"","VehicleLatitude":45.069340,"VehicleLongitude":-93.325650},{"Bearing":315,"BlockNumber":2305,"Direction":4,"LocationTime":"\/Date(1508860915000-0500)\/","Odometer":0,"Route":"724","Speed":0,"Terminal":"","VehicleLatitude":45.093520,"VehicleLongitude":-93.380340}]
                  ');
            });


$app->run();
