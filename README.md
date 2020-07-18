<!--
 * @Author: haha_giraffe
 * @Date: 2020-01-30 17:09:09
 * @Description: file content
 -->

# Introduction

HAHAGO是一个轻量级的基础网络库，基于此实现RPC。开发此项目的原因主要是作为Golang学习过程中的一个练手项目。

1. 消息封装模块，对消息数据进行简单封装，能解决TCP粘包问题。
2. 多路由模块，服务器可以注册多路由方法，对于不同类型（ID）的消息实现分路由以处理不同的情况。
3. 协程池模块，服务器启动时候开启固定数量的Goroutine用于处理客户请求，通过管道进行连接通信。
4. 连接管理模块，对客户端上线或下线时候可以注册独立的函数处理。
5. 日志和配置模块，可以设置日志输出方式（标准io/文件），可以打印日志级别，用户名，时间，文件名和行数，日志内容。同时可以根据配置的json文件导入相关配置信息。
6. RPC实现，通过GOB编码，服务器能进程方法注册，客户端能向服务器传输参数并进行方法调用，然后传递返回值（两种实现：多路由和反射）。

# Getting Started

## Prerequisites

>HAHAGO需要Go 1.11及更新版本

## Install

```go
go get -u github.com/hahgraffe/hahago
```

## Usage Examples

详细的服务器实例和RPC实例都在[例子](https://github.com/hahgiraffe/hahago/tree/master/test)

## TODO
增加负载均衡与服务发现功能