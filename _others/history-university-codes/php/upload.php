<?php
$path = "uploads/";
 
$extArr = array("jpg", "png", "gif");
 
if(isset($_POST) and $_SERVER['REQUEST_METHOD'] == "POST"){
    $name = $_FILES['photoimg']['name'];
    $size = $_FILES['photoimg']['size'];
     
    if(empty($name)){
        echo '��ѡ��Ҫ�ϴ���ͼƬ';
        exit;
    }
    $ext = extend($name);
    if(!in_array($ext,$extArr)){
        echo 'ͼƬ��ʽ����';
        exit;
    }
    if($size>(100*1024)){
        echo 'ͼƬ��С���ܳ���100KB';
        exit;
    }
    $image_name = time().rand(100,999).".".$ext;
    $tmp = $_FILES['photoimg']['tmp_name'];
    if(move_uploaded_file($tmp, $path.$image_name)){
        echo '<img src="'.$path.$image_name.'"  class="preview">';
    }else{
        echo '�ϴ������ˣ�';
    }
    exit;
}
exit;
 
 
 
 
function extend($file_name){
    $extend = pathinfo($file_name);
    $extend = strtolower($extend["extension"]);
    return $extend;
}
?>