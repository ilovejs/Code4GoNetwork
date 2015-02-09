/*
 Simple "Echo" Server

 1.	Single-threaded.
 2. While a client has a connection open to it, no other cllient can connect.
    Other clients are blocked, and will probably time out.
*/
package main

//import (
//	"fmt"
//	"net"
//	"os"
//)

//func main() {
//	service := ":1201"
//	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
//	checkError(err)

//	listener, err := net.ListenTCP("tcp", tcpAddr)
//	checkError(err)

//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			continue
//		}

//		//core
//		handleClient(conn)
//		conn.Close() // we're finished
//	}
//}
//func handleClient(conn net.Conn) {
//	var buf [512]byte
//	//why wrap in for ???
//	for {
//		n, err := conn.Read(buf[0:])
//		if err != nil {
//			return
//		}
//		//print everything
//		//Tips: string()
//		fmt.Println(string(buf[0:]))

//		//write back with n ??
//		_, err2 := conn.Write(buf[0:n])
//		if err2 != nil {
//			//
//			return
//		}
//	}
//}
//func checkError(err error) {
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
//		os.Exit(1)
//	}
//}
