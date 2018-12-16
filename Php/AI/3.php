<?php

function ai($input){
	
	$target = ['吗', '？', '?'];
	$replace = ['', '！', '!'];

	return str_replace($target, $replace, $input);
}

while(true){

	$input = readline('Me: ');
	$output = ai($input);
	echo 'AI: ' . $output . '\n';
}
