<?php

error_reporting(E_ALL & ~E_NOTICE & ~E_WARNING);
error_reporting(2047);
ini_set("display_errors", 1);
ini_set("error_prepend_string", "<div class='alert alert-error alert-block'>");
ini_set("error_append_string", "</div>");
ini_set("html_errors", 1);
ini_set("error_log", "/var/www/vhosts/ixxocart.net/evolution/errors.log");

echo phpinfo();
?>