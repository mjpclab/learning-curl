# 设置请求头

## 通用设置选项

curl提供了许多选项来设置请求头，最一般的用法是通过`-H`或`--header`来设置，其值格式为`Header-Name: value`，可以出现多次。例如：

```shell
$ curl -H 'Accept: text/html' -H 'Accept-Language: zh-CN' http://localhost:8080

================================
Request 21
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: text/html
Accept-Language: zh-CN
User-Agent: curl/8.10.1
```

也可以设置为`@file_path`的格式从文件加载：

```shell
$ cat << EOF > /tmp/headers.txt
> Accept: text/html,text/xml;q=0.9
> Accept-Language: en-US
> EOF

$ curl -H @/tmp/headers.txt http://localhost:8080

================================
Request 22
================================

GET / HTTP/1.1
Host: localhost:8080
Accept: text/html,text/xml;q=0.9
Accept-Language: en-US
User-Agent: curl/8.10.1
```
