<?php

/**
 * Syc <github.com/SycAlright>
 * Beast_SDK PHP(5.4+)
 */

$beast = ['嗷', '呜', '啊', '~'];

function encode($str)
{
	$beast = $GLOBALS['beast'];
	$code = null;
	$hexArray = str_split_unicode(bin2hex($str));
	foreach ($hexArray as $k => $v) {
		$x = base_convert($v, 16, 10) + $k % 16;
		if ($x >= 16) {
			$x -= 16;
		}
		$code .= $beast[($x / 4)] . $beast[$x % 4];
	}
	return $code;
}

function decode($str)
{
	$beast = $GLOBALS['beast'];
	$code = null;
	$hexArray = str_split_unicode($str);
	$n = count($hexArray);
	for ($i = 0; $i < $n; $i++) {
		if ($i % 2 == 0) {
			if (empty($hexArray[$i + 1])) {
				break;
			}
			$A = array_search($hexArray[$i], $beast);
			$B = array_search($hexArray[$i + 1], $beast);
			$x = (($A * 4) + $B) - (($i / 2) % 16);
			if ($x < 0) {
				$x += 16;
			}
			$code .= dechex($x);
		}
	}
	return pack("H*", $code);
}

function str_split_unicode($str, $l = 0)
{
	if ($l > 0) {
		$ret = array();
		$len = mb_strlen($str, "UTF-8");
		for ($i = 0; $i < $len; $i += $l) {
			$ret[] = mb_substr($str, $i, $l, "UTF-8");
		}
		return $ret;
	}
	return preg_split("//u", $str, -1, PREG_SPLIT_NO_EMPTY);
}


