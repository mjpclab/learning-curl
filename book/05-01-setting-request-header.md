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

## 快捷设置选项

除了通用选项，curl还为常用请求头提供了快捷设置选项，使用时只需指定请求头的值，而无需显式指明请求头名称。

### 设置用户代理（User Agent）
-A, --user-agent

### 设置引用来源（Referrer）
-e, --referer
