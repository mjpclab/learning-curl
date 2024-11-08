# 最大重定向次数

curl默认最多跟随重定向50次，可以通过选项`--max-redirs`来修改，设为`-1`代表不限制次数。

让我们对EHFS命令行参数做一些修改，使其产生多层重定向：

```shell
ehfs -l 8081 -r /tmp \
--redirect @^/docs@/files/ \
--redirect @^/archive@http://127.0.0.31:8081/files/ \
--redirect @^/manual1@/manual2 \
--redirect @^/manual2@/manual3 \
--redirect @^/manual3@/files/
```

新增了重定向路径`/manual1` -> `/manual2` -> `/manual3` -> `/files/`，共3次重定向。

使用curl的默认最大重定向次数来请求`/manual1`：

```shell
$ curl -v -L http://localhost:8081/manual1
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* Connected to localhost (::1) port 8081
* using HTTP/1.x
> GET /manual1 HTTP/1.1		# <==== 请求/manual1
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: /manual2		# <==== 重定向到/manual2
< Date: Sat, 09 Nov 2024 02:55:41 GMT
< Content-Length: 43
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8081/manual2'
* Re-using existing connection with host localhost
> GET /manual2 HTTP/1.1		# <==== 请求/manual2
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: /manual3		# <==== 重定向到/manual3
< Date: Sat, 09 Nov 2024 02:55:41 GMT
< Content-Length: 43
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8081/manual3'
* Re-using existing connection with host localhost
> GET /manual3 HTTP/1.1		# <==== 请求/manual3
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: /files/		# <==== 重定向到/files/
< Date: Sat, 09 Nov 2024 02:55:41 GMT
< Content-Length: 42
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8081/files/'
* Re-using existing connection with host localhost
> GET /files/ HTTP/1.1		# <==== 请求/files/
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
< Date: Sat, 09 Nov 2024 02:55:41 GMT
< Content-Length: 1294
< 
（略）
```

现在，指定最大重定向次数为2：

```shell
$ curl -v -L --max-redirs 2 http://localhost:8081/manual1
* Host localhost:8081 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8081...
* Connected to localhost (::1) port 8081
* using HTTP/1.x
> GET /manual1 HTTP/1.1		# <==== 请求/manual1
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: /manual2		# <==== 重定向到/manual2
< Date: Sat, 09 Nov 2024 02:57:30 GMT
< Content-Length: 43
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8081/manual2'
* Re-using existing connection with host localhost
> GET /manual2 HTTP/1.1		# <==== 请求/manual2
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: /manual3		# <==== 重定向到/manual3
< Date: Sat, 09 Nov 2024 02:57:30 GMT
< Content-Length: 43
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Issue another request to this URL: 'http://localhost:8081/manual3'
* Re-using existing connection with host localhost
> GET /manual3 HTTP/1.1		# <==== 请求/manual3
> Host: localhost:8081
> User-Agent: curl/8.10.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: /files/		# <==== 重定向到/files/
< Date: Sat, 09 Nov 2024 02:57:30 GMT
< Content-Length: 42
* Ignoring the response-body
* setting size while ignoring
< 
* Connection #0 to host localhost left intact
* Maximum (2) redirects followed
curl: (47) Maximum (2) redirects followed		# <==== 已达最大重定向次数
```
