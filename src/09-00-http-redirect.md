# HTTP重定向

有时候，当我们请求一个资源时，服务器端有可能会返回一个状态码为3XX的重定向信息，指示所请求的资源在另一个URL位置。

本章我们通过EHFS来配置一个重定向URL路径。假设我们共享了本地目录`/tmp`，且其中有一个`files`子目录。现在我们希望当用户请求`/docs`目录时能够重定向到`/files/`下；请求`/archive`目录时能够重定向到`http://127.0.0.31:8081/files/`：

```shell
$ mkdir /tmp/files
$ ehfs -l 8081 -r /tmp \
--redirect @^/docs@/files/ \
--redirect @^/archive@http://127.0.0.31:8081/files/
```

尝试用浏览器访问`http://localhost:8081/docs`，会被重定向到`http://localhost:8081/files/`。
