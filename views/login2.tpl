<head>
    <title>简单登陆页面</title>
    <script>
        function r()
        {
            var username=document.getElementById("username");
            var pass=document.getElementById("password");
            if(username.value=="")
            {
                alert("请输入用户名");
                username.focus();
                return ;
            }
            else  if(pass.value=="")
            {
                alert("请输入密码");
                return;
            }
            else  if((pass.value=="123456")&&(username.value=="W"))
            {
                frm.submit();
            }
            else
            {
                alert("登录失败，请检查用户名和密码是否有误！");
                return;
            }
            return true;
        }
    </script>
</head>
<body>
<form name="frm" action="0.html" method="post">
    <table  >
        <tr align=center>
            <td>用户名:</td><td><input type="text" name="username"  id="username"></td>
        </tr>
        <tr align=center><td>密 码:</td><td><input type="password" name="password" id="password"></td></tr>
        <tr align=center><td colspan="2"><input type="button" value="登 录" onclick="r();"/>
            <input type="reset" value="重 置"/></td></tr>
    </table>
</form>
</body>
</html>