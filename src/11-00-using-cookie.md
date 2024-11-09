# 使用Cookie

Cookie使服务器端保持客户端状态成为可能。当服务器端通过响应头`Set-Cookie`设置一个键值对时，客户端可以将其保存，并在后续发起请求时附带`Cookie`请求头，将键值对再发回给服务器，于是服务器端就可以重建上次请求时的状态信息。

本章主要使用EHFS来做实验，在3个路径下，服务器端会分别设置3个不同的Cookie：

```shell
$ ehfs -l 8081 -r /tmp/ --header \
:/foo:set-cookie:foo=1 \
:/bar:set-cookie:bar=2 \
:/baz:set-cookie:baz=3
```

请求`/foo`将设置`foo=1`，请求`/bar`将设置`bar=2`，请求`/baz`将设置`baz=3`。
