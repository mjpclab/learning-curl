# 追踪请求与响应

## 追踪传输数据

可使用`--trace`选项来追踪请求与响应数据，其功能与`--verbose`有些相似，不过侧重点在于传输的数据，该选项要求提供一个文件路径作为追踪数据日志的写入目标，可使用`-`指定输出到标准输出：

```shell
$ curl --trace - http://localhost:8080/
== Info: Host localhost:8080 was resolved.
== Info: IPv6: ::1
== Info: IPv4: 127.0.0.1
== Info:   Trying [::1]:8080...
== Info: Connected to localhost (::1) port 8080
== Info: using HTTP/1.x
=> Send header, 78 bytes (0x4e)
0000: 47 45 54 20 2f 20 48 54 54 50 2f 31 2e 31 0d 0a GET / HTTP/1.1..
0010: 48 6f 73 74 3a 20 6c 6f 63 61 6c 68 6f 73 74 3a Host: localhost:
0020: 38 30 38 30 0d 0a 55 73 65 72 2d 41 67 65 6e 74 8080..User-Agent
0030: 3a 20 63 75 72 6c 2f 38 2e 31 30 2e 31 0d 0a 41 : curl/8.10.1..A
0040: 63 63 65 70 74 3a 20 2a 2f 2a 0d 0a 0d 0a       ccept: */*....
== Info: Request completely sent off
<= Recv header, 17 bytes (0x11)
0000: 48 54 54 50 2f 31 2e 31 20 32 30 30 20 4f 4b 0d HTTP/1.1 200 OK.
0010: 0a                                              .
<= Recv header, 26 bytes (0x1a)
0000: 43 6f 6e 74 65 6e 74 2d 54 79 70 65 3a 20 74 65 Content-Type: te
0010: 78 74 2f 70 6c 61 69 6e 0d 0a                   xt/plain..
<= Recv header, 37 bytes (0x25)
0000: 44 61 74 65 3a 20 57 65 64 2c 20 31 36 20 4f 63 Date: Wed, 16 Oc
0010: 74 20 32 30 32 34 20 31 32 3a 33 31 3a 34 38 20 t 2024 12:31:48 
0020: 47 4d 54 0d 0a                                  GMT..
<= Recv header, 21 bytes (0x15)
0000: 43 6f 6e 74 65 6e 74 2d 4c 65 6e 67 74 68 3a 20 Content-Length: 
0010: 31 35 36 0d 0a                                  156..
<= Recv header, 2 bytes (0x2)
0000: 0d 0a                                           ..
<= Recv data, 156 bytes (0x9c)
0000: 0a 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d .===============
0010: 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d ================
0020: 3d 0a 52 65 71 75 65 73 74 20 31 32 0a 3d 3d 3d =.Request 12.===
0030: 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d ================
0040: 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 3d 0a 0a 47 =============..G
0050: 45 54 20 2f 20 48 54 54 50 2f 31 2e 31 0d 0a 48 ET / HTTP/1.1..H
0060: 6f 73 74 3a 20 6c 6f 63 61 6c 68 6f 73 74 3a 38 ost: localhost:8
0070: 30 38 30 0d 0a 41 63 63 65 70 74 3a 20 2a 2f 2a 080..Accept: */*
0080: 0d 0a 55 73 65 72 2d 41 67 65 6e 74 3a 20 63 75 ..User-Agent: cu
0090: 72 6c 2f 38 2e 31 30 2e 31 0d 0a 0d             rl/8.10.1...

================================
Request 12
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
== Info: Connection #0 to host localhost left intact
```

可以看到请求与响应的数据都给出了16进制表示。

## 仅显示传输内容

如不需要显示16进制数据，可以使用`--trace-ascii`选项代替`--trace`：

```shell
$ curl --trace-ascii - http://localhost:8080/
== Info: Host localhost:8080 was resolved.
== Info: IPv6: ::1
== Info: IPv4: 127.0.0.1
== Info:   Trying [::1]:8080...
== Info: Connected to localhost (::1) port 8080
== Info: using HTTP/1.x
=> Send header, 78 bytes (0x4e)
0000: GET / HTTP/1.1
0010: Host: localhost:8080
0026: User-Agent: curl/8.10.1
003f: Accept: */*
004c: 
== Info: Request completely sent off
<= Recv header, 17 bytes (0x11)
0000: HTTP/1.1 200 OK
<= Recv header, 26 bytes (0x1a)
0000: Content-Type: text/plain
<= Recv header, 37 bytes (0x25)
0000: Date: Wed, 16 Oct 2024 12:54:02 GMT
<= Recv header, 21 bytes (0x15)
0000: Content-Length: 156
<= Recv header, 2 bytes (0x2)
0000: 
<= Recv data, 156 bytes (0x9c)
0000: .================================.Request 13.===================
0040: =============..GET / HTTP/1.1
005f: Host: localhost:8080
0075: Accept: */*
0082: User-Agent: curl/8.10.1
009b: .

================================
Request 13
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
== Info: Connection #0 to host localhost left intact
```

## 显示传输与连接标识

无论是`--trace`还是`--trace-ascii`，如一次请求多个URL，都可以附加`--trace-ids`在每条日志的开头显示`[传输号-连接号]`的标识：

```shell
$ curl --trace-ascii - --trace-ids http://localhost:8080/foo http://localhost:8080/bar
[0-0] == Info: Host localhost:8080 was resolved.
[0-0] == Info: IPv6: ::1
[0-0] == Info: IPv4: 127.0.0.1
[0-0] == Info:   Trying [::1]:8080...
[0-0] == Info: Connected to localhost (::1) port 8080
[0-0] == Info: using HTTP/1.x
[0-0] => Send header, 81 bytes (0x51)
0000: GET /foo HTTP/1.1
0013: Host: localhost:8080
0029: User-Agent: curl/8.10.1
0042: Accept: */*
004f: 
[0-0] == Info: Request completely sent off
[0-0] <= Recv header, 17 bytes (0x11)
0000: HTTP/1.1 200 OK
[0-0] <= Recv header, 26 bytes (0x1a)
0000: Content-Type: text/plain
[0-0] <= Recv header, 37 bytes (0x25)
0000: Date: Tue, 29 Oct 2024 15:59:43 GMT
[0-0] <= Recv header, 21 bytes (0x15)
0000: Content-Length: 159
[0-0] <= Recv header, 2 bytes (0x2)
0000: 
[0-0] <= Recv data, 159 bytes (0x9f)
0000: .================================.Request 19.===================
0040: =============..GET /foo HTTP/1.1
0062: Host: localhost:8080
0078: Accept: */*
0085: User-Agent: curl/8.10.1
009e: .

================================
Request 19
================================

GET /foo HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
[0-0] == Info: Connection #0 to host localhost left intact
[1-0] == Info: Re-using existing connection with host localhost
[1-0] => Send header, 81 bytes (0x51)
0000: GET /bar HTTP/1.1
0013: Host: localhost:8080
0029: User-Agent: curl/8.10.1
0042: Accept: */*
004f: 
[1-0] == Info: Request completely sent off
[1-0] <= Recv header, 17 bytes (0x11)
0000: HTTP/1.1 200 OK
[1-0] <= Recv header, 26 bytes (0x1a)
0000: Content-Type: text/plain
[1-0] <= Recv header, 37 bytes (0x25)
0000: Date: Tue, 29 Oct 2024 15:59:43 GMT
[1-0] <= Recv header, 21 bytes (0x15)
0000: Content-Length: 159
[1-0] <= Recv header, 2 bytes (0x2)
0000: 
[1-0] <= Recv data, 159 bytes (0x9f)
0000: .================================.Request 20.===================
0040: =============..GET /bar HTTP/1.1
0062: Host: localhost:8080
0078: Accept: */*
0085: User-Agent: curl/8.10.1
009e: .

================================
Request 20
================================

GET /bar HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
[1-0] == Info: Connection #0 to host localhost left intact
```

上例中第一个请求的标识为`[0-0]`，传输号为`0`，连接号也为`0`。而第二个请求的标识为`[1-0]`，传输号为`1`，以区别于第一个URL请求，而连接号依旧为`0`，说明curl复用了第一个请求所使用的连接，从日志`[1-0] == Info: Re-using existing connection with host localhost`中也能印证这一点。

现在，我们故意使用两个不同的主机名去访问两个URL，看看标识会有何不同：

```shell
$ curl --trace-ascii - --trace-ids http://127.0.0.1:8080/foo http://127.0.0.2:8080/bar
[0-0] == Info:   Trying 127.0.0.1:8080...
[0-0] == Info: Connected to 127.0.0.1 (127.0.0.1) port 8080
[0-0] == Info: using HTTP/1.x
[0-0] => Send header, 81 bytes (0x51)
0000: GET /foo HTTP/1.1
0013: Host: 127.0.0.1:8080
0029: User-Agent: curl/8.10.1
0042: Accept: */*
004f: 
[0-0] == Info: Request completely sent off
[0-0] <= Recv header, 17 bytes (0x11)
0000: HTTP/1.1 200 OK
[0-0] <= Recv header, 26 bytes (0x1a)
0000: Content-Type: text/plain
[0-0] <= Recv header, 37 bytes (0x25)
0000: Date: Tue, 29 Oct 2024 16:10:31 GMT
[0-0] <= Recv header, 21 bytes (0x15)
0000: Content-Length: 159
[0-0] <= Recv header, 2 bytes (0x2)
0000: 
[0-0] <= Recv data, 159 bytes (0x9f)
0000: .================================.Request 23.===================
0040: =============..GET /foo HTTP/1.1
0062: Host: 127.0.0.1:8080
0078: Accept: */*
0085: User-Agent: curl/8.10.1
009e: .

================================
Request 23
================================

GET /foo HTTP/1.1
Host: 127.0.0.1:8080
Accept: */*
User-Agent: curl/8.10.1
[0-0] == Info: Connection #0 to host 127.0.0.1 left intact
[1-1] == Info:   Trying 127.0.0.2:8080...
[1-1] == Info: Connected to 127.0.0.2 (127.0.0.2) port 8080
[1-1] == Info: using HTTP/1.x
[1-1] => Send header, 81 bytes (0x51)
0000: GET /bar HTTP/1.1
0013: Host: 127.0.0.2:8080
0029: User-Agent: curl/8.10.1
0042: Accept: */*
004f: 
[1-1] == Info: Request completely sent off
[1-1] <= Recv header, 17 bytes (0x11)
0000: HTTP/1.1 200 OK
[1-1] <= Recv header, 26 bytes (0x1a)
0000: Content-Type: text/plain
[1-1] <= Recv header, 37 bytes (0x25)
0000: Date: Tue, 29 Oct 2024 16:10:31 GMT
[1-1] <= Recv header, 21 bytes (0x15)
0000: Content-Length: 159
[1-1] <= Recv header, 2 bytes (0x2)
0000: 
[1-1] <= Recv data, 159 bytes (0x9f)
0000: .================================.Request 24.===================
0040: =============..GET /bar HTTP/1.1
0062: Host: 127.0.0.2:8080
0078: Accept: */*
0085: User-Agent: curl/8.10.1
009e: .

================================
Request 24
================================

GET /bar HTTP/1.1
Host: 127.0.0.2:8080
Accept: */*
User-Agent: curl/8.10.1
[1-1] == Info: Connection #1 to host 127.0.0.2 left intact
```

由于请求了两个URL，依旧有两个不同的传输号。而这次由于目标主机不同，curl不得不创建新的TCP连接，因而也有两个不同的连接号。

## 显示时间

无论是`--trace`还是`--trace-ascii`，都可以附加`--trace-time`在每条日志的开头打印当前时间：

```shell
$ curl --trace-ascii - --trace-time http://localhost:8080/
20:54:53.540961 == Info: Host localhost:8080 was resolved.
20:54:53.540994 == Info: IPv6: ::1
20:54:53.540996 == Info: IPv4: 127.0.0.1
20:54:53.541025 == Info:   Trying [::1]:8080...
20:54:53.541090 == Info: Connected to localhost (::1) port 8080
20:54:53.541094 == Info: using HTTP/1.x
20:54:53.541139 => Send header, 78 bytes (0x4e)
0000: GET / HTTP/1.1
0010: Host: localhost:8080
0026: User-Agent: curl/8.10.1
003f: Accept: */*
004c: 
20:54:53.541161 == Info: Request completely sent off
20:54:53.541408 <= Recv header, 17 bytes (0x11)
0000: HTTP/1.1 200 OK
20:54:53.541444 <= Recv header, 26 bytes (0x1a)
0000: Content-Type: text/plain
20:54:53.541451 <= Recv header, 37 bytes (0x25)
0000: Date: Wed, 16 Oct 2024 12:54:53 GMT
20:54:53.541459 <= Recv header, 21 bytes (0x15)
0000: Content-Length: 156
20:54:53.541465 <= Recv header, 2 bytes (0x2)
0000: 
20:54:53.541470 <= Recv data, 156 bytes (0x9c)
0000: .================================.Request 14.===================
0040: =============..GET / HTTP/1.1
005f: Host: localhost:8080
0075: Accept: */*
0082: User-Agent: curl/8.10.1
009b: .

================================
Request 14
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.1
20:54:53.541531 == Info: Connection #0 to host localhost left intact
```

## 配置要显示的附加信息

实际上，`--trace-ids`和`--trace-time`都是`--trace-config`的特例。`--trace-config`接受一个逗号分割的列表，指出要附加的信息。例如`--trace-config ids,time`相当于分别指定`--trace-ids`和`--trace-time`。

`--trace-config`还支持一些其它值，详见curl官方文档。
