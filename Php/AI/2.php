<?php

echo "Hello ! 请问你想聊什么 ？\n";

$life = true;

while( $file ){
	
	$str = fgets( STDIN );

	$res = sprintf("%s\n", $str);
	$res = str_replace('?', '!', $res);
	$res = str_replace('吗', '', $res);
	
	echo $res;
}

