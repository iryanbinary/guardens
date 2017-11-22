# Author: @iryanb Fibonacci Bench Score
# Software License: MIT
# Copyright 2017 guarden.io founder Ryan Bagnulo
# All Rights Reserved. 

require 'prime'

start = Time.now
puts start.to_s
n=0
x,y=0,1                                                              
1000000.times do                                                          
puts y
elapsed = Time.now - start

p=Prime.instance.prime?(y).to_s
n=n+1
l=y.to_s.length
puts n.to_s + " TIME:" + Time.now.to_s + " ELAPSED :" + elapsed.to_s + " seconds.  LENGTH = " + l.to_s + " Prime: " + p
x,y=y,x+y
end                                                                  
puts Time.now.to_s 

## Baseline = computes the first 93 fibonacci numbers and finds 12 of them are prime in less than 110 seconds
