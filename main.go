package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const dataBaseAddr = "localhost:6379"

func main(){
	conn, err := net.Dial("tcp", dataBaseAddr)
	if err != nil{
		fmt.Println("could not connect to database server")
		os.Exit(1)

	}
	cmdReader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Print(string(buf[:n]))
	str, err := cmdReader.ReadString('\n')
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	conn.Write([]byte(str + "\f"))
	for{
		bufCopy := make([]byte, len(buf))
		copy(buf, bufCopy)
		str, err := cmdReader.ReadString('\n')
		if err != nil{
			fmt.Println(err)
			os.Exit(1)
		}
		conn.Write([]byte(str + "\f"))
		n, _ = conn.Read(buf)
		fmt.Print(string(buf[:n]))
	}
}