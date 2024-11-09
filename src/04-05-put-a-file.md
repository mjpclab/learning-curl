# 使用PUT方法上传资源

一些服务器支持以`PUT`方法上传单个资源到目标路径，请求体的内容即为资源内容。可以用`-T`或`--upload-file`指定本地文件用于上传：

```shell
$ echo -n 'hello world' > /tmp/greeting.txt
$ curl -T /tmp/greeting.txt http://localhost:8080

================================
Request 7
================================

PUT /greeting.txt HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 11
User-Agent: curl/8.10.1

hello world
```

可以看出，`-T`其实是一种快捷方式，我们可以用`-X`和`--data-binary`的组合实现同样的请求：

```shell
$ curl -X PUT --data-binary @/tmp/greeting.txt http://localhost:8080

================================
Request 9
================================

PUT / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 11
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.1

hello world
```
