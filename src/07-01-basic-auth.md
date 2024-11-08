# 基本身份验证

HTTP协议支持用户身份验证，其中最基础的一种验证方式为基本认证（Basic Auth）。

本节我们使用GHFS来做实验，先准备一下环境，然后启动GHFS：

```shell
$ mkdir /tmp/share
$ echo -n '0123456789abcdef' > /tmp/share/hex.txt
$ ghfs -l 8081 -r /tmp/share/ --auth / --user foobar:secrets
```

从URL的根`/`开始要求基本身份验证，可用于验证的用户名为`foobar`，密码为`secrets`。

## 通过浏览器开发者工具探索基本认证协议

让我们暂时离开一下命令行，用浏览器打开GHFS服务器地址，你将会看到身份认证质询对话框。现在先取消认证，打开浏览器开发者工具，一般位于主菜单->更多工具->开发者工具中，切换到“网络”面板，然后重新刷新页面。“网络”面板中应该会出现一条请求`/`的信息，选中它以查看请求的详细信息：

![要求基本验证](image/basic-auth-403.png)

响应有如下特征：

- 状态码为`401 Unauthorized`，代表用户身份认证失败
- 响应头中出现`www-authenticate: Basic realm="REALM"`

当同时满足以上条件，浏览器就会弹出认证质询对话框，要求用户输入用户名和密码。

现在让我们输入正确的用户名和密码，并提交。我们会发现，现在可以正常显示页面了。再次选择查看`/`的请求信息：

![基本验证通过](image/basic-auth-200.png)

请求有如下特征

- 请求头包含`authorization: Basic Zm9vYmFyOnNlY3JldHM=`

响应有如下特征：

- 状态码为`200 OK`
- 响应头中依旧出现`www-authenticate: Basic realm="REALM"`

此种情况即代表用户认证已经成功。

请求头中`authorization`的认证数据实际上就是`用户名:密码`经过base64编码后的值，我们可以用`base64`命令来解码：

```shell
$ echo -n 'Zm9vYmFyOnNlY3JldHM=' | base64 -d
foobar:secrets
```

也可以用base64来编码，看看和浏览器编码后的值是否相同：

```shell
$ echo -n 'foobar:secrets' | base64
Zm9vYmFyOnNlY3JldHM=
```

如在认证过程中输入错误的用户名或密码，那么服务器端逻辑经过计算后认为认证失败，会再次回到认证质询的状态。

## 在curl中完成基本认证

选项`-u`或`--user`用于指定认证时使用的用户名及密码，格式为`username[:password]`。让我们先给出错误的用户名或密码：

```shell
$ curl -I -u foo:secrets http://localhost:8081
HTTP/1.1 401 Unauthorized
Www-Authenticate: Basic realm="files"
（略）
```

```shell
$ curl -I -u foobar:idontknow http://localhost:8081
HTTP/1.1 401 Unauthorized
Www-Authenticate: Basic realm="files"
（略）
```

返回的状态码始终为`401 Unauthorized`。

一旦提供了正确的用户名和密码，响应状态码将变为`200 OK`：

```shell
$ curl -I -u foobar:secrets http://localhost:8081
HTTP/1.1 200 OK
Www-Authenticate: Basic realm="files"
（略）
```

指定用户名和密码也可以不使用`-u`而是附加在主机名之前，并添加`@`作为分隔符：

```shell
$ curl -I http://foobar:secrets@localhost:8081
HTTP/1.1 200 OK
Www-Authenticate: Basic realm="files"
（略）
```

无论使用那种方式提供用户名和密码，都不够安全，因为shell通常会把最近用过的命令保存在历史记录中。可以在`-u`选项后面只提供用户名，curl会以交互式方式询问密码：

```shell
$ curl -I -u foobar http://localhost:8081
Enter host password for user 'foobar': #（在此输入密码）
HTTP/1.1 200 OK
Www-Authenticate: Basic realm="files"
（略）
```

由于`-u`选项只是组装了`authorization`请求头，自然也可以手动构造base64编码的认证信息：

```shell
curl -I \
-H 'authorization: Basic Zm9vYmFyOnNlY3JldHM=' \
http://localhost:8081
```

输出效果与上例相同。
