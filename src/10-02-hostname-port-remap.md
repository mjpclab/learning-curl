# 主机名和端口重映射

`--resolve`可以重新映射域名地址，但无法重写目标端口。`--connect-to`可以同时重写目标主机和端口，并保持请求主机头为原始值。其参数格式为`SOURCE_HOST:SOURCE_PORT:DEST_HOST:DEST_PORT`。

假设我们想把`1.2.3.4:8888`的请求重写为`localhost:8000`，那么可以：

```shell
$ curl -v \
--connect-to 1.2.3.4:8888:localhost:8080 \
http://1.2.3.4:8888/foo/bar
* Connecting to hostname: localhost
* Connecting to port: 8080
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> GET /foo/bar HTTP/1.1
> Host: 1.2.3.4:8888
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Content-Type: text/plain
< Date: Wed, 16 Oct 2024 14:42:45 GMT
< Content-Length: 161
< 

================================
Request 33
================================

GET /foo/bar HTTP/1.1
Host: 1.2.3.4:8888
Accept: */*
User-Agent: curl/8.10.1
* Connection #0 to host localhost left intact
````
