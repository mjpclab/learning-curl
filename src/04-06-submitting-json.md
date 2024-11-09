# 提交JSON数据

基于Restful API的服务通常接受来自客户端的JSON格式数据，且将执行结果用JSON格式返回。`--json`选项可以方便地调用这类API。

```shell
$ curl --json '{"foo":"bar","baz":null}' http://localhost:8080

================================
Request 4
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: application/json
Content-Length: 25
Content-Type: application/json
User-Agent: curl/8.10.0

{"foo":"bar","barz":null}
```

```shell
$ echo -n '{"x":10,"y":20}' > /tmp/data.json

$ curl --json @/tmp/data.json http://localhost:8080

================================
Request 5
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: application/json
Content-Length: 15
Content-Type: application/json
User-Agent: curl/8.10.0

{"x":10,"y":20}
```

`--json`选项其实是以下选项组合的快捷方式：
```
--data ARG
--header "Content-Type: application/json"
--header "Accept: application/json"
```

它将设置请求头`Content-Type: application/json`，表示其请求体内容为JSON格式。

同时设置请求头`Accept: application/json`表明期待从服务器端接收JSON格式的数据，如服务器端逻辑实现正确，那么应当返回JSON格式的数据。
