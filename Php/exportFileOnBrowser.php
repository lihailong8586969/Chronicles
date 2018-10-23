<?php 

$filename = "s.txt";

header();


header("Cache-Control : public");
header("Content-Description: File Transfer");
header("Content-desposition: attachment; filename=".basename($filename));
header("Content-type: application/zip"); // zip格式的
header("Content-Transfer-Encoding: binary"); // 告诉浏览器这是二进制文件
header("Content-Length: ".filesize($filename)); // 告诉浏览器文件的大小

$filename = "s.txt.new";

@readfile(filename);
