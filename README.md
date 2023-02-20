# qiniu-cli

如果你选择了七牛云作为图床

而又不想在浏览器中上传图片或者文件等

那么你可以选择这个工具

qiniu-cli一个基于命令行的文件上传工具

仅仅需要使用`qn put filename`即可

如果你是一个开发者，请继续往下看

如果你只是寻找一个免费图床，请点击下方链接

## 配置

首先在七牛云上找到账号ak sk [链接](https://portal.qiniu.com/user/key)

并选择你要将文件存储在哪个空间(bucket)中 [链接](https://portal.qiniu.com/kodo/bucket)

## 使用

```shell
go install github.com/jaronnie/qiniu-cli@latest
// 将 qiniu-cli 重命名为 qn
```


第一次使用时需要配置你的ak sk 以及bucket

`qn account -w ak sk bucket(把sk sk bucket换成你的ak sk以及你选择存储的空间)`

### 示例

```shell
# 假设test.png的位置是/path/to/test.png
qn put /home/jay/images/test.png
qn put ../images/test.png
qn put test.png
# 结果为:
upload successfully
外链为:http(s)://resource.jaronnie.com/test.png

# 你也可以上传多个文件
qn put ~/.pic/maven-compile.png ~/.pic/maven-package.png
# 结果为:
upload successfully
外链为:http(s)://resource.jaronnie.com/test.png

upload successfully
外链为:http(s)://resource.jaronnie.com/test.png

## 如果你要覆盖上传 请使用-w选项
qn put -w test.png
qn put -w test1.png test2.png

# 可以将网页文件上传
qn put https://xx/test.png
外链为:http(s)://resource.jaronnie.com/test.png
```

使用qn -h查看支持的命令

## 与 Typora 天然集成

如果你使用的是 typora 这款软件编写 markdown 

那么这个工具将是你的不二之选

能够极大的提高你的效率

[点击这个地方查看详细配置](docs/typora.md)

## 目前支持的命令

| 子命令  | 如何使用                   | 注意事项                                                     |
| ------- | -------------------------- | ------------------------------------------------------------ |
| account | qn account -w ak sk bucket | 替换你在七牛云上的ak sk bucket                               |
| put     | qn put [-w] filename [filename] | filename可填写绝对路径和相对路径，注意使用/而不是\，可同时传多个文件，用空格隔开 |
| ls      | qn ls 或 qn ls *.xxx       | xxx为文件的后缀 如 qn ls *.png                               |
