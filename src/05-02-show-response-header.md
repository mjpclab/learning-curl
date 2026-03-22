# 显示响应头

curl默认不显示服务器端返回的响应头，可以用`-i`或`--show-headers`（旧版本为`--include`）启用：

```shell
$ curl -i http://localhost:8080
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Sun, 22 Mar 2026 03:05:24 GMT
Content-Length: 156


================================
Request 10
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.14.1
```

注意回显输出的顶部包含了本次请求对应的响应头，而下方显示的是客户端请求头的回显。

另外，可以用`-D`或`--dump-header`将响应头输出到文件。

```shell
$ curl -D /tmp/response-headers.txt http://localhost:8080

================================
Request 11
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.14.1
```

```shell
$ cat /tmp/response-headers.txt
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Sun, 22 Mar 2026 03:10:46 GMT
Content-Length: 155
```

如果将输出文件指定为`-`，则响应头会从标准输出（stdout）输出，其效果与`-i`选项相同：

```shell
$ curl -D - http://localhost:8080
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Sun, 22 Mar 2026 03:15:12 GMT
Content-Length: 156


================================
Request 12
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.14.1
```

`-D`可以和`-i`一同使用，如果指定的输出文件为`-`，那么就会在终端上得到两份响应头：

```shell
$ curl -D - -i http://localhost:8080
HTTP/1.1 200 OK
HTTP/1.1 200 OK
Content-Type: text/plain
Content-Type: text/plain
Date: Sun, 22 Mar 2026 03:17:02 GMT
Date: Sun, 22 Mar 2026 03:17:02 GMT
Content-Length: 156
Content-Length: 156



================================
Request 13
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.14.1
```
