# 跟随重定向

让我们用curl看看到底发生了什么：

```shell
$ curl -i http://localhost:8081/docs
HTTP/1.1 301 Moved Permanently
Content-Type: text/html; charset=utf-8
Location: /files/
Date: Sat, 09 Nov 2024 02:38:58 GMT
Content-Length: 42

<a href="/files/">Moved Permanently</a>.
```

客户端收到了一个状态码为`301 Moved Permanently`的重定向响应，并通过`Location`响应头指示资源所在的位置。

可以手动发起对新URL的访问，也可以使用curl的`-L`或`--location`选项自动跟随重定向：

```shell
$ curl -i -L http://localhost:8081/docs
HTTP/1.1 301 Moved Permanently
Content-Type: text/html; charset=utf-8
Location: /files/
Date: Sat, 09 Nov 2024 02:40:13 GMT
Content-Length: 42

HTTP/1.1 200 OK
Cache-Control: public, max-age=0
Content-Type: text/html; charset=utf-8
Vary: accept, accept-encoding
X-Content-Type-Options: nosniff
Date: Sat, 09 Nov 2024 02:40:13 GMT
Content-Length: 1294

（略）
```
