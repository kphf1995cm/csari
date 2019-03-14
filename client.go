package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("This is test0")
	con, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("err dialing:", err.Error())
		return
	}
	defer con.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		str, _ := inputReader.ReadString('\n')
		data := strings.Trim(str, "\n")
		if data == "quit" {
			return
		}
		_, err := con.Write([]byte(data))
		if err != nil {
			fmt.Println("send data error:", err)
			return
		}
		buf := make([]byte, 512)
		n, err := con.Read(buf)
		fmt.Println("from server:", string(buf[:n]))
	}
}
