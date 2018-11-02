<?php

if($fp=fopen($fileName,'a')){
 $startTime=microtime();
 do{
  $canWrite=flock($fp,LOCK_EX);
  if(!$canWrite){
   usleep(round(rand(0,100)*1000));
  }
 }while((!$canWrite)&&((microtime()-$startTime)<1000));
 if($canWrite){
  fwrite($fp,$dataToSave);
 }
 fclose($fp);
}
