<?php


// 在执行完 composer dump 之后，会重新生成 autoload.php 文件
// 必须添加此行代码，否则不会引入composer的
require __DIR__ . '/vendor/autoload.php';


$loop = React\EventLoop\Factory::create();

$server = stream_socket_server('tcp://127.0.0.1:8082');
stream_set_blocking($server, false);

$loop->addReadStream($server, function ($server) use ($loop) {
    $conn = stream_socket_accept($server);
    $data = "HTTP/1.1 200 OK\r\nContent-Length: 3\r\n\r\nHi\n";
    $loop->addWriteStream($conn, function ($conn) use (&$data, $loop) {
        $written = fwrite($conn, $data);
        if ($written === strlen($data)) {
            fclose($conn);
            $loop->removeWriteStream($conn);
        } else {
            $data = substr($data, $written);
        }
    });
});

$loop->addPeriodicTimer(5, function () {
    $memory = memory_get_usage() / 1024;
    $formatted = number_format($memory, 3).'K';
    echo "Current memory usage: {$formatted}\n";
});

$loop->run();

