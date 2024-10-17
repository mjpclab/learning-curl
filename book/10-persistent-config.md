# 持久化配置

有些选项可能在每次执行curl时都希望加上（例如强制解析某个域名的IP），可以将他们写入配置文件。

默认的配置文件路径通常为用户目录下的`.curlrc`。也可以在执行curl命令时，通过选项`-K`或`--config`指定配置文件位置。

```shell
$ echo 'resolve = "localhost:8080:127.0.0.15"' > ~/.curlrc
$ curl -v http://localhost:8080
* Added localhost:8080:127.0.0.15 to DNS cache
* Hostname localhost was found in DNS cache
*   Trying 127.0.0.15:8080...
* Connected to localhost (127.0.0.15) port 8080
（略）
```

```shell
$ echo -e '-v\n--user-agent mycurl/1.0.0' > /tmp/curl-config

$ curl -I -K /tmp/curl-config http://localhost:8080
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> HEAD / HTTP/1.1
> Host: localhost:8080
> User-Agent: mycurl/1.0.0
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 200 OK
HTTP/1.1 200 OK
< Content-Type: text/plain
Content-Type: text/plain
< Date: Wed, 16 Oct 2024 15:18:19 GMT
Date: Wed, 16 Oct 2024 15:18:19 GMT
< Content-Length: 158
Content-Length: 158
< 

* Connection #0 to host localhost left intact
```

可见配置文件中的格式，可以使用和命令行选项相同的形式，也可以使用`长选项名称（不带--前缀） = 值`的形式，如值中间包含空格，须用**双引号**包裹，每个选项占用单独的一行。
