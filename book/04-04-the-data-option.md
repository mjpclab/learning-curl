# 提交数据的选项

## 选项比较

在提交表单和上传文件章节我们用到了多个用于提交数据的选项，它们之间有一些不那么显而易见的区别，现在让我们再次进行归纳。

|选项|一次指定多个键值对|`@`引用外部文件|引用外部文件时保留换行|
|----|---------------|-------------|------------------|
|`-d`, `--data`, `--data-ascii`|是|是|否|
|`--data-raw`|是|否||
|`--data-urlencode`|否|是|否|
|`--data-binary`|是|是|是|

## 构造GET请求的查询字符串

实际上，我们可以用以上这些选项构造GET请求的查询字符串，只需额外指定`-G`或`--get`。

```shell
$ echo -n 'a=1&b=2' > /tmp/params.txt

$ echo -n 'Tom&Jerry' > /tmp/username.txt

$ curl -G -d 'x=3&y=4' -d @/tmp/params.txt --data-urlencode username@/tmp/username.txt http://localhost:8080/?foo=bar

================================
Request 5
================================

GET /?foo=bar&x=3&y=4&a=1&b=2&username=Tom%26Jerry HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
```

可以发现，数据选项指定的值会和URL上已有的查询字符串进行合并。
