<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>登录</title>
</head>
<body>
    <div class="">
        <form action="http://127.0.0.1:9090/login" method="POST">
        <table>
            <tr>
                <td>用户名：</td>
                <td>
                    <input type="text" name="username" />
                </td>
                <td>密码：</td>
                <td>
                    <input type="text" name="password" />
                </td>
                <td colspan="2">
                    <input type="submit" name="提交"  />
                </td>
            </tr>
        </table>
        </form>
    </div>
</body>
</html>