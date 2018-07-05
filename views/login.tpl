
<!DOCTYPE html>

<html>

<head>

    <title>Login page</title>

</head>

<body >

<form id="loginform">

    id:<input type="text" name="username" value="{{.Username}}" id="userinput"><br>

    pw:<input type="password" name="password" value="{{.Password}}"><br>

    <input type="button" value="submit" id="loginbtn">

</form>

<script type="text/javascript" src="/static/js/jquery-3.2.1.js"></script>

<script type="text/javascript">

    $(function(){

        $('#loginbtn').bind('click', function(){


            $.ajax({

                url: "/login",

                type: "post",

                asyns: true,

                data: $('#loginform').serialize(),

// dataType: "json",

                success: function(res){

                    console.log(res);

                    if(res){

                        var x = document.getElementById("userinput").value;
                        console.log(x);
                        window.location.href = "/v1/"+x+"/index";

                    }else{

                        window.location.href = "/";

                    }

// console.log({{.flash.success}});

// {{if .flash.success}}

// window.location.href = {{.flash.success}};

// {{end}}

                },

                error: function(XMLHttpRequest, testStatus, errorThrown){

                    console.log(XMLHttpRequest.status);

                    console.log(XMLHttpRequest.readyState);

                    console.log(testStatus);

                }

            });

        })

    });

</script>

</body>

</html>