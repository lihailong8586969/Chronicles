<?php


$arr = [
	
	[ "id"=> 1,  "pid" => 0, "name" => "中国" ],
	[ "id"=> 2,  "pid" => 0, "name" => "美国" ],
	
	
	[ "id"=> 3,  "pid" => 1, "name" => "山东省" ],
	[ "id"=> 4,  "pid" => 1, "name" => "河南省" ],
	
	
	[ "id"=> 5,  "pid" => 3, "name" => "济南市" ],
	[ "id"=> 6,  "pid" => 3, "name" => "青岛市" ],
	
	[ "id"=> 7,  "pid" => 4, "name" => "郑州市" ],
	[ "id"=> 8,  "pid" => 4, "name" => "洛阳市" ],
	
	[ "id"=> 9,  "pid" => 2, "name" => "纽约" ],
	[ "id"=> 10,  "pid" => 2, "name" => "洛杉矶" ],

];

function traverse($arr, $pid, $level){

	foreach( $arr as $node ){
		
		if($node['pid'] === $pid){
			
			$node['name'] = str_repeat("-",$level) . $node['name'] ;
			
			echo $node['name']."\n";
			
			traverse($arr, $node['id'], $level+1);
		}

	}
	
}

traverse($arr, 0, 0);




