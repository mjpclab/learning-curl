# 范围请求

范围请求（Range Request）提供了一种只请求资源主体一部分片段的能力，通常用于`
GET`请求。

## 服务器端支持

当用`HEAD`或`GET`请求服务器端某个URL时，如果服务器支持范围请求，那么在响应头中会包含`Accept-Ranges: bytes`，以此告知客户端服务器支持范围请求。

我们先看下不支持范围请求的回显服务器的输出内容，用`-i`指示curl输出服务器端响应头：

```shell
$ curl -i http://localhost:8080
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Sat, 12 Oct 2024 13:22:31 GMT
Content-Length: 155


================================
Request 7
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

```shell
$ curl -i -I http://localhost:8080
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Sat, 12 Oct 2024 13:22:36 GMT
Content-Length: 156
```

现在我们创建一个共享目录，并用GHFS共享出该目录：

```shell
$ mkdir /tmp/share
$ echo -n '0123456789abcdef' > /tmp/share/hex.txt
$ ghfs -l 8081 -r /tmp/share/
```

让我们测试一下GHFS是否支持范围请求：

```shell
$ curl -i http://localhost:8081/hex.txt
HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 16
（略）

0123456789abcdef
```

响应头中的`Accept-Ranges: bytes`说明GHFS支持对该URL进行范围请求。

## 范围请求格式

curl的`-r`或`--range`选项用于发起范围请求，先让我们用回显服务器看一下请求格式：

```shell
$ curl -r 0-3 http://localhost:8080/

================================
Request 12
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
Range: bytes=0-3
User-Agent: curl/8.10.1
```

范围请求使用请求头`Range: bytes=RANGE`来表示。我们请求了`/`这个资源的0～3这个范围，也就是起始的4个字节，范围的第一个字节地址是0,也可以认为是偏移量。

如果想要请求从偏移量为n的字节开始到末尾，可以写成`n-`：

```shell
$ curl -r 4- http://localhost:8080/

================================
Request 13
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
Range: bytes=4-
User-Agent: curl/8.10.1
```

如果要请求最末尾的n个字节，可以写成`-n`：

```shell
$ curl -r -4 http://localhost:8080/

================================
Request 15
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
Range: bytes=-4
User-Agent: curl/8.10.1
```

## 范围请求结果

当服务器接受客户端对指定URL的范围请求时，其会输出状态码`206 Partial Content`，而不是`200 OK`。如果服务器拒绝输出范围，或者不支持范围请求，那么会像往常一样输出`200`状态码以及URL资源的完整内容。

现在，我们对刚才启动的GHFS发起范围请求，看看其如何响应。

```shell
$ curl -i -r 0-4 http://localhost:8081/hex.txt
HTTP/1.1 206 Partial Content
Accept-Ranges: bytes
Content-Length: 5
Content-Range: bytes 0-4/16
Content-Type: text/plain; charset=utf-8
（略）

01234
```

```shell
$ curl -i -r 8- http://localhost:8081/hex.txt
HTTP/1.1 206 Partial Content
Accept-Ranges: bytes
Content-Length: 8
Content-Range: bytes 8-15/16
Content-Type: text/plain; charset=utf-8
（略）

89abcdef
```

```shell
$ curl -i -r -3 http://localhost:8081/hex.txt
HTTP/1.1 206 Partial Content
Accept-Ranges: bytes
Content-Length: 3
Content-Range: bytes 13-15/16
Content-Type: text/plain; charset=utf-8
（略）

def
```

响应头中的`Content-Range: bytes start-end/total`指示了本响应输出对应的字节范围，由于范围地址从0开始，因此`end`的最大值为`total-1`。

当请求的范围超出资源的实际大小时，只要两者有重合部分，还是会正常输出，否则服务器端会报错，返回状态码`416 Requested Range Not Satisfiable`。

```shell
$ curl -i -r 12-100 http://localhost:8081/hex.txt
HTTP/1.1 206 Partial Content
Accept-Ranges: bytes
Content-Length: 4
Content-Range: bytes 12-15/16
Content-Type: text/plain; charset=utf-8
（略）

cdef
```

```shell
$ curl -i -r 100-200 http://localhost:8081/hex.txt
HTTP/1.1 416 Requested Range Not Satisfiable
Content-Range: bytes */16
Content-Type: text/plain; charset=utf-8
Content-Length: 33
（略）

invalid range: failed to overlap
```

现在，让我们试一下，如果在一个请求中指定了多个范围会发生什么：

```shell
$ curl -i -r 0-3,8-11,14- http://localhost:8081/hex.txt
HTTP/1.1 206 Partial Content
Accept-Ranges: bytes
Content-Length: 493
Content-Type: multipart/byteranges; boundary=68139b69c478491c38a766657527deba0052d7b57fcd50c315e6e2a9a5a1
（略）

--68139b69c478491c38a766657527deba0052d7b57fcd50c315e6e2a9a5a1
Content-Range: bytes 0-3/16
Content-Type: text/plain; charset=utf-8

0123
--68139b69c478491c38a766657527deba0052d7b57fcd50c315e6e2a9a5a1
Content-Range: bytes 8-11/16
Content-Type: text/plain; charset=utf-8

89ab
--68139b69c478491c38a766657527deba0052d7b57fcd50c315e6e2a9a5a1
Content-Range: bytes 14-15/16
Content-Type: text/plain; charset=utf-8

ef
--68139b69c478491c38a766657527deba0052d7b57fcd50c315e6e2a9a5a1--
```

服务器端的输出格式居然变了！其响应头`Content-Type: multipart/byteranges`表示响应体包含多个part，这与我们在上传文件章节中看到的结构非常相似，区别有两点：

一是上传文件时，传输方向为客户端向服务器端传送多个part，而范围响应的多个part则由服务器端传输给客户端。

二是其`Content-Type`头的值为`multipart/byteranges`，而上传文件时用的是`multipart/form-data`。
