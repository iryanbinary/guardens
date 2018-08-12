# Author: Ryan Bagnulo @iryanb 
# API emulator
# Copyright 2018 All Rights Reserved. guarden.io 
# Usage License: MIT

require 'rest-client'
#require "net/http"
require 'json'
require 'optparse'
require 'time'

gettime = Time.now

puts gettime.to_s

#uri = URI("http://api.spitcast.com/api/top/spots")

#http = Net::HTTP.new(uri.host, uri.port)

#req = Net::HTTP::get.new(uri.path)
#req.content_type = "application/json"

n=0                                                           
1.times do                                                          
#http.request(req)

responseMsg = RestClient.get "http://api.spitcast.com/api/top/spots", 'Accept'=>"application/json"
responseMsgparsed = JSON.parse(responseMsg)
fieldA = responseMsgparsed[0]["avg_max_size"]
#fieldB = responseMsgparsed[0]["FieldB"]
#fieldC = responseMsgparsed[0]["FieldC"]

n=n+1

#parsed
#puts n.to_s + " Response Message "  + fieldA + " " + fieldB + " " + fieldC 

puts n.to_s + " Response Message " + responseMsgparsed.to_s + " Avg Wave Size = " + fieldA.to_s 

end                                                                  

elapsed = Time.now - gettime

puts "This test completed " + n.to_s + "api requests " + " in " + elapsed.to_s + " seconds." 

puts "The Time now is " + Time.now.to_s 
