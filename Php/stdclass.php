<?php

// stdclass 类似 Java 中的 Object , 是所有类的祖先
$obj = new stdClass();

$obj -> name = "lvsi" ;
$obj -> age = 21 ;

echo json_encode($obj) ;
