## kratos 使用

布局

```bash
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── api // 下面维护了微服务使用的proto文件以及根据它们所生成的go文件
│   └── helloworld
│       └── v1
│           ├── error_reason.pb.go
│           ├── error_reason.proto
│           ├── error_reason.swagger.json
│           ├── greeter.pb.go
│           ├── greeter.proto
│           ├── greeter.swagger.json
│           ├── greeter_grpc.pb.go
│           └── greeter_http.pb.go
├── cmd  // 整个项目启动的入口文件
│   └── server
│       ├── main.go
│       ├── wire.go  // 我们使用wire来维护依赖注入
│       └── wire_gen.go
├── configs  // 这里通常维护一些本地调试用的样例配置文件
│   └── config.yaml
├── generate.go
├── go.mod
├── go.sum
├── internal  // 该服务所有不对外暴露的代码，通常的业务逻辑都在这下面，使用internal避免错误引用
│   ├── biz   // 业务逻辑的组装层，类似 DDD 的 domain 层，data 类似 DDD 的 repo，而 repo 接口在这里定义，使用依赖倒置的原则。
│   │   ├── README.md
│   │   ├── biz.go
│   │   └── greeter.go
│   ├── conf  // 内部使用的config的结构定义，使用proto格式生成
│   │   ├── conf.pb.go
│   │   └── conf.proto
│   ├── data  // 业务数据访问，包含 cache、db 等封装，实现了 biz 的 repo 接口。我们可能会把 data 与 dao 混淆在一起，data 偏重业务的含义，它所要做的是将领域对象重新拿出来，我们去掉了 DDD 的 infra层。
│   │   ├── README.md
│   │   ├── data.go
│   │   └── greeter.go
│   ├── server  // http和grpc实例的创建和配置
│   │   ├── grpc.go
│   │   ├── http.go
│   │   └── server.go
│   └── service  // 实现了 api 定义的服务层，类似 DDD 的 application 层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑
│       ├── README.md
│       ├── greeter.go
│       └── service.go
└── third_party  // api 依赖的第三方proto
    ├── README.md
    ├── google
    │   └── api
    │       ├── annotations.proto
    │       ├── http.proto
    │       └── httpbody.proto
    └── validate
        ├── README.md
        └── validate.proto
```


**增加api**

```bash
kratos proto add api/helloworld/v1/student.proto
```

根据proto文件生成对应的service

```bash
kratos proto server api/helloworld/v1/student.proto -t internal/service
```

生成的服务在internal的service里面

service 实现了 api 定义的服务，其实就是 student.proto 里定义的服务。它要把数据传输对象(比如 http request data) 传入到 internal/biz 里进行处理，它一般不会涉及业务逻辑代码。业务逻辑的组装会在 biz 里实现。

有了service/student.go
接下来，我们使用wire来注入代码

在/internal/service.go里面添加

```go
var ProviderSet = wire.NewSet(NewGreeterService, NewStudentService)
```

接下来在student里面添加服务，使其可以通过http来访问

internal/server/http.go
internal/server/grpc.go
```go

```