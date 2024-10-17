# 回显服务器

回显服务器程序会简单地回显收到的请求，包括请求的方法、路径、头部和主体，读者可据此判断所发送的请求是否符合预期。

后续章节的示例演示，也会默认以回显服务器返回的输出内容来展示curl的执行效果。

## 编译和运行

回显服务器使用Go语言开发，因此需要安装Go开发环境来编译，详见[Go官方网站](https://go.dev/)或[Go中国镜像站](https://golang.google.cn/)。

在安装配置完Go工具链后，进入回显服务器源代码目录：

```shell
$ cd echo-server
```

有两种方式可以运行服务器，一种是直接运行，Go会在后台编译源代码，并将生成的可执行程序放置到临时目录并执行：

```shell
$ go run .
Start listening on :8080
```

另一种方式是先一次性编译出可执行程序，以后直接运行编译后的程序即可：

```shell
# 编译
$ go build .
```

```shell
# 运行
$ ./server
Start listening on :8080
```

## 指定运行参数

### 指定监听端口

监听端口默认为8080，可通过选项`-port`更改：

```shell
$ go run . -port 3000
Start listening on :3000
```

```shell
$ ./server -port 3000
Start listening on :3000
```

### 以HTTPS方式运行

服务器可以按HTTPS（TLS）连接方式运行：

- 选项`-cert`指定TLS证书文件路径，PEM格式
- 选项`-key`指定TLS证书对应私钥文件路径，PEM格式
- HTTPS模式下默认端口为8443，同样可以通过选项`-port`更改

```shell
go run . -port 9443 -cert /tmp/mysite.crt -key /tmp/mysite.key
```

```shell
./server -port 9443 -cert /tmp/mysite.crt -key /tmp/mysite.key
```

#### 生成测试证书

可以借助openssl命令行工具生成测试证书：

```shell
openssl req -x509 -newkey rsa:2048 -keyout /tmp/mysite.key -nodes -subj '/C=CN/ST=ZheJiang/L=HangZhou/O=CompanyName/OU=DepartmentName/CN=www.mysite.com' -days 365 -sha256 -out /tmp/mysite.crt
```

## 测试运行效果

假设服务器已经用默认选项启动，即监听在端口8080且非HTTPS连接模式，此时在命令行中没有任何输出内容，可能给人造成一种“假死”的错觉，请保持该命令行运行而不要中断或关闭它。

现在，再打开一个新的命令行终端，让我们来试一下用curl作为客户端发请求给服务器：

```shell
$ curl http://localhost:8080/foo/bar

================================
Request 1
================================

GET /foo/bar HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/8.10.0
```

在运行curl的终端上，如果看到类似以上的内容，那么恭喜，您已经成功运行了回显服务器。

与此同时，在运行服务器的终端上也会打印出相同的请求信息。这很重要，因为当使用`HEAD`方法来请求服务器时，服务器是无法将请求信息放入响应体（Response Body）输出给客户端的，此时就只能通过服务器端的输出，观察请求发送正确与否。

注意最上方的标题"Request N"，其中`N`是请求的顺序编号。

现在请尝试使用图形化的Web浏览器打开服务器地址，通过服务器回显观察浏览器给服务器发送了哪些请求头。然后尝试刷新页面，如果顺序编号每次总比上次变化+2，请通过观察服务器端终端输出的信息来解释其中的原因。
