<?php
 
    session_start();
 
    $check_code = "" ;                                          // ���浽SESSION�е���֤��
 
    $image = imagecreatetruecolor(100, 30);                     // ���ɵ�ͼ , ȫ�������� $image �˵�ͼ��
 
    $bgcolor = imagecolorallocate($image, 240, 240, 240 ) ;     // ���ɸı��ͼ����ɫ
     
    imagefill( $image, 0, 0, $bgcolor ) ;                       //  ����ɫ��䵽��ͼ��
     
     
//  
//  /* ----------������֤��(������)������䵽��ͼ���� -----------*/
//  
//  
//  $str_length = 4 ;       // ��֤�����ֵĳ��� 
//  
//  $str_font_size  = 6 ;       // ��֤�����ֵ������С               
//  
//  
//  for( $i = 0 ; $i <= $str_length -1 ; ++$i ){
//      
//      $str_font_color = imagecolorallocate($image, rand(0, 120), rand(0,120), rand(0,120) ) ;     //��֤�����ֵ���ɫ (�������ɫ)
//      $str_font_content = rand( 0 , 9 ) ;                                                         // ����һ����֤������
//      
//      $x =  $i*25 + rand(5, 10) ;                                                                             //  ��֤��������д����ͼʱ�� x ����
//      $y = rand( 5,10 ) ;                                                                                     //  ��֤��������д����ͼʱ�� y ����
//      
//      $check_code = $check_code . $str_font_content ;
//
//      imagestring( $image , $str_font_size, $x , $y, $str_font_content , $str_font_color ) ;
//      
//  }                                                           
//  /* ----------������֤��(������)������䵽��ͼ���� -----------*/  
//  
     
     
     
     
     
     
     
     
    /* ----------������֤��(���ֺ���ĸ),����䵽��ͼ�� -----------*/
     
    $str_length = 4 ;       // ��֤�����ֵĳ��� 
     
    $str_font_size  = 6 ;       // ��֤�����ֵ������С     
     
    $str_font_content_store = "3456789abcdefghijkmnopqrstuvwxy" ;   // ��֤��  ���� + ��ĸ  �ֿ�  
     
     
    for( $i = 0 ; $i<=$str_length - 1 ; ++$i ){
         
        $str_font_color = imagecolorallocate( $image , rand(0,120) , rand(0,120) , rand(0,120) ) ;
         
//      $str_font_content = substr( $str_font_content_store, rand( 0 , strlen($str_font_content_store) ) , 1 ) ;    // ������ȡ�����ַ����������������ַ���
         
        do{
         
             $str_font_content =  substr( $str_font_content_store, rand( 0 , strlen($str_font_content_store) ) , 1 ) ;
         
        }while( $str_font_content == '') ;  //  ������ȡ�����ַ�
 
        $x =  $i*25 + rand(5, 10) ;
                                                                                        //  ��֤��������д����ͼʱ�� x ����
        $y = rand( 5,10 ) ; 
 
        $check_code = $check_code . $str_font_content ;
         
        imagestring( $image , $str_font_size , $x , $y , $str_font_content , $str_font_color ) ;
         
    }
      
    /* ----------������֤��(���ֺ���ĸ),����䵽��ͼ�� -----------*/
     
     
     
    /* ----------���ɸ���Ԫ�أ�ѩ�� / �� �� ----------- */
     
    $snow_count = 200       ;                                                // ���ŵ㣨ѩ����������
     
    for( $i = 0 ; $i<= $snow_count - 1 ; ++$i ){
         
        $snow_color = imagecolorallocate( $image , rand(50,200), rand(50,200) , rand(50,200) )  ; // ���ŵ�(ѩ��) ����ɫҪ����ǳ
     
        imagesetpixel( $image , rand(1,99) , rand(1,29) , $snow_color ) ;                       //�����ŵ㣨ѩ��������ͼ����
     
    }
 
    /* ----------���ɸ���Ԫ�أ�ѩ�� / �� �� ----------- */
     
     
     
    /* ----------���ɸ����� ------------------------- */
     
    $line_count = 3        ;                                                    // �����ߵ�����
     
    for( $i = 0 ; $i<=$line_count ; ++$i ){
         
        $line_color = imagecolorallocate( $image , rand(80,220) , rand(80,220) , rand(80,220) ) ;   // �����ߵ���ɫ
        imageline( $image , rand(1,99) , rand(1,29) ,  rand(1,99) , rand(1,29) , $line_color )  ;   // ��Ӹ����ߵ���ͼ��
    }
     
    /* ----------���ɸ����� ------------------------- */
     
     
     
     
     
     
    /*------------������֤�� -------------------*/
     
    $_SESSION[ session_id() ] = $check_code ;       
     
    /*------------������֤�� -------------------*/
     
     
     
     
     
     
     
     
     
    header( "content-type:image/png " );                        //  �������ݵ�����
 
    imagepng($image);                                           //  ���ͼƬ
 
     
    imagedestroy($image);                                       //  ����ͼƬ
 
 
 
?>