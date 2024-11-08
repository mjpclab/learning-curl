# 保持认证信息

在跟随重定向后，认证信息`Authorization`请求头默认不会发送给非本域主机，因为它们默认是不受信任的，而本域主机默认受信任。

```shell
$ curl -v -L -u foo:bar http://localhost:8081/docs
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* Connected to localhost (::1) port 8081
* using HTTP/1.x
* Server auth using Basic with user 'foo'
> GET /docs HTTP/1.1
> Host: localhost:8081
> Authorization: Basic Zm9vOmJhcg==		# <==== 原始请求附带认证信息
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: /files/
< Date: Sat, 09 Nov 2024 02:44:43 GMT
< Content-Length: 42
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://foo:bar@localhost:8081/files/'
* Re-using existing connection with host localhost
* Server auth using Basic with user 'foo'
> GET /files/ HTTP/1.1
> Host: localhost:8081
> Authorization: Basic Zm9vOmJhcg==		# <==== 本域跟随请求附带认证信息
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Cache-Control: public, max-age=0
< Content-Type: text/html; charset=utf-8
< Vary: accept, accept-encoding
< X-Content-Type-Options: nosniff
< Date: Sat, 09 Nov 2024 02:44:43 GMT
< Content-Length: 1294
< 
（略）
```

```shell
$ curl -v -L -u foo:bar http://localhost:8081/archive
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* Connected to localhost (::1) port 8081
* using HTTP/1.x
* Server auth using Basic with user 'foo'
> GET /archive HTTP/1.1
> Host: localhost:8081
> Authorization: Basic Zm9vOmJhcg==		# <==== 原始请求附带认证信息
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: http://127.0.0.31:8081/files/
< Date: Sat, 09 Nov 2024 02:45:24 GMT
< Content-Length: 64
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://127.0.0.31:8081/files/'
*   Trying 127.0.0.31:8081...
* Connected to 127.0.0.31 (127.0.0.31) port 8081
* using HTTP/1.x
> GET /files/ HTTP/1.1
> Host: 127.0.0.31:8081		# <==== 非本域请求没有附带认证信息
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Cache-Control: public, max-age=0
< Content-Type: text/html; charset=utf-8
< Vary: accept, accept-encoding
< X-Content-Type-Options: nosniff
< Date: Sat, 09 Nov 2024 02:45:24 GMT
< Content-Length: 1294
< 
（略）
```

选项`--location-trusted`告知curl，重定向后的URL可以被信任，因而可以带上认证信息：

```shell
$ curl -v -L --location-trusted -u foo:bar http://localhost:8081/archive
*   Trying 127.0.0.1:8081...
* Connected to localhost (127.0.0.1) port 8081 (#0)
* Server auth using Basic with user 'foo'
> GET /archive HTTP/1.1
> Host: localhost:8081
> Authorization: Basic Zm9vOmJhcg==		# <==== 原始请求附带认证信息
> User-Agent: curl/7.88.1
> Accept: */*
>
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: http://127.0.0.31:8081/files/
< Date: Fri, 08 Nov 2024 17:19:41 GMT
< Content-Length: 64
<
* Ignoring the response-body
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://127.0.0.31:8081/files/'
*   Trying 127.0.0.31:8081...
* Connected to 127.0.0.31 (127.0.0.31) port 8081 (#1)
* Server auth using Basic with user 'foo'
> GET /files/ HTTP/1.1
> Host: 127.0.0.31:8081
> Authorization: Basic Zm9vOmJhcg==		# <==== 非本域请求附带认证信息
> User-Agent: curl/7.88.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Cache-Control: public, max-age=0
< Content-Type: text/html; charset=utf-8
< Vary: accept, accept-encoding
< X-Content-Type-Options: nosniff
< Date: Fri, 08 Nov 2024 17:19:41 GMT
< Content-Length: 1294
<
（略）
```
