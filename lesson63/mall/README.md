# go-zero RPC服务


## 编写RPC

1. 编写pb文件，并生成代码
2. 完善配置结构体和配置文件（结构体和yaml文件一定一定一定要对应上！）
3. 完善ServiceContext
4. 完善rpc业务逻辑


### rpc服务测试工具

一个测试grpc服务的ui工具
https://github.com/fullstorydev/grpcui


安装：
```bash
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
```

确保你电脑上的 $GOPATH/bin 目录，被添加到环境变量里了

使用 

其中`localhost:8080` 是你的RPC服务的地址
```bash
grpcui -plaintext localhost:8080
```

如果出现下面的情况：
```bash
grpcui -plaintext localhost:8080
Failed to compute set of methods to expose: server does not support the reflection API
```

要想用 grpcui 测试RPC服务，需要让go-zero rpc服务以 dev或test 模式运行,
需要在配置文件中指定 mode .

```yaml
Name: user.rpc
Mode: dev
ListenOn: 0.0.0.0:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: user.rpc
Mysql:
  DataSource: root:root1234@tcp(127.0.0.1:3306)/db3?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheRedis:
  - Host: 127.0.0.1:6379
```


