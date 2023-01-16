# go-shentong
神通数据库官方驱动

# 安装
```shell
go get github.com/Mystery00/go-shentong
```

新建文件 aci.pc，内容如下：
```
includedir=%aci头文件%
libdir=%aci库文件%

Name: aci
Description: ShenTong Database Client
Version: 7.0.8
Cflags: -I${includedir}
Libs: -L${libdir} -laci
```

- aci头文件项目中已经带了，修改路径指向include就行，假设仓库源码存放在 /root/shentong 下，那么 includedir=/root/shentong/include/
- aci库文件需要根据部署平台从数据库中的安装目录中拷贝$SZ_OSCAR_HOME/aci/，比如在arm平台下，将$SZ_OSCAR_HOME/aci/arm64/libaci.so文件拷贝到仓库源码目录下，然后指定libdir为正确路径

如下，是一个示例：
```
includedir=/root/shentong/include/
libdir=/root/shentong/lib/linux64

Name: aci
Description: ShenTong Database Client
Version: 7.0.8
Cflags: -I${includedir}
Libs: -L${libdir} -laci
```
示例文件路径树如下，当前路径为仓库目录。
```
.
├── aci.go
├── aci.go.h
├── aci.pc
├── c_helpers.go
├── connection.go
├── connector.go
├── dsn.go
├── globals.go
├── go.mod
├── include
│   ├── aciapi.h
│   ├── acidef.h
│   ├── aciext.h
│   ├── aci.h
│   ├── acitype.h
│   └── oci.h
├── lib
│   ├── linux64
│   │   └── libaci.so
│   └── win64
│       ├── aci.dll
│       └── aci.lib
├── rows.go
├── statement.go
└── test
    └── test.go
```

# 编译
设置环境变量 `PKG_CONFIG_PATH` ，值为aci.pc所在目录，然后执行 go build 即可。

# 运行
记得先执行以下代码
```go
shentong.Init()
```
