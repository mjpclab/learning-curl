# 下载文件

本节会大量使用GHFS，我们先准备好目录并启动GHFS，且通过`--archive`选项允许用户打包目录并下载：

```shell
$ mkdir /tmp/share
$ echo -n '0123456789abcdef' > /tmp/share/hex.txt
$ echo -n 'foo bar' > /tmp/share/foobar.txt
$ ghfs -l 8081 -r /tmp/share/ --archive /
```

## 指定输出位置

默认情况下，curl会把请求的响应体通过标准输出（stdout）打印到终端上。我们可以通过输出重定向或curl自身的`-o`或`--output`选项，将响应体输出到外部文件。

```shell
# 请求GHFS URL
$ curl http://localhost:8081/hex.txt > /tmp/hex1
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    16  100    16    0     0  22253      0 --:--:-- --:--:-- --:--:-- 16000

$ curl -o /tmp/hex2 http://localhost:8081/hex.txt
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    16  100    16    0     0  19047      0 --:--:-- --:--:-- --:--:-- 16000

$ cat /tmp/hex1
0123456789abcdef
$ cat /tmp/hex2
0123456789abcdef
```

当我们指定输出目标时，curl默认会在终端通过标准错误（stderr）输出当前的下载状态信息，例如进度，速率等，可以通过指定`-s`或`--silent`避免输出状态信息。

比较特别的是，`-o`只对一个URL生效。如果curl命令后提供了多个URL，第一个`-o`对应第一个URL，第二个`-o`对应第二个URL，以此类推：

```shell
$ curl -o 1.txt -o 2.txt http://localhost:8081/hex.txt http://localhost:8081/foobar.txt

$ cat 1.txt
0123456789abcdef
$ cat 2.txt
foo bar
```

## 自动从URL提取文件名

通常我们希望下载后的文件名与原始资源保持一直，通过使用`-O`或`--remote-name`选项，curl可以从URL中提取出文件名部分，把它当作下载后的本地文件名。与`-o`类似，一个`-O`也只针对一个URL。

```shell
$ cd ~

# 请求GHFS
$ curl -O http://localhost:8081/hex.txt

$ cat hex.txt
0123456789abcdef
```

`-O`选项默认将文件保存在当前目录，如需改变保存目录，可以用`--output-dir`指定：

```shell
$ rm -f ~/Downloads/hex.txt

# 请求GHFS
$ curl -O --output-dir ~/Downloads/ http://localhost:8081/hex.txt

$ cat ~/Downloads/hex.txt
0123456789abcdef
```

## 自动从`Content-Disposition`响应头提取文件名

有些资源是通过URL对应的服务器端逻辑动态生成的，无法通过URL末尾部分正确地推断出文件名，通常服务器端程序会通过`Content-Disposition`响应头给出参考文件名。让我们先试试请求GHFS打包目录到zip文件的`HEAD`调用：

```shell
$ curl -I 'http://localhost:8081/?zip'
HTTP/1.1 200 OK
Content-Disposition: attachment; filename=share.zip; filename*=UTF-8''share.zip
（略）
```

curl提供了`-J`或`--remote-header-name`来提取`Content-Disposition`响应头中的文件名，需要配合前面介绍的`-O`来将文件保存的文件系统：

```shell
$ rm -f ~/Downloads/curl_response
$ curl -O -J --output-dir ~/Downloads/ 'http://localhost:8081/?zip'
$ ls Downloads/
share.zip
```

## 保留资源的修改时间

如果获取的资源包含响应头`Last-Modified`，那么只要启用了`--remote-time`，curl可以将下载后的文件日期也改成该值。

```shell
$ curl -O http://localhost:8081/hex.txt

# 下载后的文件与原始文件日期并不相同
$ ls -l hex.txt /tmp/share/hex.txt
-rw-r--r-- 1 marjune marjune 16 Oct 13 21:09 hex.txt
-rw-r--r-- 1 marjune marjune 16 Oct 13 18:32 /tmp/share/hex.txt
```

```shell
$ rm -f hex.txt

$ curl -I http://localhost:8081/hex.txt
HTTP/1.1 200 OK
Last-Modified: Sun, 13 Oct 2024 10:32:09 GMT
#（略）

$ curl -O --remote-time http://localhost:8081/hex.txt

$ ls -l hex.txt /tmp/share/hex.txt
-rw-r--r-- 1 marjune marjune 16 Oct 13 18:32 hex.txt
-rw-r--r-- 1 marjune marjune 16 Oct 13 18:32 /tmp/share/hex.txt
```

## 避免覆盖已有文件

如想要避免覆盖现有文件，可以指定`--no-clobber`，curl会把额外的后缀添加到文件名之后。

```shell
$ echo dummy > hex.txt

$ ls hex*
hex.txt

$ curl -O --no-clobber http://localhost:8081/hex.txt

$ ls hex*
hex.txt  hex.txt.1

$ cat hex.txt.1
0123456789abcdef

$ cat hex.txt
dummy
```

## 续传文件

如果文件下载到一半中断，想要接着下载而不是从头开始，只要服务器支持范围请求，就可以利用curl的`-C`或`--continue-at`选项可以从指定的字节偏移开始下载。

```shell
$ echo -n '0123' > /tmp/hex.txt

$ cat /tmp/hex.txt
0123

$ curl -o /tmp/hex.txt -C 4 http://localhost:8081/hex.txt
** Resuming transfer from byte position 4

$ cat /tmp/hex.txt
0123456789abcdef
```

续传的本质还是利用了服务器对范围请求的支持，这次添加`-v`或`--verbose`选项打印请求与响应头：

```shell
$ echo -n '0123' > /tmp/hex.txt

$ curl -v -o /tmp/hex.txt -C 4 http://localhost:8081/hex.txt
** Resuming transfer from byte position 4
> GET /hex.txt HTTP/1.1
> Host: localhost:8081
> Range: bytes=4-
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 206 Partial Content
< Accept-Ranges: bytes
< Content-Length: 12
< Content-Range: bytes 4-15/16
< Content-Type: text/plain; charset=utf-8
< Last-Modified: Sun, 13 Oct 2024 10:32:09 GMT
< 
{ [12 bytes data]
```

使用`-C -`可以让curl根据文件大小自动给出偏移量：

```shell
$ echo -n '0123' > /tmp/hex.txt

$ curl -o /tmp/hex.txt -C - http://localhost:8081/hex.txt
** Resuming transfer from byte position 4
```
