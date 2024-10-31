# 输出详细信息

使用`-v`或`--verbose`可以显示curl当前所执行的步骤，一行开头的`>`代表curl发出的请求头；`<`代表接收到的响应头；`*`代表其它额外信息。

```shell
$ curl -v http://localhost:8080/
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Content-Type: text/plain
< Date: Wed, 16 Oct 2024 12:31:04 GMT
< Content-Length: 156
< 

================================
Request 11
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
* Connection #0 to host localhost left intact
```
