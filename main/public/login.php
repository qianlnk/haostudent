<?php
session_start();

//ע����¼
if(@($_GET['action'] == "logout")){
    unset($_SESSION['userid']);
    unset($_SESSION['username']);
    echo 'ע����¼�ɹ�������˴� <a href="login.html">��¼</a>';
    exit;
}

//��¼
if(!isset($_POST['submit'])){
    exit('�Ƿ�����!');
}
$username = ($_POST['username']);//htmlspecialchars
$password = ($_POST['password']);

//�������ݿ������ļ�
//include('conn.php');
/*****************************
*���ݿ�����
*****************************/
/*
$conn = @mysql_connect("10.4.8.232","root","");
if (!$conn){
    die("�������ݿ�ʧ�ܣ�" . mysql_error());
}
mysql_select_db("test", $conn);
//�ַ�ת��������
mysql_query("set character set 'gbk'");
//д��
mysql_query("set names 'gbk'");*/

$link = mysqli_connect("10.4.8.232", "root", "", "test");

/* check connection */
if (mysqli_connect_errno()) {
    printf("Connect failed: %s\n", mysqli_connect_error());
    exit();
}
//����û����������Ƿ���ȷ
$check_query = mysqli_multi_query($link,"select uid from user where username='$username' and password='$password'");

if($result = mysqli_store_result($link)){
	if($row = mysqli_fetch_row($result)) 
	{
    //��¼�ɹ�
    //$_SESSION['username'] = $username;
    //$_SESSION['userid'] = $result['uid'];
$userid = $row[0];
//echo $userid,'hehehehhehehe</a><br />';
    //echo $username,' ��ӭ�㣡���� <a href=my.php?userid= ".$userid.">�û�����</a><br />';
	echo "<a href= 'my.php?userid= $userid'>�ҵĿɲ���ģ��</a><br />";
    echo '����˴� <a href="login.php?action=logout">ע��</a> ��¼��<br />';
    exit;
}
} else {
    exit('��¼ʧ�ܣ�����˴� <a href="javascript:history.back(-1);">����</a> ����');
}
//http://www.kuqin.com/php5_doc/function.mysqli-multi-query.html
?>


