<?php
session_start();

//����Ƿ��¼����û��¼��ת���¼����
if(!isset($_SESSION['userid'])){
    header("Location:login.html");
    exit();
}

//�������ݿ������ļ�
include('conn.php');
$userid = $_GET['userid'];
$username = $_SESSION['username'];
//$user_query = mysql_query("select * from user where uid=$userid");
$user_query = mysql_query("select * from tModule where moduleid in(select moduleid from tLimit where uid = $userid)");
//$row = mysql_fetch_array($user_query);
echo '�ҵĿɲ���ģ�飺<br />';
//echo '�û�ID��',$userid,'<br />';
//echo '�û�����',$username,'<br />';
//echo '���䣺',$row['email'],'<br />';
//echo 'ע�����ڣ�',date("Y-m-d", $row['regdate']),'<br />';

while ($row = mysql_fetch_assoc($user_query)) {
    echo 'ģ����:'.$row["moduleid"];
    echo '----�� ģ������:','<a href="http://www.baidu.com">'.$row["modulename"]. '</a><br />';

}

echo '<a href="login.php?action=logout">ע��</a> ��¼<br />';
?>
