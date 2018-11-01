<?php



function queue($count){

    for ($i = 0; $i <= $count-1; ++$i) {

        yield $i;
    }
}

foreach ( queue(1000) as $item => $value){

// echo md5( time()."{$item}" .rand(-9999999,9999999) ) . "<br>" ;

// $micro_time = explode(".",microtime());

// echo "20150123 ".rand(100000,999999)." ".time(). " " . intval($micro_time[1]) ." ".rand(2000,9999)." ".$item."<br>";

// echo date("YmdHis") . rand(10000001,99999999). $item ."<br>";

    echo date("YmdHis") . " " .uniqid() . mt_rand(10000001,99999999) . " " . $value . "<br>";
}







