# What if the 1st two Fibonacci numbers are .072... and 1.17... instead of 1 and 1 
# if 1 = 0.72723460248141
# if 2 = 1.17669030460994
# Software License: MIT
# Copyright 2017 guarden.io founder Ryan Bagnulo
# All Rights Reserved. 

start = Time.now
puts start.to_s
n=1
x,y=0.72723460248141,1.17669030460994                                                              
puts x 
puts n.to_s + " TIME = " + Time.now.to_s + " ELAPSED = 0 LEN
GTH = 1"

1477.times do                                                          
puts y
elapsed = Time.now - start
n=n+1
l=y.to_s.length
r=y/x
puts n.to_s + " TIME:" + Time.now.to_s + " ELAPSED :" + elapsed.to_s + " s.  LENGTH = " + l.to_s + " deciPHI.r = " + r.to_s
x,y=y,x+y
end                                                                  
puts Time.now.to_s 

## Baseline = 1475 numbers to floating point decimal infinity in less than 0.06 seconds 
# (floating point PHI math performance test)
# 1475 TIME:2017-11-21 14:44:33 -0600 ELAPSED :0.058206 s.  LENGTH = 22 deciPHI.r = 1.618033988749895
# 1.3135418179411196e+308
# 1476 TIME:2017-11-21 14:44:33 -0600 ELAPSED :0.058236 s.  LENGTH = 23 deciPHI.r = 1.618033988749895
# Infinity
# 1477 TIME:2017-11-21 14:44:33 -0600 ELAPSED :0.058261 s.  LENGTH = 8 deciPHI.r = Infinity
# Infinity
# 1478 TIME:2017-11-21 14:44:33 -0600 ELAPSED :0.05828 s.  LENGTH = 8 deciPHI.r = NaN
