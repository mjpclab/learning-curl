# 带有TLS的请求

目前线上实际使用的服务器都打开了HTTPS功能，也就是TLS协议。TLS协议在TCP层和HTTP层中间加入了一层安全机制。有关TLS的详细信息超出了本书的范围，不会加以讨论。

## 运行TLS回显服务器

我们可以再打开一个终端来运行回显服务器，指定其工作在HTTPS（即TLS）模式。

先生成一张自签名的测试证书：

```shell
openssl req -x509 -newkey rsa:2048 -keyout /tmp/mysite.key -nodes -subj '/C=CN/ST=ZheJiang/L=HangZhou/O=CompanyName/OU=DepartmentName/CN=www.mysite.com' -days 365 -sha256 -out /tmp/mysite.crt
```

然后启动回显服务器，指定证书和私钥位置：

```shell
$ cd echo-server/

$ go run . -cert /tmp/mysite.crt -key /tmp/mysite.key
Start listening on :8443
```

服务器默认监听在端口8443上。

## 尝试请求TLS回显服务器

```shell
$ curl https://localhost:8443
curl: (60) SSL certificate problem: self-signed certificate
More details here: https://curl.se/docs/sslcerts.html

curl failed to verify the legitimacy of the server and therefore could not
establish a secure connection to it. To learn more about this situation and
how to fix it, please visit the webpage mentioned above.
```

居然出错了！这里的主要原因是我们生成的自签名证书默认是不受信任的，因此在TLS协议握手过程中，在校验证书的阶段会失败。

对于如何将测试证书加入到受信任的列表中，不同的操作系统有不同的方法，请查阅相关资料，本书不作展开。不推荐将测试证书加入到受信任的证书列表中，以免引起安全问题。

由于只是练习，我们只需简单指定选项`-k`或`--insecure`就可以忽略安全验证中遇到的错误。

```shell
$ curl -k https://localhost:8443

================================
Request 1
================================

GET / HTTP/2.0
Host: localhost:8443
Accept: */*
User-Agent: curl/8.10.1
```

## 协商TLS版本

目前常用的TLS版本有1.1、1.2和1.3。客户端会和服务器端协商出双方都能支持的（一般选最大的）版本。可以通过选项`-v`让curl打印出详细的日志信息来检查：

```shell
 curl -k -v https://localhost:8443
* Host localhost:8443 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8443...
* ALPN: curl offers h2,http/1.1
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_128_GCM_SHA256 / x25519 / RSASSA-PSS # <==== 注意这里，协商使用了TLS 1.3版本
* ALPN: server accepted h2
（略）
```

可以通过一些选项来限制curl作为客户端可用于协商的TLS版本：

- `-1`或`--tlsv1`指定协商的TLS版本下限为1.0
- `--tlsv1.0`指定协商的TLS版本下限为1.0
- `--tlsv1.1`指定协商的TLS版本下限为1.1
- `--tlsv1.2`指定协商的TLS版本下限为1.2
- `--tlsv1.3`指定协商的TLS版本下限为1.3
- `--tls-max`设置协商的TLS版本上限值

```shell
$ curl -k -v --tlsv1.1 --tls-max 1.2 https://localhost:8443
* Host localhost:8443 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8443...
* ALPN: curl offers h2,http/1.1
* TLSv1.2 (OUT), TLS handshake, Client hello (1):
* TLSv1.2 (IN), TLS handshake, Server hello (2):
* TLSv1.2 (IN), TLS handshake, Certificate (11):
* TLSv1.2 (IN), TLS handshake, Server key exchange (12):
* TLSv1.2 (IN), TLS handshake, Server finished (14):
* TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
* TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.2 (OUT), TLS handshake, Finished (20):
* TLSv1.2 (IN), TLS handshake, Finished (20):
* SSL connection using TLSv1.2 / ECDHE-RSA-AES128-GCM-SHA256 / x25519 / RSASSA-PSS # <==== 协商使用了TLS 1.2版本
* ALPN: server accepted h2
（略）
```

## 指定HTTP应用层版本

对于运行在TCP+TLS层上的HTTP，可使用HTTP 1.1或HTTP 2版本。在上面的示例中，curl日志提示了应用层协议协商（ALPN）的结果为，服务器端选择使用HTTP 2：

```
* ALPN: curl offers h2,http/1.1
* ALPN: server accepted h2
```

有一些选项可以指定HTTP的版本：

- `--http0.9`
- `-0`或`--http1.0`
- `--http1.1`
- `--http2`
- `--http3`

并不是每一种版本都能随意使用，例如HTTP 2只能运行在TCP+TLS层之上，HTTP 1.0只能运行在纯粹的TCP之上。

```shell
$ curl -k -v --http1.1 https://localhost:8443
* ALPN: curl offers http/1.1
* ALPN: server accepted http/1.1
（略）
```
