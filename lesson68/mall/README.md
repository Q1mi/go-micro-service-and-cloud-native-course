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



## RPC调用传递metadata

### 几个知识点
1. 什么是metadata ?什么样的数据应该存入metadata? 它和请求参数有什么区别？
2. gRPC的拦截器：客户端的拦截器和服务端的拦截器

### go-zero项目添加client端拦截器

order服务的search接口中添加拦截器，添加一些requestID、token、userID等数据

几个关键点：
1. 什么时候存入metadata
2. 怎么存metadata
3. 拦截器中如何通过context传值
4. context存值取值操作

### go-zero项目添加server端拦截器
1. 拦截器怎么加？什么时候加？
2. 拦截器的业务逻辑怎么写？
3. 服务端拦截器如何从metadata取值



## 错误处理

```json
{
  "code": 10001,
  "msg": "内部错误"
}
```
1. 定义自定义错误格式
2. 业务代码中按需返回自定义的错误
3. 告诉go-zero框架处理一下我们的自定义错误


## go-zero框架中goctl模板

模板的用处：用来生成代码的 

goctl指令生成代码时就是根据模板来生成代码的。 
```bash
goctl api go -api order.api -dir . -style=goZero
```

### goctl template

关于模板的官方文档：https://go-zero.dev/cn/docs/goctl/template-cmd

查看默认的存放模板文件的路径： GOCTL_HOME=/Users/liwenzhou/.goctl
```bash
goctl env  
```

初始化模板，在自己电脑上生成一份模板文件
```bash
goctl template init 
```

具体使用：
1. 找到模板文件并按需修改
2. 生成代码（有同名文件就不会生成）
