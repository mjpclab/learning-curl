curl默认不显示服务器端返回的响应头，可以用`-i`或`--show-headers`（旧版本为`--include`）启用：

```shell
$ curl -i http://localhost:8080
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Fri, 11 Oct 2024 12:41:19 GMT
Content-Length: 156


================================
Request 18
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```
