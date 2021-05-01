<?php
include('./beast.php');

$enStr = encode("你好");
echo "兽语：" . $enStr;

$deStr = decode("呜嗷嗷嗷啊嗷嗷~啊呜~啊~呜呜嗷");
echo '<br>音译：' . $deStr;