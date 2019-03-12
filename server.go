package main

import (
	"fmt"
	"net"
)

func handle(con net.Conn) {
	defer con.Close()
	fmt.Println(con.RemoteAddr())
	for {
		buf := make([]byte, 180)
		n, err := con.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print("data size: %d msg: %s\n", n, string(buf[0:n]))
		msg := []byte("Hello World\n")
		con.Write(msg)
	}
}

func main() {
	fmt.Println("start server...")
	listen, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		fmt.Println("Listen failed! msg:", err)
		return
	}
	for {
		con, errs := listen.Accept()
		if errs != nil {
			fmt.Println("accept failed")
			continue
		}
		go handle(con)
	}

}
