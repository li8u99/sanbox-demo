<?php
$text = $_POST['info'];
if($text != null){
    $text = $_POST['info'];
    $file = fopen('info.txt', "a");
    fwrite($file, $text);
    fclose($file);
    echo "ok";
}else{
    echo "HelloWorld";
}
?>