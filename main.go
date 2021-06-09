package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("hello, world")
	// for idx, arg := range os.Args {
	// 	// fmt.Println("参数", idx, ":", arg)
	// 	fmt.Printf("参数%d: %s\n", idx, arg)
	// }
	var userName string
	var passwd string
	var host string
	var port int

	flag.StringVar(&userName, "u", "", "用户名， 默认为空")
	flag.StringVar(&passwd, "p", "", "密码，默认为空")
	flag.StringVar(&host, "h", "127.0.0.1", "主机，默认为localhost")
	flag.IntVar(&port, "P", 6666, "端口号，默认为6666")

	flag.Parse()

	fmt.Printf("userName: %v, passwd: %v, host: %v, port: %v\n", userName, passwd, host, port)
}
