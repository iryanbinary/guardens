# Author: Ryan Bagnulo @iryanb 
# Grobot Supply Chain Automation Fork Lift Emulator
# Copyright 2018 All Rights Reserved. guarden.io 
# Usage License: MIT
# This emulator is designed to send a REST API POST request to an influxdb time series database in binary octet stream format.
# Transformation of the data to json is more effeciently done by the user interface after the data is persisted.

#require 'rest-client'
require "net/http"
#require 'json'
#require 'optparse'
require 'time'

posttime = Time.now

puts posttime.to_s

uri = URI("http://localhost:99/write?db=SCAinfluxdb")

http = Net::HTTP.new(uri.host, uri.port)

req = Net::HTTP::Post.new(uri.path)
req.body = "\x154\x151\x146\x164\x167\x157\x162\x153\x54\x114\x151\x146\x164\x127\x157\x162\x153\x104\x145\x163\x143\x162\x151\x160\x164\x151\x157\x156\x75\x167\x157\x162\x153\x144\x145\x163\x143\x162\x151\x160\x164\x151\x157\x156\x164\x145\x170\x164\x54\x114\x151\x146\x164\x127\x157\x162\x153\x111\x144\x75\x167\x157\x162\x153\x151\x144\x164\x145\x170\x164\x54\x114\x151\x146\x164\x127\x157\x162\x153\x123\x164\x141\x162\x164\x124\x151\x155\x145\x75\x154\x151\x146\x164\x167\x157\x162\x153\x163\x164\x141\x162\x164\x164\x151\x155\x145\x164\x145\x170\x164\x54\x114\x151\x146\x164\x127\x157\x162\x153\x123\x164\x157\x160\x124\x151\x155\x145\x75\x154\x151\x146\x164\x167\x157\x162\x153\x163\x164\x157\x160\x164\x151\x155\x145\x164\x145\x170\x164\x54\x114\x151\x146\x164\x127\x157\x162\x153\x114\x157\x143\x141\x164\x151\x157\x156\x111\x144\x75\x154\x151\x146\x164\x167\x157\x162\x153\x154\x157\x143\x141\x164\x151\x157\x156\x164\x145\x170\x164\x54\x114\x151\x146\x164\x114\x157\x141\x144\x103\x145\x156\x164\x145\x162\x75\x60\x56\x60\x60\x54\x114\x151\x146\x164\x127\x157\x162\x153\x127\x145\x151\x147\x150\x164\x75\x60\x56\x60\x60"
req.content_type = "application/octet-stream"

#equivalent curl command
#http://localhost:99/write?db=SCAinfluxdb' --data-binary 'liftwork,LiftWorkDescription=workdescriptiontext,LiftWorkId=workidtext,LiftWorkStartTime=liftworkstarttimetext,LiftWorkStopTime=liftworkstoptimetext,LiftWorkLocationId=liftworklocationtext,LiftLoadCenter=0.00,LiftWorkWeight=0.00'

n=0
#to emulate 4000 lifts sending a POST work status update                                                            
#4000.times do     

#only one                                                      
1.times do 
http.request(req)

n=n+1

end                                                                  

elapsed = Time.now - posttime

puts "This test completed " + n.to_s + " fork lift work status POST emulated messages " + " in " + elapsed.to_s + " seconds." 
