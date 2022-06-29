package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 关闭连接
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		// 读取数据
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("reader.Read error:", err)
			break
		}
		recvData := string(buf[:n])
		fmt.Println("receive data:", recvData)
		// 把数据再发给客户端
		conn.Write([]byte("ok"))
	}
}

func main() {
	// 开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	for {
		// 等待客户端来建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		// 启动一个单独的goroutine去处理连接
		go process(conn)
	}
}