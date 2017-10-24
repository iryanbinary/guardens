<?php
require_once __DIR__ . '/vendor/autoload.php';

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Silex\Application;

$app = new Silex\Application();


$app->GET('//NexTrip/oauth20/authorize', function(Application $app, Request $request) {
            $grant_type = $request->get('grant_type');
            $client_id = $request->get('client_id');
            $redirect_uri = $request->get('redirect_uri');
            return new Response('How about implementing getAuthCode as a GET method ?');
            });


$app->GET('//NexTrip/oauth20/token', function(Application $app, Request $request) {
            $grant_type = $request->get('grant_type');
            $client_id = $request->get('client_id');
            $client_secret = $request->get('client_secret');
            return new Response('How about implementing getTokenRequest as a GET method ?');
            });


$app->POST('//NexTrip/oauth20/token', function(Application $app, Request $request) {
            return new Response('How about implementing postTokenRequest as a POST method ?');
            });


$app->POST('//NexTrip/update/{rOUTE}', function(Application $app, Request $request, $ROUTE) {
            return new Response('How about implementing updateroute as a POST method ?');
            });


$app->GET('//NexTrip/{sTOPID}', function(Application $app, Request $request, $STOPID) {
            return new Response('How about implementing getDepartures as a GET method ?');
            });


$app->GET('//NexTrip/Directions/{rOUTE}', function(Application $app, Request $request, $ROUTE) {
            return new Response('How about implementing getDirections as a GET method ?');
            });


$app->GET('//NexTrip/Providers', function(Application $app, Request $request) {
            return new Response('How about implementing getProviders as a GET method ?');
            });


$app->GET('//NexTrip/Routes', function(Application $app, Request $request) {
            return new Response('How about implementing getRoutes as a GET method ?');
            });


$app->GET('//NexTrip/Stops/{rOUTE}/{dIRECTION}', function(Application $app, Request $request, $ROUTE, $DIRECTION) {
            return new Response('How about implementing getStops as a GET method ?');
            });


$app->GET('//NexTrip/{rOUTE}/{dIRECTION}/{sTOP}', function(Application $app, Request $request, $ROUTE, $DIRECTION, $STOP) {
            return new Response('How about implementing getTimepointDepartures as a GET method ?');
            });


$app->GET('//NexTrip/VehicleLocations/{rOUTE}', function(Application $app, Request $request, $ROUTE) {
            return new Response('How about implementing getVehicleLocations as a GET method ?');
            });


$app->run();
