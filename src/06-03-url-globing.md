# 多URL通配符

当需要一次请求多个有一定规则的URL地址时，可以使用curl支持的通配符，curl会把它们展开成多个URL。配合下载相关的选项如`-O`、`--output-dir`等，可以批量下载资源到本地。

## 列表通配符

列表通配符是一系列枚举值的集合，被包裹在`{}`中，值之间用`,`分割，即`{value,value,...}`的形式。

```shell
$ curl http://localhost:8080/book/{curl,shell}.pdf

================================
Request 1
================================

GET /book/curl.pdf HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 2
================================

GET /book/shell.pdf HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

## 序列通配符

可将英文字母及数字用于序列通配符，用于指定值的范围，被包裹在`[]`中，可选指定的步进值，格式为`[start-end]`或`[start-end:step]`。

```shell
$ curl http://localhost:8080/dictionary/[a-c]/words

================================
Request 3
================================

GET /dictionary/a/words HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 4
================================

GET /dictionary/b/words HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 5
================================

GET /dictionary/c/words HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

```shell
$ curl http://localhost:8080/magzine/[2021-2024]/index.html

================================
Request 6
================================

GET /magzine/2021/index.html HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 7
================================

GET /magzine/2022/index.html HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 8
================================

GET /magzine/2023/index.html HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 9
================================

GET /magzine/2024/index.html HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

```shell
$ curl http://localhost:8080/magzine/2024/[1-5:2]/index.html

================================
Request 10
================================

GET /magzine/2024/1/index.html HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 11
================================

GET /magzine/2024/3/index.html HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 12
================================

GET /magzine/2024/5/index.html HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

数字值还可以有多个前导`0`指示最小位数：

```shell
$ curl http://localhost:8080/archive/[008-012].zip

================================
Request 13
================================

GET /archive/008.zip HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 14
================================

GET /archive/009.zip HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 15
================================

GET /archive/010.zip HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 16
================================

GET /archive/011.zip HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 17
================================

GET /archive/012.zip HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

## 递归展开

组合使用多个通配符，curl会以笛卡尔乘积的形式展开成多个URL：

```shell
$ curl http://localhost:8080/videos/{curl,shell}/chapter[01-05:2].mp4

================================
Request 26
================================

GET /videos/curl/chapter01.mp4 HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 27
================================

GET /videos/curl/chapter03.mp4 HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 28
================================

GET /videos/curl/chapter05.mp4 HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 29
================================

GET /videos/shell/chapter01.mp4 HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 30
================================

GET /videos/shell/chapter03.mp4 HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

================================
Request 31
================================

GET /videos/shell/chapter05.mp4 HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

## 避免shell解析

由于`{}`和`[]`也是shell模式匹配的语法，如果本地文件系统刚好有和URL模式匹配的文件路径，会被shell执行展开匹配，该匹配逻辑与curl并不相同。为避免发生此类问题，最好为URL包裹引号。
