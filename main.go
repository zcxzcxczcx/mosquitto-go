package main

import (
	"net"
	"fmt"
	"os"
)

func main(){
	if os.Args[1] == "server" {
		server, err := net.Listen("tcp", ":7777")
		if err == nil { //若果err为nil，则成功创建
			for {
				con, error_ := server.Accept()
				fmt.Println("有一个客户连接\n")
				if error_ == nil {
					con.Write([]byte("hello world!"))
				}
			}
		} else {
			return
		}
	} else {
		conn, err := net.Dial("tcp", "127.0.0.1:7777")
		if err == nil {
			for {
				buf := make([]byte, 1024)
				if length, err := conn.Read(buf); err == nil {
					if length > 0 {
						buf[length] = 0
						fmt.Printf("%s", string(buf[0:length]))
					}
				}
			}
		} else {
			return
		}
	}
}
