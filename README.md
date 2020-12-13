# qiniu-cli

如果你选择了七牛云作为图床

而又不想在浏览器中上传图片或者文件等

那么你可以选择这个工具

qiniu-cli一个基于命令行的文件上传工具

仅仅需要使用`qn put filename`即可

## 下载

| 平台  | 下载地址                                                     |
| ----- | ------------------------------------------------------------ |
| linux | [https://picture.nj-jay.com/qn_linux](https://picture.nj-jay.com/qn_linux) |
| win   | [https://picture.nj-jay.com/qn_win.exe](https://picture.nj-jay.com/qn_win.exe) |
| mac   | [https://picture.nj-jay.com/qn_mac](https://picture.nj-jay.com/qn_mac) |

## 配置

首先在七牛云上找到账号ak sk [链接](https://portal.qiniu.com/user/key)

并选择你要将文件存储在哪个空间(bucket)中 [链接](https://portal.qiniu.com/kodo/bucket)

### linux

```shell
# 假设qn_linux文件所在的路径为 /path/to/qn_linux
mv /path/to/qn_linux  /path/to/qn
chmod +x /path/to/qn
# 在~/.bashrc或者~/.zshrc中添加一行
export PATH="/path/to:$PATH"
```

### win

```shell
# 修改qn_win.exe为qn.exe
#把qn.exe所在的目录添加为环境变量即可
```

### mac

```shell
# 假设qn_mac文件所在的路径为 /path/to/qn_mac
mv /path/to/qn_mac /path/to/qn
chmod +x /path/to/qn
# 在~/.bashrc或者~/.zshrc中添加一行
export PATH="/path/to:$PATH"
```

## 使用

第一次使用时需要配置你的ak sk 以及bucket

`qn account -w ak sk bucket(把sk sk bucket换成你的ak sk以及你选择存储的空间)`

### 注意事项

填写文件路径的时候请使用`/`而不是`\`

并保证文件名中没有空格

正确的路径为 

* ../filename(相对路径)
* filename(同一目录)
* /home/jay/filename(绝对路径)

### 示例

```shell
# 假设test.png的位置是/path/to/test.png
qn put /home/jay/images/test.png
qn put ../images/test.png
qn put test.png
# 结果为:
upload successfully
外链为:http://picture.nj-jay.com/test.png
```

使用qn -h查看支持的命令

## 与typora天然集成

如果你使用的是typora这款软件编写markdown

那么这个工具将是你的不二之选

能够极大的提高你的效率

[点击这个地方查看详细配置](typora.md)

## 目前支持的命令

* put
> qn put filepath
> 如果七牛上已经有该文件,默认不会覆盖上传.当你使用-w选项时可以进行覆盖上传

## TODO

- [ ] 支持通配符(批量传入)

- [ ] 改进路径问题(假设传入的路径中使用了`\`或者有空格,在传输之前首先处理这些问题(把路径中的`\`改为/并去除空格)
- [ ] 当七牛中已经存在同名的文件时,提示该信息，并告知使用-w选项可以覆盖该文件
