# Author: @iryanb 
# Next Bus using MetroTransit NexTrip API 
# Copyright 2017 All Rights Reserved. 
# Usage License: MIT

require 'rest-client'
require 'json'
require 'optparse'
# require 'time'

searchtime = Time.now

# puts searchtime.to_s

options = {}
OptionParser.new do |opts|
  opts.banner = "Usage: example.rb [options]"

  opts.on("-v", "--[no-]verbose", "Run verbosely") do |v|
    options[:verbose] = v
  end
end.parse!

# Route ID
# p ARGV[0] 

# Direction 1 = South, 2 = East, 3 = West, 4 = North
# p ARGV[1]

# Stop ID 
# p ARGV[2] 

responseMsg = RestClient.get "http://svc.metrotransit.org/NexTrip/" + ARGV[0] + '/' + ARGV[1] + '/' + ARGV[2], 'Accept'=>"application/json"
responseMsgparsed = JSON.parse(responseMsg)
nextBus = responseMsgparsed[0]["DepartureText"]
nextBusDirection = responseMsgparsed[0]["RouteDirection"]
nextBusDescription = responseMsgparsed[0]["Description"]

# note the json response format of the DepartureTime Key must be transformed as it is not a standard datetime object 
# nextBusTime = responseMsgparsed[0]["DepartureTime"] 


n=0                                                           
1.times do                                                          
# puts responseMsg.body
# puts responseMsgparsed
# puts nextBus

elapsed = Time.now - searchtime
n=n+1

puts n.to_s + " The Next Bus on Route " + ARGV[0] + " going " + nextBusDirection + " " + nextBusDescription + " departs from " + ARGV[2] + " at " + nextBus 
end                                                                  
puts Time.now.to_s 

