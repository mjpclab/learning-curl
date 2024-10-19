# 用HTML提交表单

实际上，以`GET`方式提交表单在整个互联网上都非常罕见，几乎所有的表单都以`POST`方法提交。因此如未加说明，一般表单提交特指以`POST`方式提交表单。让我们来修改之前的HTML，将方法改为`POST`：

```diff
- <form action="http://localhost:8080/" method="get">
+ <form action="http://localhost:8080/" method="post">
```

最终的HTML如下：

```html
<html>
  <head>
    <title>form POST test</title>
  </head>
  <body>
    <form action="http://localhost:8080/" method="post">
      <div>x: <input name="x" /></div>
      <div>y: <input name="y" /></div>
      <div><button type="submit">Submit</button></div>
    </form>
  </body>
</html>
```

再次填写并提交表单，看看回显信息会有何变化：

```
================================
Request 11
================================

POST / HTTP/1.1
Host: localhost:8080
Content-Length: 11
Content-Type: application/x-www-form-urlencoded
（略）

x=100&y=200
```

有3处明显的不同之处：

- 请求方法为`POST`
- 请求头包含`Content-Type: application/x-www-form-urlencoded`
- 参数列表被移到了请求体（Request Body）中

满足以上3个条件，那么该请求就是一个表单提交。另外请求头`Content-Length`也提示了请求体的字节长度。

请求体中的键值对应该对URL元字符进行编码，以区分正常的元字符`&`和`=`、转义元字符`%`，以及其它URL元字符（如`+`），这也是“urlencoded”的含义所在。让我们再修改一下HTML文件：

```diff
- <div>x: <input name="x" /></div>
- <div>y: <input name="y" /></div>
+ <div>xy: <input name="x&y" /></div>
+ <div>username: <input name="user=name" /></div>
```

修改后的HTML如下：

```html
<html>
  <head>
    <title>form POST test</title>
  </head>
  <body>
    <form action="http://localhost:8080/" method="post">
      <div>xy: <input name="x&y" /></div>
      <div>username: <input name="user=name" /></div>
      <div><button type="submit">Submit</button></div>
    </form>
  </body>
</html>
```

用浏览器打开该文件，第一个输入框填写"100+200"，第二个输入框填写"Tom&Jerry"，提交表单后回显如下：

```
================================
Request 19
================================

POST / HTTP/1.1
Host: localhost:8080
Content-Length: 39
Content-Type: application/x-www-form-urlencoded
（略）

x%26y=100%2B200&user%3Dname=Tom%26Jerry
```

从中可以观察到提交的数据顺利被服务器接收，键值中的元字符被编码成了`%`+字节16进制表示：

- 数据1
  - 键：`x%26y`，即`x&y`编码后（url encoded）的字符串
  - 值：`100%2B200`，即`100+200`编码后的字符串
- 数据2
  - 键：`user%3Dname`，即`user=name`编码后的字符串
  - 值：`Tom%26Jerry`，即`Tom&Jerry`编码后的字符串

理论上，表单数据和查询字符串可以同时存在，在服务器端开发框架中，一般提供单独提取其中一种类型的值，以及提取两种类型合并后值的方法，例如`Request`、`Request.Form`和`Request.QueryString`等。

请尝试修改HTML，为`form`元素的`action`属性末尾添加`?key=value`，观察提交后回显的值。

# 用curl提交表单

## 在命令行指定提交数据

了解了表单提交数据的格式，我们就可以用curl来模仿了。选项`-d`或`--data`用于指定发送的请求体数据，它会使用POST方法提交，同时为请求头增加`Content-Type: application/x-www-form-urlencoded`：

```shell
$ curl --data 'username=Tom%26Jerry' http://localhost:8080

================================
Request 29
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 20
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.0

username=Tom%26Jerry
```

手工编码元字符比较麻烦，此时`--data-urlencode`就可以派上用场了：

```shell
$ curl --data-urlencode 'username=Tom&Jerry' http://localhost:8080

================================
Request 31
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 20
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.0

username=Tom%26Jerry
```

注意`--data-urlencode`假设键已被正确编码，只有其值需要处理。

一个`--data-urlencode`选项只能指定一个键值对，提交多个参数需要多次使用该选项来分别指定：

```shell
$ curl --data-urlencode 'pair1=Alice&Bob' --data-urlencode 'pair2=Tom&Jerry' http://localhost:8080

================================
Request 50
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 35
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.0

pair1=Alice%26Bob&pair2=Tom%26Jerry
```

`--data-urlencode`后跟随的字符串中首次出现的`=`被视为键和值的分隔符，因此指定`--data-urlencode a=b=c`则键为`a`，值为`b=c`（编码为`b%3Dc`），无法指定第n个`=`（n≥2）才是键值对的分隔符。

```shell
$ curl --data-urlencode 'a=b=c' http://localhost:8080/

================================
Request 51
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 7
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.0

a=b%3Dc
```

## 从文件加载提交数据

有时候请求体内容较大，对于需要多次发送的请求，每次在命令行上重复输入非常不便。这时可以把请求体保存在文件中，在curl相关选项值中使用特殊语法来引用文件。

对于不做编码处理的`-d`或`--data`，使用`@data_file`的格式指定文件来源：

```shell
$ echo -n 'username=Tom%26Jerry' > /tmp/data.txt

$ curl -d @/tmp/data.txt http://localhost:8080

================================
Request 1
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 20
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.0

username=Tom%26Jerry
```

对于`--data-urlencode`，文件中的内容会全部被编码，键需要写在命令行上，文件里只保留值的部分，使用`key@value_file`的格式来指定：

```shell
$ echo -n 'Alice&Bob' > /tmp/pair1.txt

$ echo -n 'Tom&Jerry' > /tmp/pair2.txt

$ curl --data-urlencode pair1@/tmp/pair1.txt --data-urlencode pair2@/tmp/pair2.txt http://localhost:8080

================================
Request 2
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 35
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.0

pair1=Alice%26Bob&pair2=Tom%26Jerry
```

如果把键也写入文件，所有内容都会被编码，结果可能是非预期的：

```shell
$ echo -n 'pair1=Alice&Bob' > /tmp/pair1.txt

$ echo -n 'pair2=Tom&Jerry' > /tmp/pair2.txt

$ curl --data-urlencode @/tmp/pair1.txt --data-urlencode @/tmp/pair2.txt http://localhost:8080

================================
Request 3
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 39
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.0

pair1%3DAlice%26Bob&pair2%3DTom%26Jerry
```

那么，万一真的需要提交以`@`开头的数据该怎么办呢？curl提供了`--data-raw`选项，它不会将`@`开头的数据当作特殊指令处理，而仅仅是字面量值：

```shell
$ curl --data-raw @/tmp/data.txt http://localhost:8080

================================
Request 4
================================

POST / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 14
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/8.10.0

@/tmp/data.txt
```

当请求体不是键值对的格式时，服务器端上层应用框架无法正确解析出字段。应用开发者如果定制了非标准格式，仍然可以通过自定义逻辑来解析请求体数据。
