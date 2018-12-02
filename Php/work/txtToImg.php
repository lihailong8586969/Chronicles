<?php 
// Set the content-type 
header('Content-type: image/png'); 
 
// Create the image 
$im = imagecreatetruecolor(1200, 600); 
 
// Create some colors 
$white = imagecolorallocate($im, 255, 255, 255); 
$grey = imagecolorallocate($im, 128, 128, 128); 
$black = imagecolorallocate($im, 0, 0, 0); 
imagefilledrectangle($im, 0, 0, 1200, 600, $white); 
 
// The text to draw 
$text = '小明又出现了…… 
老师：“多位数减法，遇到低位数不够减时，就向高位数去借。” 
小明：“高位数不借怎么办？” 
老师：“你出去..！ 
 
老师讲圣经，讲到大洪水把地球上生物全淹死了。 
小明问老师：你确定？ 
老师说：确定。 
小明：那鱼呢？ 
老师：你出去！'; 
// Replace path by your own font path 
$font = 'hei.ttf'; 
 
// Add some shadow to the text 
imagettftext($im, 12, 0, 51, 101, $grey, $font, $text); 
 
// Add the text 
//imagettftext($im, 12, 0, 50, 100, $black, $font, $text); 
 
// Using imagepng() results in clearer text compared with imagejpeg() 
imagepng($im); 
imagedestroy($im); 
?> 