package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 与服务器连接
	conn ,err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Print("dial failed, err:%v\n", err)
		return
	}
	// 关闭连接
	defer  conn.Close()
	// 输入数据
	inputReader := bufio.NewReader(os.Stdin)
	for {
		// 读取用户输入
		input, _ := inputReader.ReadString('\n')
		// 截断
		inputInfo := strings.Trim(input, "\r\n")
		// 读取到用户输入q或者Q就退出
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		// 把输入的数据发送给服务端
		_, err := conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		// 从服务端接收回复的消息
		buf := [256]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("conn.Read error :", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
