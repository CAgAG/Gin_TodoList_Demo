## 基于 Gin + Go-Micro V4 + RabbitMQ 实现的简单备忘录

组件:
- 服务发现: etcd
- 鉴权(Token 验证): JWT
- 网关和模块通信: gRPC
- 微服务: Go-Micro V4
- 异步消息(备忘录创建): RabbitMQ

## 主要功能

- 用户登录、注册(jwt-go鉴权)
- 增删查改 备忘录
  - 增: 异步消息创建

## 项目结构
### 项目总体
```
/
├── app/                   // 微服务
│   ├── gateway            // 网关
│   ├── task               // 任务模块微服务
│   └── user               // 用户模块微服务
├── conf                   // 配置文件
├── grpc_proto/            // gRPC protoc文件
│   └── pb                 // 生成的pb文件
├── pkg/                   // 自定义的辅助包
│   ├── status             // 状态码
└── └── utils              // 工具函数(JWT、中间件和http响应)
```

### gateway 网关部分
```
app/gateway/
├── cmd                    // 启动入口
├── http_func              // Gin处理HTTP请求的函数
├── middleware             // 中间件(JWT)
├── router                 // 路由
└── rpc                    // rpc 调用
```

### 任务模块
```
app/task/
├── cmd                    // 启动入口
├── repository/            // 持久层
│    ├── db/               // 视图层
│    │    ├── dao          // 对数据库进行操作
│    │    └── model        // 定义数据库的模型
│    └── mq/               // RabbitMQ 函数
│         └── listener     // 监听 RabbitMQ 消息
├── script                 // 监听 RabbitMQ 的脚本
└── service                // task 服务实现
```

### 用户模块
```
app/task/
├── cmd                    // 启动入口
├── repository/            // 持久层
│    └── db/               // 视图层
│         ├── dao          // 对数据库进行操作
│         └── model        // 定义数据库的模型
└── service                // user 服务实现
```

### 项目配置
```
# debug开发模式,release生产模式
[appmode]
AppMode = debug

[mysql]
Db = mysql
# mysql的ip地址
DbHost = "127.0.0.1"
# mysql的端口号,默认3306
DbPort = 3306
# mysql user
DbUser = test_root
# mysql password
DbPassWord = 123456
# 数据库名字
DbName = todo_list_demo

[rabbitmq]
RabbitMQ = amqp
# RabbitMQ 用户名
RabbitMQUser = guest
# RabbitMQ 密码
RabbitMQPassWord = guest
# RabbitMQ 地址
RabbitMQHost = localhost
# RabbitMQ 端口
RabbitMQPort = 5672

[etcd]
# Etcd 端口
EtcdHost = localhost
# Etcd 端口
EtcdPort = 2379

[server]
# 网关服务地址
GatewayServiceAddress = localhost:4000
# 用户服务地址
UserServiceAddress = 127.0.0.1:8082
# 任务(备忘录)服务地址
TaskServiceAddress = 127.0.0.1:8083
```

## 项目运行
### 本机运行
#### 安装好对应的软件
- Mysql
- RabbitMq
- Protoc
- Etcd

#### 或是使用docker 配置环境
安装 docker 和 docker-compose 后, 
在项目根目录运行命令
```bash
docker-compose up -d
```
> Ubuntu
> 
> 启动 docker `sudo service docker start`
> 
> 关闭 docker `sudo service docker stop`
>
> 查看数据库 `docker exec -it mysql /bin/bash` 

### 运行命令
```bash
go mod tidy
go run app/user/cmd/main.go
go run app/task/cmd/main.go
go run app/gateway/cmd/main.go
```

## 接口

### 使用
Postman导入TodoList_demo.postman_collection.json

测试是否可以连接

Get: http://127.0.0.1:4000/api/v1/ping

### 用户模块

用户注册

Post: http://127.0.0.1:4000/api/v1/user/register

用户登录

Post: http://127.0.0.1:4000/api/v1/user/login

### 任务(备忘录)模块

备忘录创建

Post: http://127.0.0.1:4000/api/v1/task/create

备忘录更新

Post: http://127.0.0.1:4000/api/v1/task/update

备忘录删除

Post: http://127.0.0.1:4000/api/v1/task/delete

备忘录获取

Get: http://127.0.0.1:4000/api/v1/task/get

用户所有备忘录获取

Get: http://127.0.0.1:4000/api/v1/task/list

[参考](https://github.com/CocaineCong/micro-todoList)
