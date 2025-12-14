'use strict';

var i0 = require('i0ek3_demo');
var http = require('http');

var str = "i0Ek3";

var server = http.createServer(function (request, response) {
    console.log(request.method + ': ' + request.url);
    response.writeHead(200, {'Content-Type': 'text/html'});
    response.end(i0.hi(str));
});

server.listen(9999);
console.log('Server is running at http://127.0.0.1:9999/');
