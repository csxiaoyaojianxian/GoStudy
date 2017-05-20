<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }
    header {
      padding: 100px 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 50px 0;
      color: #999;
    }
    a {
      color: #444;
      text-decoration: none;
    }
  </style>
</head>

<body>
  <header>
    <div>
      分支结构
      {{if .TrueCond}}
      true condition
      {{end}}

      {{if .FalseCond}}
      {{else}}
      false condition
      {{end}}

    </div>

    <div>
      非嵌套方式
      {{.User.Name}};{{.User.Age}};{{.User.Sex}}

      嵌套方式
      {{with .User}}
      {{.Name}};{{.Age}};{{.Sex}}
      {{end}}
    </div>
    
    <div>
      直接打印
      {{.Nums}}
      循环打印
      {{range .Nums}}
      {{.}}
      {{end}}
    </div>

    <div>
      模板变量
      {{$tplVar := .TplVar}}
      {{$tplVar}}
    </div>

    <div>
      模板函数
      {{str2html .Html}}
      字符转义
      {{.Pipe | htmlquote}}
    </div>

    <div>
      模板嵌套
      {{template "test"}}
    </div>
  </header>
  <footer>
    <div>
      Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      <a href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  
  <script src="/static/js/reload.min.js"></script>
</body>
</html>

{{define "test"}}
<div>
  this is test template  
</div>
{{end}}