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



## 订单服务的检索接口


/api/order/search : 根据订单id查询订单信息
  -RPC-> userID -> user.GetUser 


课后作业：
1. 把订单服务自己完善一下


## go-zero中通过RPC调用其他服务

1. 配置RPC客户端(配置结构体和yaml配置文件都要加RPC客户端配置，注意：etcd的key要对应上)
2. 修改 ServiceContext （告诉生成的代码我现在有RPC的客户端了）
  - go-zero中的RPC服务会自动生成一份客户端代码
3. 编写业务逻辑（可以直接通过RPC客户端发起RPC调用了）



## 使用Consul作为注册中心

### 服务注册
1. 修改配置(配置结构体和yaml配置文件)
  - 引入github.com/zeromicro/zero-contrib/zrpc/registry/consul
  - 注释掉原来默认的etcd,添加consul相关配置
2. 服务启动的时候将服务注册到consul
  - consul.RegisterService(c.ListenOn, c.Consul)


### 服务发现
1. 修改yaml配置文件
  - Target: consul://127.0.0.1:8500/consul-user.rpc?wait=14s
2. 程序启动时 import _ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"

