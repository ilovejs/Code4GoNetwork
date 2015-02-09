/* GetHeadInfo
 */
package main

//import (
//	"fmt"
//	"io/ioutil"
//	"net"
//	"os"
//)

//func main() {
//	/*
//		GetHeadInfo www.google.com:80
//	*/
//	if len(os.Args) != 2 {
//		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
//		os.Exit(1)
//	}
//	service := os.Args[1]
//	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
//	//syntax error in the address specified
//	checkError(err)

//	//func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
//	conn, err := net.DialTCP("tcp", nil, tcpAddr)
//	//attempt to connect to remote service may fail
//	checkError(err)

//	//coerce into byte array
//	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
//	//Write fails if connection died suddenly / network time out
//	checkError(err)

//	//result, err := readFully(conn)
//	result, err := ioutil.ReadAll(conn)
//	//read may fail
//	checkError(err)

//	fmt.Println(string(result))
//	os.Exit(0)
//}

////idea is good
//func checkError(err error) {
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
//		os.Exit(1)
//	}
//}
