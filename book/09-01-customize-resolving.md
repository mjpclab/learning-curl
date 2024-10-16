# 自定义域名解析地址

`--resolve`选项可以自定义域名解析，就好象是`hosts`文件的另一个版本。该选项的值格式为`主机名:端口:解析地址[, ...]`。注意需要指定端口，当在一个curl命令中请求多个URL，他们主机名相同但端口不同，那么需要针对每个端口分别指定解析规则。

让我们把`localhost`指定解析为`127.0.0.2`：

```shell
$ curl -v --resolve localhost:8080:127.0.0.2 http://localhost:8080
* Added localhost:8080:127.0.0.2 to DNS cache
* Hostname localhost was found in DNS cache
*   Trying 127.0.0.2:8080...
* Connected to localhost (127.0.0.2) port 8080
* using HTTP/1.x
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.10.1
> Accept: */*
（略）
```

可以看到，自定义解析已经成功，`localhost`被解析到`127.0.0.2`，而请求中的主机头依旧和原始URL相同，即`Host: localhost:8080`。
