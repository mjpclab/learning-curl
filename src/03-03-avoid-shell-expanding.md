# 防止shell展开

当URL中包含查询字符串（Query String）时，最好用引号将其包裹起来，以防止shell将其视为特殊字符，例如`?`会被识别为通配符，`&`在一些shell中表示将命令放入后台运行。不同shell在引号的语法上有所差异，在`sh`中可以用单引号或双引号，后者可以在字符串中嵌入变量表达式。

```shell
$ curl 'http://localhost:8080/foo/bar?x=1&y=2'

================================
Request 5
================================

GET /foo/bar?x=1&y=2 HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

```shell
$ A=100
$ B=200
$ curl "http://localhost:8080/foo/bar?x=$A&y=$B"

================================
Request 6
================================

GET /foo/bar?x=100&y=200 HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```
