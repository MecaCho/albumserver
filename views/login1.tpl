<!DOCTYPE html>
<html>
<head>
    <script>
        function checkForm()
        {
            alert("表单已提交！");
        }
    </script>
</head>
<body>

<form action="/demo/demo_form.asp" onsubmit="checkForm()">
    用户名：<input type="text" name="username" value="{{.Username}}"><br>
    密码：<input type="password" name="password" value="{{.Password}}"><br>
    <input type="submit" value="提交">

    <p>用户登录</p>
</form>
</body>
</html>