# 用HEAD方法请求资源

`HEAD`请求和`GET`效果相同，唯一的区别是`HEAD`请求不返回响应体。当需要检查资源（例如是否存在、字节大小）等，可通过`HEAD`方法避免传输不必要的响应体。

选项`-I`或`--head`用于发起`HEAD`请求。在客户端终端执行`HEAD`方法将显示服务器端响应头：

```shell
$ curl -I http://localhost:8080
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Fri, 11 Oct 2024 12:58:49 GMT
Content-Length: 157
```

此处看不到回显的响应体，这是`HEAD`请求定义的行为，服务器端逻辑应当遵守。此时切换到运行回显服务器的终端，可以看到服务器输出的实际请求内容：

```
================================
Request 20
================================

HEAD / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

服务器端程序应当保证`HEAD`请求和`GET`请求具有相同的输出逻辑，但可以对不需要响应体输出的`HEAD`方法进行一定的优化，例如在生成内容之前可能无法得知其长度，所以`Content-Length`头的值可能存在不确定性。
