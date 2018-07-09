<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
    <title>电子相册</title>
    <style type="text/css">
        div{
            width:800px;
            height:700px;
            background-color:beige;
            margin:40px auto;
            position:relative;
            list-style:none;
        }
        a:hover img{
            position:absolute;
            display:block;
            overflow:hidden;


        }
        a:hover img{
            top:100px;
            left:180px;
            width:600px;
            height:500px;
        }
        a{
            width:110px;
            height:80px;
            border-color:beige;
            display:block;
            overflow:hidden;
        }
    </style>
</head>


<body>
<div id="div1">
    <ul>
        <br />
        <a><img src="1.jpg" /></a><br />
        <a><img src="2.jpg" /></a><br />
        <a><img src="3.jpg" /></a><br />
        <a><img src="6.jpg" /></a><br />
        <a><img src="4.jpg" /></a><br />
        <a><img src="6.jpg" /></a><br />
    </ul>
</div>
</body>
</html>