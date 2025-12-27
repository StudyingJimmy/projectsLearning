package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// SocketServer 简单的socket服务器
type SocketServer struct {
	address string
}

// NewSocketServer 创建新的socket服务器
func NewSocketServer(address string) *SocketServer {
	return &SocketServer{address: address}
}

// Start 启动服务器
func (s *SocketServer) Start() error {
	// 监听TCP连接
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return fmt.Errorf("监听失败: %v", err)
	}
	defer listener.Close()

	fmt.Printf("服务器启动，监听地址: %s\n", s.address)
	fmt.Println("等待客户端连接...")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("接受连接失败: %v\n", err)
			continue
		}

		// 处理客户端连接
		go s.handleConnection(conn)
	}
}

// handleConnection 处理客户端连接
func (s *SocketServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("客户端连接: %s\n", clientAddr)

	// 创建读写器
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// 发送欢迎消息
	welcomeMsg := fmt.Sprintf("欢迎来到Socket服务器! 你的地址: %s\n", clientAddr)
	writer.WriteString(welcomeMsg)
	writer.Flush()

	// 循环读取客户端消息
	for {
		// 读取消息（以换行符为结束符）
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("客户端 %s 断开连接: %v\n", clientAddr, err)
			break
		}

		// 去除换行符
		message = strings.TrimSpace(message)
		fmt.Printf("收到来自 %s 的消息: %s\n", clientAddr, message)

		// 如果是退出命令
		if message == "quit" || message == "exit" {
			response := "服务器: 再见!\n"
			writer.WriteString(response)
			writer.Flush()
			break
		}

		// 回显消息给客户端
		response := fmt.Sprintf("服务器回复: %s\n", message)
		writer.WriteString(response)
		writer.Flush()
	}

	fmt.Printf("客户端 %s 已断开\n", clientAddr)
}

// SocketClient socket客户端
type SocketClient struct {
	address string
}

// NewSocketClient 创建新的socket客户端
func NewSocketClient(address string) *SocketClient {
	return &SocketClient{address: address}
}

// Connect 连接到服务器
func (c *SocketClient) Connect() error {
	// 连接到服务器
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return fmt.Errorf("连接服务器失败: %v", err)
	}
	defer conn.Close()

	fmt.Printf("已连接到服务器: %s\n", c.address)

	// 创建读写器
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	consoleReader := bufio.NewReader(os.Stdin)

	// 启动一个goroutine读取服务器消息
	go func() {
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("与服务器的连接断开: %v\n", err)
				os.Exit(0)
			}
			fmt.Print(message)
		}
	}()

	// 主循环：读取用户输入并发送给服务器
	for {
		fmt.Print("请输入消息 (输入quit退出): ")
		input, err := consoleReader.ReadString('\n')
		if err != nil {
			fmt.Printf("读取输入失败: %v\n", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "quit" || input == "exit" {
			writer.WriteString(input + "\n")
			writer.Flush()
			break
		}

		// 发送消息到服务器
		_, err = writer.WriteString(input + "\n")
		if err != nil {
			fmt.Printf("发送消息失败: %v\n", err)
			break
		}
		writer.Flush()

		// 小延迟，避免消息混乱
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

// 演示Socket的基本概念
func demonstrateSocketBasics() {
	fmt.Println("=== Socket基础概念演示 ===")
	fmt.Println()
	
	// 1. 创建监听socket
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Printf("创建监听socket失败: %v\n", err)
		return
	}
	defer listener.Close()

	// 获取实际监听的地址
	addr := listener.Addr().String()
	fmt.Printf("1. 监听Socket创建成功，地址: %s\n", addr)
	fmt.Println("   - Socket是操作系统提供的网络通信端点")
	fmt.Println("   - 包含IP地址和端口号")
	fmt.Println()

	// 2. 演示连接过程
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		defer conn.Close()
		
		fmt.Printf("2. 接受连接: %s\n", conn.RemoteAddr())
		fmt.Println("   - Accept()返回一个新的socket用于通信")
		fmt.Println("   - 原监听socket继续监听新的连接")
		fmt.Println()

		// 读取数据
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		
		fmt.Printf("3. 收到数据: %s\n", string(buffer[:n]))
		fmt.Println("   - Read()从socket读取数据")
		fmt.Println("   - 数据通过内核缓冲区传输")
		fmt.Println()

		// 发送响应
		response := "Hello from server!"
		conn.Write([]byte(response))
		fmt.Printf("4. 发送响应: %s\n", response)
		fmt.Println("   - Write()向socket写入数据")
		fmt.Println("   - 数据通过TCP协议传输")
	}()

	// 客户端连接
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Printf("客户端连接失败: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Printf("2. 客户端连接到服务器: %s\n", addr)
	fmt.Println("   - Dial()创建客户端socket")
	fmt.Println("   - 通过三次握手建立连接")
	fmt.Println()

	// 发送数据
	message := "Hello from client!"
	conn.Write([]byte(message))
	fmt.Printf("3. 客户端发送数据: %s\n", message)
	fmt.Println()

	// 读取响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return
	}
	
	fmt.Printf("4. 客户端收到响应: %s\n", string(buffer[:n]))
	fmt.Println()

	time.Sleep(1 * time.Second)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法:")
		fmt.Println("  go run useSocket.go server  - 启动服务器")
		fmt.Println("  go run useSocket.go client  - 启动客户端")
		fmt.Println("  go run useSocket.go demo    - 演示Socket基础概念")
		fmt.Println()
		fmt.Println("服务器默认监听: localhost:8080")
		return
	}

	switch os.Args[1] {
	case "server":
		server := NewSocketServer("localhost:8080")
		if err := server.Start(); err != nil {
			fmt.Printf("服务器启动失败: %v\n", err)
		}

	case "client":
		client := NewSocketClient("localhost:8080")
		if err := client.Connect(); err != nil {
			fmt.Printf("客户端连接失败: %v\n", err)
		}

	case "demo":
		demonstrateSocketBasics()

	default:
		fmt.Printf("未知命令: %s\n", os.Args[1])
		fmt.Println("请使用: server, client 或 demo")
	}
}