## 安静模式

有时候，curl会主动输出一些额外信息，最典型的是当使用`-o`或`-O`将响应重定向到文件的时候，会显示一个下载状态指示信息：

```shell
$ curl -o /tmp/output http://localhost:8080
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   156  100   156    0     0   189k      0 --:--:-- --:--:-- --:--:--  152k
```

如要避免输出这类信息，可使用`-s`或`--silent`选项让curl保持安静：

```shell
$ curl -s -o /tmp/output http://localhost:8080
# 无输出
```

## 仍然输出错误信息

有时候，我们需要让curl工作在安静模式，但是当发生错误时仍需给出提示。选项`-S`或`--show-error`就是用于这个目的的，下面的例子中，我们故意写一个无效的IP地址URL来测试：

```shell
$ curl -o /tmp/output http://256.1.1.1/
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0curl: (6) Could not resolve host: 256.1.1.1 # <==== 这里先显示状态指示，然后是错误信息
```

```shell
$ curl -s -o /tmp/output http://256.1.1.1/
# 无输出
```

```shell
$ curl -sS -o /tmp/output http://256.1.1.1/
curl: (6) Could not resolve host: 256.1.1.1
```

## 直接指定不显示状态指示

curl 7.67.0起新增了`--no-progress-meter`用于禁用进度状态指示。

```shell
$ curl --no-progress-meter -o /tmp/output http://localhost:8080
# 无输出
```
