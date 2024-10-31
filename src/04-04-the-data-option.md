# 提交数据的选项

## 选项比较

在提交表单和上传文件章节我们用到了多个用于提交数据的选项，它们之间有一些不那么显而易见的区别，现在让我们再次进行归纳。

`--data`用于一般情况下的数据提交；`--data-urlencode`可以对键值对的值进行URL编码转义；`--data-raw`和`--data-binary`基本上都是不作转换直接将原始值提交给服务器，但`--data-binary`可以指定外部文件作为数据来源。

|选项|一次指定多个键值对|`@`引用外部文件|引用外部文件时保留换行|
|----|---------------|-------------|------------------|
|`-d`, `--data`, `--data-ascii`|是|是|否|
|`--data-urlencode`|否|是|否|
|`--data-raw`|是|否||
|`--data-binary`|是|是|是|

## 构造GET请求的查询字符串

### 构造查询字符串

实际上，我们可以用以上这些选项构造GET请求的查询字符串，只需额外指定`-G`或`--get`。

```shell
$ echo -n 'a=1&b=2' > /tmp/params.txt

$ echo -n 'Tom&Jerry' > /tmp/username.txt

$ curl \
-G \
-d 'x=3&y=4' \
-d @/tmp/params.txt \
--data-urlencode username@/tmp/username.txt \
'http://localhost:8080/?foo=bar'

================================
Request 5
================================

GET /?foo=bar&x=3&y=4&a=1&b=2&username=Tom%26Jerry HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

可以发现，数据选项指定的值会和URL上已有的查询字符串进行合并。

对于`-G`和`--data-urlencode`的组合，有一个更为便捷的选项可以使用：`--url-query`，它可以像`--data-urlencode`一样编码参数的值，也可以引用外部文件：

```shell
$ echo -n 'Alice&Bob' > /tmp/pair1.txt

$ curl \
--url-query pair1@/tmp/pair1.txt \
--url-query 'pair2=Tom&Jerry' \
'http://localhost:8080?x=1&y=2'

================================
Request 3
================================

GET /?x=1&y=2&pair1=Alice%26Bob&pair2=Tom%26Jerry HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

如要避免被编码，或者传递给`--url-query`的数据是已经过编码的，在其选项值的开头加上`+`即可：

```shell
$ curl \
--url-query +pair1@/tmp/pair1.txt \
--url-query '+pair2=Tom&Jerry' \
http://localhost:8080

================================
Request 4
================================

GET /?pair1@/tmp/pair1.txt&pair2=Tom&Jerry HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

以上写法构造的请求最终将被服务器端解析为3个参数，分别是

- 键`pair1@/tmp/pair1.txt`，值为空
- 键`pair2`，值为`Tom`
- 键`Jerry`，值为空

### 同时构造查询字符串和表单数据

我们只需同时指定`--url-query`和普通的数据提交选项（`--data`，`--data-raw`等），就可以构造出同时含有查询字符串和表单数据请求体的请求，默认使用POST方法：

```shell
$ curl --url-query x=1 --data y=2 http://localhost:8080

================================
Request 5
================================

POST /?x=1 HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 3
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.1

y=2
```

也可以通过`-X`来改变请求方法：

```shell
$ curl -X GET --url-query x=1 --data y=2 http://localhost:8080

================================
Request 6
================================

GET /?x=1 HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 3
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.1

y=2
```
