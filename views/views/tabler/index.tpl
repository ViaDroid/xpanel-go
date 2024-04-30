<!DOCTYPE HTML>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="robots" content="noindex">
    <title>{{index .config "appName"}}</title>
    {{if .isLogin}}
        <script>
            window.location.href = "/user"
        </script>
    {{else}}
        <script>
            window.location.href = "/auth/login"
        </script>
    {{end}}
</head>
</html>
