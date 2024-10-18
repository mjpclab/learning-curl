# 像浏览器一样发出请求

当在浏览器地址栏输入网址并前往后，其会使用`GET`方法请求对应URL。可以使用curl来模仿这种行为：

```shell
curl http://localhost:8080/index.html
```

# 指定HTTP请求的方法（method）

选项`-X`或`--request`指定请求所使用的方法，由于在curl中HTTP请求的默认方法就是`GET`，因此没必要像下面这样显式指定使用`GET`方法：

```shell
curl -X GET http://localhost:8080/index.html
```

如果要使用`POST`方法来发送请求，需要显式指定：

```shell
curl -X POST http://localhost:8080/notify
```

也可以在URL末尾带上查询字符串（Query String）作为参数，注意用引号包裹URL，以防shell对特殊字符的处理：

```shell
curl -X POST "http://localhost:8080/notify?user=foo&event=bar"
```

# 指定请求头

如要在请求中发送指定的请求头，通过`-H`或`--header`指定：

```shell
curl -X POST -H 'accept: application/json' http://localhost:8080/notify
```

# 指定请求体

使用选项`-d`或`--data`来指定请求体的内容。例如请求一个Restful的API来发出一篇帖子，可能是这样的：

```shell
curl -X POST -H 'accept: application/json' -d '{"title":"Post Title","content":"Post Content"}' http://localhost:8080/add-post
```

# 参考手册

无论何时，不要忘记man手册，当你想不起来某一功能的选项时，请通过man查阅：

```shell
man curl
```
