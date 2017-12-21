window.addEventListener('load', function(e) {
  document.querySelector('#SqRt2').innerHTML = 'incommensurability created by Ryan B ( @iryanb ) on 12.20.2017 © ';
}, false);

var a = 1;
var b = 1;
var n = 1;
var c = 2;
var d = 3;
var start = Date.now();
document.writeln("<br /> if a = 1 & b = 1 & c = a+b & d = 2a+b then the Square Root of 2 = d / c as it is close to 1.414213562373095 after 20 iterations until reaching 32 bit infinity of scientific notation floating point precision after 806 iterations" + "<br />");

for(n = 1; n < 807; n++){
c=a+b;
d=(2 * a) + b;
sqrt=d/c;

document.writeln(" <p />"); document.writeln("for iteration " + n + " c = " + c + " d = " + d + "<p /> √ 2 is " + sqrt);
a=c;
b=d;


document.writeln(" <br />");
}
var speed = "average";
t=Math.abs(Date.now() - start);
if (t < 24) { speed = "fast";}
if (t > 50) { speed = "slow";}
document.writeln(" <br />");
document.writeln("Compute Time = " + t + " (ms) . " + speed);




