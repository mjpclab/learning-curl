# 设置请求头

## 通用设置选项

curl提供了许多选项来设置请求头，最一般的用法是通过`-H`或`--header`来设置，其值格式为`Header-Name: value`，可以出现多次。例如：

```shell
$ curl -H 'Accept: text/html' -H 'Accept-Language: zh-CN' http://localhost:8080

================================
Request 21
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: text/html
Accept-Language: zh-CN
User-Agent: curl/8.10.1
```

也可以设置为`@file_path`的格式从文件加载：

```shell
$ cat << EOF > /tmp/headers.txt
> Accept: text/html,text/xml;q=0.9
> Accept-Language: en-US
> EOF

$ curl -H @/tmp/headers.txt http://localhost:8080

================================
Request 22
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: text/html,text/xml;q=0.9
Accept-Language: en-US
User-Agent: curl/8.10.1
```

要删除某个已有的请求头，将其设为空即可，例如：

```shell
$ curl http://localhost:8080

================================
Request 11
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1

$ curl -H 'Accept:' http://localhost:8080

================================
Request 12
================================

GET / HTTP/1.1
Host: localhost:8080
User-Agent: curl/8.10.1
```

## 快捷设置选项

除了通用选项，curl还为常用请求头提供了快捷设置选项，使用时只需指定请求头的值，而无需显式指明请求头名称。

### 设置用户代理（User Agent）
`-A`或`--user-agent`用于设置用户代理（User Agnet）字符串，效果和`--header 'User-Agent: value'`相同。

```shell
$ curl -A 'mycurl/0.0.1-alpha' http://localhost:8080

================================
Request 13
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: mycurl/0.0.1-alpha
```

### 设置引用来源（Referrer）
`-e`或`--referer`用于设置来源URL,效果同`--header 'Referer: value'`：

```shell
$ curl -e http://example.com/foo/bar http://localhost:8080

================================
Request 17
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
Referer: http://example.com/foo/bar
User-Agent: curl/8.10.1
```
