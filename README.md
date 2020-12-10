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
export PATH="/path/to/qn:$PATH"
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
export PATH="/path/to/qn:$PATH"
```

## 使用

第一次使用时需要配置你的ak sk 以及bucket

`qn account -w ak sk bucket(把sk sk bucket换成你的ak sk以及你选择存储的空间)`

```shell
# 假设test.png的位置是/path/to/test.png
qn put /path/to/test.png
# 结果为:
upload successfully
外链为:http://picture.nj-jay.com/test.png
```

使用qn -h查看支持的命令

## 目前支持的命令

* put
> qn put filepath

## TODO

支持通配符

qn put ./*.png

qn put ./*