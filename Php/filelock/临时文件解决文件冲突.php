<?php

$dir_fileopen='tmp';
function randomid(){
    return time().substr(md5(microtime()),0,rand(5,12));
}
function cfopen($filename,$mode){
    global $dir_fileopen;
    clearstatcache();
    do{
  $id=md5(randomid(rand(),TRUE));
        $tempfilename=$dir_fileopen.'/'.$id.md5($filename);
    } while(file_exists($tempfilename));
    if(file_exists($filename)){
        $newfile=false;
        copy($filename,$tempfilename);
    }else{
        $newfile=true;
    }
    $fp=fopen($tempfilename,$mode);
    return $fp?array($fp,$filename,$id,@filemtime($filename)):false;
}
function cfwrite($fp,$string){
 return fwrite($fp[0],$string);
}
function cfclose($fp,$debug='off'){
    global $dir_fileopen;
    $success=fclose($fp[0]);
    clearstatcache();
    $tempfilename=$dir_fileopen.'/'.$fp[2].md5($fp[1]);
    if((@filemtime($fp[1])==$fp[3])||($fp[4]==true&&!file_exists($fp[1]))||$fp[5]==true){
        rename($tempfilename,$fp[1]);
    }else{
        unlink($tempfilename);
  //说明有其它进程 在操作目标文件，当前进程被拒绝
        $success=false;
    }
    return $success;
}
$fp=cfopen('lock.txt','a+');
cfwrite($fp,"welcome to beijing.n");
fclose($fp,'on');