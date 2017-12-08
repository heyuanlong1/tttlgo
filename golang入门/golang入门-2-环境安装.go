安装包下载地址为：
.官网： https://golang.org/dl/。（已被墙）
.Go语言中文网： https://studygolang.com/dl
.golang中国： https://www.golangtc.com/download

linux环境：

为了构建 Go 1.x （x ≥ 5），需要先安装 Go 1.4 到 $GOROOT_BOOTSTRAP.（因为Go 1.5以后 将使用 Go 1.4 进行构建，）
第一步：先下载1.4源码和1.7源码
第二步：cd到go1.4/src里，执行.all.bash...........安装完毕
第三步：export GOROOT_BOOTSTRAP=...go1.4
第四步：cd到go1.7/src里，执行.all.bash...........安装完毕
第五步：设置go1.7的环境
{
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOARCH=amd64
export GOOS=linux
}

相关网站
http://blog.csdn.net/h1023417614/article/details/52326681
http://blog.csdn.net/u011019726/article/details/77584708
http://blog.csdn.net/a1160712069/article/details/78257307


windows环境：

下载msi文件
双击*.msi安装包，按提示安装，推荐一路默认。

相关网站
https://studygolang.com/articles/10918


关于GOPATH环境变量：

GOPATH允许多个目录，当有多个目录时，请注意分隔符，多个目录的时候Windows是分号，Linux系统是冒号，当有多个GOPATH时，默认会将go get的内容放在第一个目录下

以上 $GOPATH 目录约定有三个子目录：
1.src 存放源代码（比如：.go .c .h .s等）
2.pkg 编译后生成的文件（比如：.a）
3.bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）

