<?php
session_start();

//注销登录
if(@($_GET['action'] == "logout")){
    unset($_SESSION['userid']);
    unset($_SESSION['username']);
    echo '注销登录成功！点击此处 <a href="login.html">登录</a>';
    exit;
}

//登录
if(!isset($_POST['submit'])){
    exit('非法访问!');
}
$username = ($_POST['username']);//htmlspecialchars
$password = ($_POST['password']);

//包含数据库连接文件
//include('conn.php');
/*****************************
*数据库连接
*****************************/
/*
$conn = @mysql_connect("10.4.8.232","root","");
if (!$conn){
    die("连接数据库失败：" . mysql_error());
}
mysql_select_db("test", $conn);
//字符转换，读库
mysql_query("set character set 'gbk'");
//写库
mysql_query("set names 'gbk'");*/

$link = mysqli_connect("10.4.8.232", "root", "", "test");

/* check connection */
if (mysqli_connect_errno()) {
    printf("Connect failed: %s\n", mysqli_connect_error());
    exit();
}
//检测用户名及密码是否正确
$check_query = mysqli_multi_query($link,"select uid from user where username='$username' and password='$password'");

if($result = mysqli_store_result($link)){
	if($row = mysqli_fetch_row($result)) 
	{
    //登录成功
    //$_SESSION['username'] = $username;
    //$_SESSION['userid'] = $result['uid'];
$userid = $row[0];
//echo $userid,'hehehehhehehe</a><br />';
    //echo $username,' 欢迎你！进入 <a href=my.php?userid= ".$userid.">用户中心</a><br />';
	echo "<a href= 'my.php?userid= $userid'>我的可操作模块</a><br />";
    echo '点击此处 <a href="login.php?action=logout">注销</a> 登录！<br />';
    exit;
}
} else {
    exit('登录失败！点击此处 <a href="javascript:history.back(-1);">返回</a> 重试');
}
//http://www.kuqin.com/php5_doc/function.mysqli-multi-query.html
?>


