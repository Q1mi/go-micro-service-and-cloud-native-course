# lesson18


## 准备环境

之前gRPC版本的hello服务（server端和client端）


### pb定义

`hello.proto`存放于项目下的pb文件夹中。
```protobuf
syntax = "proto3";  // 版本声明

option go_package = "hello_client/pb";  // 项目中import导入生成的Go代码的名称

package pb;  // client与server必须一致！


// 定义服务
service Greeter {
    // 定义方法
    rpc SayHello (HelloRequest)returns (HelloResponse){}
}

// 定义的消息
message HelloRequest{
    string name = 1;  // 字段序号
}

message HelloResponse{
    string reply = 1; 
}
```

### protoc编译命令

```bash
protoc --proto_path=pb \
--go_out=pb --go_opt=paths=source_relative \
--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
hello.proto
```

## 引入外部定义
```go
import errpb "google.golang.org/genproto/googleapis/rpc/errdetails"
```