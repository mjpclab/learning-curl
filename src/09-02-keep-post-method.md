# 保持POST请求

当使用诸如`-d`这样的选项间接导致curl使用POST方法时，在跟随重定向后curl默认使用`GET`方法来请求后续资源：

```shell
$ curl -d 'foo=bar' -L -v http://localhost:8081/docs
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* Connected to localhost (::1) port 8081
* using HTTP/1.x
> POST /docs HTTP/1.1		# <==== POST 请求
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> Content-Length: 7
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 7 bytes
< HTTP/1.1 301 Moved Permanently
* Need to rewind upload for next request
< Location: /files/
< Date: Sat, 09 Nov 2024 02:40:54 GMT
< Content-Length: 0
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8081/files/'
* Switch from POST to GET
* Re-using existing connection with host localhost
> GET /files/ HTTP/1.1		# <==== GET 请求
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Cache-Control: public, max-age=0
< Content-Type: text/html; charset=utf-8
< Vary: accept, accept-encoding
< X-Content-Type-Options: nosniff
< Date: Sat, 09 Nov 2024 02:40:54 GMT
< Content-Length: 1294
<
（略）
```

注意以上日志中提示了请求方法的切换：

```
* Switch from POST to GET
```

如要对后续请求继续保持`POST`方法，针对状态码`301`，`302`和`303`，可以使用选项：

- `--post301`
- `--post302`
- `--post303`

```shell
$ curl -d 'foo=bar' -L --post301 -v http://localhost:8081/docs
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* Connected to localhost (::1) port 8081
* using HTTP/1.x
> POST /docs HTTP/1.1		# <==== POST 请求
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> Content-Length: 7
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 7 bytes
< HTTP/1.1 301 Moved Permanently
* Need to rewind upload for next request
< Location: /files/
< Date: Sat, 09 Nov 2024 02:42:19 GMT
< Content-Length: 0
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8081/files/'
* Re-using existing connection with host localhost
> POST /files/ HTTP/1.1		# <==== POST 请求
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> Content-Length: 7
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 7 bytes
< HTTP/1.1 200 OK
< Cache-Control: public, max-age=0
< Content-Type: text/html; charset=utf-8
< Vary: accept, accept-encoding
< X-Content-Type-Options: nosniff
< Date: Sat, 09 Nov 2024 02:42:19 GMT
< Content-Length: 1294
< 
（略）
```

不过，对于用`-X`或`--request`显式指定了方法的请求，curl在跟随重定向后将继续保持该方法：

```shell
$ curl -X PUT -L -v http://localhost:8081/docs
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* Connected to localhost (::1) port 8081
* using HTTP/1.x
> PUT /docs HTTP/1.1		# <==== PUT 请求
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 301 Moved Permanently
< Location: /files/
< Date: Sat, 09 Nov 2024 02:43:10 GMT
< Content-Length: 0
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8081/files/'
* Re-using existing connection with host localhost
> PUT /files/ HTTP/1.1		# <==== PUT 请求
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
< Cache-Control: public, max-age=0
< Content-Type: text/html; charset=utf-8
< Vary: accept, accept-encoding
< X-Content-Type-Options: nosniff
< Date: Sat, 09 Nov 2024 02:43:10 GMT
< Content-Length: 1294
< 
（略）
```
