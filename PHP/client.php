// client
//

<?php

$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
$conn = socket_connect($socket, "127.0.0.1", 8888);
if ($conn) {
    while ($buff = socket_read($socket, 1024)) {
        echo 'The message is: ' . $buff . PHP_EOL;
    }
}
socket_close($socket);

?>
