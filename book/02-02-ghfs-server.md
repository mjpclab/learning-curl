# GHFS（Go HTTP File Server）

GHFS是基于命令行的HTTP文件共享服务器。在后续章节如文件上传、范围请求等，单单通过回显服务器回显请求，还是过于抽象。为了进行更真实的演示，将使用GHFS来演示和验证这些操作。

GHFS使用Go语言开发，项目地址：https://github.com/mjpclab/go-http-file-server

可以直接通过发布页下载编译好的二进制程序，也可以自行编译：

```shell
go build .
```

假设要把/tmp共享出来，服务器运行在8081端口，那么只需运行

```shell
ghfs -l 8081 -r /tmp/
```

根据提示用浏览器打开URL便能看到目录列表。我们将在用到GHFS时再做详细介绍。
