<?php
 
    session_start();
 
    $check_code = "" ;                                          // 保存到SESSION中的验证码
 
    $image = imagecreatetruecolor(100, 30);                     // 生成底图 , 全程作用在 $image 此底图上
 
    $bgcolor = imagecolorallocate($image, 240, 240, 240 ) ;     // 生成改变底图的颜色
     
    imagefill( $image, 0, 0, $bgcolor ) ;                       //  把颜色填充到底图上
     
     
//  
//  /* ----------生成验证码(纯数字)，并填充到底图上面 -----------*/
//  
//  
//  $str_length = 4 ;       // 验证码数字的长度 
//  
//  $str_font_size  = 6 ;       // 验证码数字的字体大小               
//  
//  
//  for( $i = 0 ; $i <= $str_length -1 ; ++$i ){
//      
//      $str_font_color = imagecolorallocate($image, rand(0, 120), rand(0,120), rand(0,120) ) ;     //验证码数字的颜色 (较深的颜色)
//      $str_font_content = rand( 0 , 9 ) ;                                                         // 生成一个验证码数字
//      
//      $x =  $i*25 + rand(5, 10) ;                                                                             //  验证码数字填写到底图时的 x 坐标
//      $y = rand( 5,10 ) ;                                                                                     //  验证码数字填写到底图时的 y 坐标
//      
//      $check_code = $check_code . $str_font_content ;
//
//      imagestring( $image , $str_font_size, $x , $y, $str_font_content , $str_font_color ) ;
//      
//  }                                                           
//  /* ----------生成验证码(纯数字)，并填充到底图上面 -----------*/  
//  
     
     
     
     
     
     
     
     
    /* ----------生成验证码(数字和字母),并填充到底图上 -----------*/
     
    $str_length = 4 ;       // 验证码数字的长度 
     
    $str_font_size  = 6 ;       // 验证码数字的字体大小     
     
    $str_font_content_store = "3456789abcdefghijkmnopqrstuvwxy" ;   // 验证码  数字 + 字母  仓库  
     
     
    for( $i = 0 ; $i<=$str_length - 1 ; ++$i ){
         
        $str_font_color = imagecolorallocate( $image , rand(0,120) , rand(0,120) , rand(0,120) ) ;
         
//      $str_font_content = substr( $str_font_content_store, rand( 0 , strlen($str_font_content_store) ) , 1 ) ;    // 可能提取出空字符，所以用下面这种方法
         
        do{
         
             $str_font_content =  substr( $str_font_content_store, rand( 0 , strlen($str_font_content_store) ) , 1 ) ;
         
        }while( $str_font_content == '') ;  //  避免提取到空字符
 
        $x =  $i*25 + rand(5, 10) ;
                                                                                        //  验证码数字填写到底图时的 x 坐标
        $y = rand( 5,10 ) ; 
 
        $check_code = $check_code . $str_font_content ;
         
        imagestring( $image , $str_font_size , $x , $y , $str_font_content , $str_font_color ) ;
         
    }
      
    /* ----------生成验证码(数字和字母),并填充到底图上 -----------*/
     
     
     
    /* ----------生成干扰元素（雪花 / 线 ） ----------- */
     
    $snow_count = 200       ;                                                // 干扰点（雪花）的数量
     
    for( $i = 0 ; $i<= $snow_count - 1 ; ++$i ){
         
        $snow_color = imagecolorallocate( $image , rand(50,200), rand(50,200) , rand(50,200) )  ; // 干扰点(雪花) 的颜色要尽量浅
     
        imagesetpixel( $image , rand(1,99) , rand(1,29) , $snow_color ) ;                       //添点干扰点（雪花）到底图上面
     
    }
 
    /* ----------生成干扰元素（雪花 / 线 ） ----------- */
     
     
     
    /* ----------生成干扰线 ------------------------- */
     
    $line_count = 3        ;                                                    // 干扰线的数量
     
    for( $i = 0 ; $i<=$line_count ; ++$i ){
         
        $line_color = imagecolorallocate( $image , rand(80,220) , rand(80,220) , rand(80,220) ) ;   // 干扰线的颜色
        imageline( $image , rand(1,99) , rand(1,29) ,  rand(1,99) , rand(1,29) , $line_color )  ;   // 添加干扰线到底图上
    }
     
    /* ----------生成干扰线 ------------------------- */
     
     
     
     
     
     
    /*------------保存验证码 -------------------*/
     
    $_SESSION[ session_id() ] = $check_code ;       
     
    /*------------保存验证码 -------------------*/
     
     
     
     
     
     
     
     
     
    header( "content-type:image/png " );                        //  声明内容的类型
 
    imagepng($image);                                           //  输出图片
 
     
    imagedestroy($image);                                       //  销毁图片
 
 
 
?>