<?php
session_start();

//检测是否登录，若没登录则转向登录界面
if(!isset($_SESSION['userid'])){
    header("Location:login.html");
    exit();
}

//包含数据库连接文件
include('conn.php');
$userid = $_GET['userid'];
$username = $_SESSION['username'];
//$user_query = mysql_query("select * from user where uid=$userid");
$user_query = mysql_query("select * from tModule where moduleid in(select moduleid from tLimit where uid = $userid)");
//$row = mysql_fetch_array($user_query);
echo '我的可操作模块：<br />';
//echo '用户ID：',$userid,'<br />';
//echo '用户名：',$username,'<br />';
//echo '邮箱：',$row['email'],'<br />';
//echo '注册日期：',date("Y-m-d", $row['regdate']),'<br />';

while ($row = mysql_fetch_assoc($user_query)) {
    echo '模块编号:'.$row["moduleid"];
    echo '----》 模块名称:','<a href="http://www.baidu.com">'.$row["modulename"]. '</a><br />';

}

echo '<a href="login.php?action=logout">注销</a> 登录<br />';
?>
