## 下载

| 平台支持      | 下载                                                         |
| ------------- | ------------------------------------------------------------ |
| windows amd64 | [qn_windows_v1.exe](https://pic.gocloudcoder.com/qn_windows_v1.exe) |
| linux amd64   | [qn_linux_v1](https://pic.gocloudcoder.com/qn_linux_v1)      |
| mac amd64     | [qn_mac_v1](https://pic.gocloudcoder.com/qn_mac_v1)          |

## 配置

### linux

```shell
# 假设qn_linux_v1文件所在的路径为 /path/to/qn_linux
mv /path/to/qn_linux  /path/to/qn
chmod +x /path/to/qn
# 在~/.bashrc或者~/.zshrc中添加一行
export PATH="/path/to:$PATH"
```

### win

```shell
# 修改qn_windows_v1.exe为qn.exe
#把qn.exe所在的目录添加为环境变量即可
```

### mac

```shell
# 假设qn_mac_v1文件所在的路径为 /path/to/qn_mac
mv /path/to/qn_mac /path/to/qn
chmod +x /path/to/qn
# 在~/.bashrc或者~/.zshrc中添加一行
export PATH="/path/to:$PATH"
```

## 使用

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

# 你也可以上传多个文件
qn put ~/.pic/maven-compile.png ~/.pic/maven-package.png
# 结果为:
upload successfully
外链为:
http://picture.nj-jay.com/maven-compile.png

upload successfully
外链为:
http://picture.nj-jay.com/maven-package.png
```

使用qn -h查看支持的命令