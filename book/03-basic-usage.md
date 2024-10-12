# 基本用法

## 命令行格式

```
curl [选项[ 值]|URL ...]
```

curl命令之后可以跟随选项、选项的值或URL，都可以出现多次。

选项如有设置值，须连续书写，中间不能插入其他参数，例如`-X POST`。

选项与URL可以混排，任何不能被识别为选项或选项值的参数都会被当作URL。

以下示例演示了同时请求2个URL，且混排指定了2个选项：

- `--request`指定使用POST方法
- `--user-agent`指定了`User-Agent`请求头

这两个选项会同时作用在混排的两个URL的请求上，而不是第一个选项对第一个URL生效，第二个选项对第二个URL生效：

```shell
$ curl --request POST http://localhost:8080/foo --user-agent 'httpClient/1.0' http://localhost:8080/bar

================================
Request 2
================================

POST /foo HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: httpClient/1.0

================================
Request 3
================================

POST /bar HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: httpClient/1.0
```

除了将无法识别为选项或其值的命令行参数自动当作URL，也可以用`--url`选项显式指定URL。

```shell
$ curl --url http://localhost:8080 -X POST

================================
Request 2
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

## 防止shell展开

当URL中包含查询字符串（Query String）时，最好用引号将其包裹起来，以防止shell将其视为特殊字符，例如`?`会被识别为通配符，`&`在一些shell中表示将命令放入后台运行。不同shell在引号的语法上有所差异，在`sh`中可以用单引号或双引号，后者可以在字符串中嵌入变量表达式。

```shell
$ curl 'http://localhost:8080/foo/bar?x=1&y=2'

================================
Request 4
================================

GET /foo/bar?x=1&y=2 HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.0
```

```shell
$ A=100
$ B=200
$ curl "http://localhost:8080/foo/bar?x=$A&y=$B"

================================
Request 5
================================

GET /foo/bar?x=100&y=200 HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.0
```
