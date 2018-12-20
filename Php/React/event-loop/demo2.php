<?php

require __DIR__ . '/vendor/autoload.php';

$loop = React\EventLoop\Factory::create();

// 只是创建一个 Socket
// Create an Internet or Unix domain server socket
// Creates a stream or datagram socket on the specified local_socket.
// This function only creates a socket, to begin accepting connections use stream_socket_accept().
$server = stream_socket_server('tcp://127.0.0.1:8082');

//  为资源流设置阻塞或者阻塞模式
stream_set_blocking($server, false); // 如果 mode 为0，资源流将会被转换为非阻塞模式；如果是1，资源流将会被转换为阻塞模式
//  在非阻塞模式下，调用 fgets() 总是会立即返回。而在阻塞模式下，将会一直等到从资源流里面获取到数据才能返回。

try{

    // 学
    $loop->addReadStream($server, function ($server) use ($loop) {

        $conn = stream_socket_accept($server);

        // HTTP/1.1 200 OK\r\nContent-Length: 3\r\n\r\n文字\n
        // 注意必须以这种形式写才可以！
        $data = "HTTP/1.1 200 OK\r\nContent-Length: 3\r\n\r\nHi\n";

        $loop->addWriteStream($conn, function ($conn) use (&$data, $loop) {

            $written = fwrite($conn, $data);

            // 如果一字不差的全部写入成功
            if ($written === strlen($data)) {

                fclose($conn);
                $loop->removeWriteStream($conn);
            } else {

                $data = substr($data, $written);
            }
        });
    });

}catch(Exception $e){


}

// 每隔 $interval 秒执行一次 addPeriodicTimer 中的回调函数
// 每隔 5 秒执行一次 addPeriodicTimer 中的回调函数
$loop->addPeriodicTimer(5, function () {

    $memory = memory_get_usage() / 1024;
    $formatted = number_format($memory, 3).'K';
    echo "Current memory usage: {$formatted}\n";
});

// 开始运行
$loop->run();

