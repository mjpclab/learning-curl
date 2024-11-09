# 直接附带Cookie值

假设在之前的请求中服务器端已经下发了Cookie值，本次请求在传回Cookie时，可以直接在命令行上使用`-b`或`--cookie`选项给出键值对。可以一次指定多个值，之间用分号分割，也可以多次使用`-b`选项：

```shell
$ curl -b 'x=1;y=2' -b 'z=3' http://localhost:8080

================================
Request 23
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
Cookie: x=1;y=2;z=3
User-Agent: curl/8.10.1
```
