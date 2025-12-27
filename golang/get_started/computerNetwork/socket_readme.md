# Socket 网络编程基础

## 概述

Socket（套接字）是计算机网络编程的基础，提供了不同主机间进程通信的端点。它是应用层与TCP/IP协议族通信的中间软件抽象层，是应用程序通过网络协议进行通信的接口。

## Socket 基本概念

### 1. 什么是Socket

Socket是操作系统提供的网络通信端点，包含以下要素：
- **IP地址**：标识网络中的主机
- **端口号**：标识主机中的进程（0-65535）
- **协议**：TCP、UDP等传输层协议

### 2. Socket类型

- **流式Socket（SOCK_STREAM）**：基于TCP协议，提供可靠的、面向连接的字节流服务
- **数据报Socket（SOCK_DGRAM）**：基于UDP协议，提供无连接的、不可靠的数据报服务
- **原始Socket（SOCK_RAW）**：允许直接访问底层协议，如IP、ICMP

## TCP Socket通信流程

### 服务器端流程

1. **创建Socket**：调用`socket()`函数创建套接字
2. **绑定地址**：调用`bind()`函数将套接字绑定到特定地址和端口
3. **监听连接**：调用`listen()`函数开始监听传入连接
4. **接受连接**：调用`accept()`函数接受客户端连接
5. **数据传输**：通过`read()/write()`或`send()/recv()`进行数据交换
6. **关闭连接**：调用`close()`关闭套接字

### 客户端流程

1. **创建Socket**：调用`socket()`函数创建套接字
2. **连接服务器**：调用`connect()`函数连接到服务器
3. **数据传输**：通过`read()/write()`或`send()/recv()`进行数据交换
4. **关闭连接**：调用`close()`关闭套接字

## Go语言中的Socket编程

Go语言的`net`包提供了简洁的Socket编程接口：

### 服务器端

```go
// 创建监听
listener, err := net.Listen("tcp", "localhost:8080")
if err != nil {
    log.Fatal(err)
}
defer listener.Close()

// 接受连接
conn, err := listener.Accept()
if err != nil {
    log.Fatal(err)
}
defer conn.Close()

// 读取数据
buffer := make([]byte, 1024)
n, err := conn.Read(buffer)

// 发送数据
conn.Write([]byte("Hello, Client!"))
```

### 客户端

```go
// 连接服务器
conn, err := net.Dial("tcp", "localhost:8080")
if err != nil {
    log.Fatal(err)
}
defer conn.Close()

// 发送数据
conn.Write([]byte("Hello, Server!"))

// 读取响应
buffer := make([]byte, 1024)
n, err := conn.Read(buffer)
```

## 运行示例程序

### 1. 演示Socket基础概念

```bash
go run useSocket.go demo
```

这个命令会演示Socket的基本工作流程，包括：
- 创建监听Socket
- 客户端连接过程
- 数据发送和接收
- 连接关闭

### 2. 启动服务器

在一个终端窗口中运行：

```bash
go run useSocket.go server
```

服务器会监听`localhost:8080`，等待客户端连接。

### 3. 启动客户端

在另一个终端窗口中运行：

```bash
go run useSocket.go client
```

客户端会连接到服务器，你可以输入消息与服务器交互。

## Socket与TCP的关系

Socket是TCP协议的具体实现接口：

1. **TCP是协议**：定义了数据传输的规则和格式
2. **Socket是接口**：提供了使用TCP协议进行编程的API
3. **TCP基于Socket实现**：TCP连接通过Socket建立和维护
4. **Socket屏蔽底层细节**：开发者无需关心TCP协议的具体实现

### TCP三次握手在Socket中的体现

1. **客户端connect()**：发送SYN包
2. **服务器accept()**：接收SYN，发送SYN+ACK
3. **客户端connect()返回**：接收SYN+ACK，发送ACK
4. **连接建立**：双方可以开始数据传输

### TCP四次挥手在Socket中的体现

1. **一方close()**：发送FIN包
2. **对方read()返回0**：接收到FIN
3. **对方close()**：发送FIN包
4. **最初一方close()**：接收到FIN，连接完全关闭

## 关键概念总结

1. **Socket是端点**：每个Socket代表网络通信的一个端点
2. **TCP是可靠的**：通过序列号、确认应答、重传机制保证数据可靠传输
3. **Socket是双向的**：建立连接后，双方都可以发送和接收数据
4. **字节流服务**：TCP提供无结构的字节流，应用层需要处理消息边界
5. **全双工通信**：连接双方可以同时发送和接收数据

## 后续学习建议

1. **理解TCP协议**：深入学习TCP的可靠传输机制
2. **并发处理**：学习如何使用goroutine处理多个客户端连接
3. **协议设计**：设计应用层协议来处理消息边界
4. **性能优化**：学习Socket编程的性能优化技巧
5. **网络安全**：了解网络安全在Socket编程中的应用

这个程序为你理解TCP如何基于Socket连接打下了基础。接下来你可以深入学习TCP协议的具体实现和优化技巧。