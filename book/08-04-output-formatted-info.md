# 输出格式化信息

有时我们希望得到有关一个请求的摘要性的信息，此时我们可以使用格式化选项来输出这样的信息。

`-w`或`--write-out`用于按指定的格式化字符串模板输出信息，模板中可以内嵌curl的预定义变量，格式为`%{variable_name}`，而要输出`%`字面量，则连续写两个百分号：`%%`。

假设我们只关心某个请求的响应状态码，那么可以这样指定格式化信息：

```shell
$ curl -w 'response status: %{http_code}\n' -o /dev/null -s http://localhost:8080
response status: 200
```

上例通过`-o /dev/null`丢弃了响应的主体，但通过`-w`输出了摘要信息，其中`http_code`为curl的内置变量。

`-w`默认输出到标准输出，如要输出到标准错误，指定特殊变量`stderr`即可：

```shell
$ curl -w 'response status: %{http_code}\n' -o /dev/null -s http://localhost:8080 2> /dev/null
response status: 200	# 丢弃stderr，但格式化信息从stdout输出，因而不受影响

$ curl -w '%{stderr}response status: %{http_code}\n' -o /dev/null -s http://localhost:8080 2> /dev/null
# 无输出
```

更多内嵌变量请参考curl的man手册。

`-w`也接受`@file_path`的格式从外部文件读取格式化字符串。

```shell
$ echo 'response status: %{http_code}\n' > /tmp/output-format
$ echo 'content type: %{content_type}\n' >> /tmp/output-format

$ curl -w @/tmp/output-format -o /dev/null -s http://localhost:8080
response status: 200
content type: text/plain
```
