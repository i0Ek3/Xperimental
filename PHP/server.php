// server
<?php

$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
socket_bind($socket, "127.0.0.1", 8888);
socket_listen($socket);
$conn = socket_accept($socket);
$write_buffer = "hello socket!";
socket_write($conn, $write_buffer);
socket_close($conn);

?>
