'use strict';

var tool = require('itool.js');
var http = require('http');

var str = "i0Ek3";

var server = http.createServer(function (request, response) {
    console.log(request.method + ': ' + request.url);
    response.writeHead(200, {'Content-Type': 'text/html'});
    response.end(tool.isPalindrome(str, 0, 4));
    response.end(tool.isPalindrome("ollo", 0, 3));
});

server.listen(9999);
console.log('Server is running at http://127.0.0.1:9999/');
