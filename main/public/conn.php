<?php
/*****************************
*���ݿ�����
*****************************/
$conn = @mysql_connect("10.4.8.232","root","");
if (!$conn){
    die("�������ݿ�ʧ�ܣ�" . mysql_error());
}
mysql_select_db("test", $conn);
//�ַ�ת��������
mysql_query("set character set 'gbk'");
//д��
mysql_query("set names 'gbk'");
?>
