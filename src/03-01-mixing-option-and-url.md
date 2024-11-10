# 选项与URL混排

选项与URL可以混排，任何不能被识别为选项或选项值的参数都会被当作URL。

以下示例演示了同时请求两个URL，且混排指定了2个选项：

- `--request`指定使用`POST`方法
- `--user-agent`指定了`User-Agent`请求头

这两个选项会同时作用在混排的两个URL的请求上，而不是第一个选项对第一个URL生效，第二个选项对第二个URL生效：

```shell
$ curl \
--request POST \
http://localhost:8080/foo \
--user-agent 'httpClient/1.0' \
http://localhost:8080/bar

================================
Request 1
================================

POST /foo HTTP/1.1          # <==== POST方法
Host: localhost:8080
Accept: */*
User-Agent: httpClient/1.0  # <==== 自定义UA

================================
Request 2
================================

POST /bar HTTP/1.1          # <==== POST方法
Host: localhost:8080
Accept: */*
User-Agent: httpClient/1.0  # <==== 自定义UA
```

要使某些选项只对其中一个URL生效，需要使用选项`-:`或`--next`将它们分割开来：

```shell
$ curl \
--request POST \
http://localhost:8080/foo \
-: \
--user-agent 'httpClient/1.0' \
http://localhost:8080/bar

================================
Request 13
================================

POST /foo HTTP/1.1      # <==== POST方法
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1 # <==== 默认UA

================================
Request 14
================================

GET /bar HTTP/1.1           # <==== 默认GET方法
Host: localhost:8080
Accept: */*
User-Agent: httpClient/1.0  # <==== 自定义UA
```

除了将无法识别为选项或其值的命令行参数自动当作URL，也可以用`--url`选项显式指定URL。

```shell
$ curl --url http://localhost:8080 --request POST

================================
Request 3
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```
