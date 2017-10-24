# publicguardens

alpha development examples

for dev and test purposes only, all rights reserved, iryanb at guarden dot io (c) 2017

Usage License: MIT 

Copyright 2017 Ryan A. Bagnulo

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


ruby API Client for MetroTransit NexTrip Bus Routes REST API 

1.5 KBytes nextbus.rb
less than 50 lines of code 


usage: 

ruby nextbus.rb ROUTEID DIRECTION STOPID 

ROUTEID is an integer
DIRECTION is an integer 1 = South, 2 = East, 3 = West, 4 = North
STOPID is a string abbreviation for the desired bus stop 

Example API Request: 

$ ruby nextbus.rb 901 1 TF22

Response: 

1 The Next Bus on Route 901 going SOUTHBOUND to Mall of America departs from TF22 at 10:37
2017-10-23 22:19:22 -0600


Also see the extra credit work that I did to create a Swagger 2.0 powered interactive Web UI API client with user friendly documentation and API server iteration suggestions to implement OAuth 2.0 with scoped tokens for finer grained client application authorization and a new update route POST operation that requires the usage of a admin token to futher restrict access to higher risk API operations that update data in the database. 

Find this work in NexTripAPI_RubyCommandLineClientApp / Swagger+docs

MetroTransit NexTrip API (Future Swagger UI Client with OAuth).pdf
MetroTransit NexTrip API (Interactive API Web UI Client with example responses).pdf
index.html
metrotransitnextripAPIbeta.json
metrotransitnextripAPIbeta.yaml


