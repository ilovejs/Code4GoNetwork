// go-network project main.go
package main

//import (
//	"fmt"
//	"net"
//	"os"
//)

//func main() {
//	if len(os.Args) != 2 {
//		//Args[0] is programe exe name
//		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
//		//non-zero means an error
//		os.Exit(1)
//	}
//	name := os.Args[1]
//	addr := net.ParseIP(name)

//	if addr == nil {
//		fmt.Println("Invalid address")
//	} else {
//		fmt.Println("The address is ", addr.String())
//	}
//	//0 means success
//	os.Exit(0)
//}
