/* LookupHost
 */
package main

//import (
//	"fmt"
//	"net"
//	"os"
//)

//func main() {
//	/*
//		xxx.exe www.taobao.com

//		> 210.176.46.41
//		> 210.176.46.51

//      cname is www.taobao.com
//	*/
//	if len(os.Args) != 2 {
//		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
//		os.Exit(1)
//	}
//	name := os.Args[1]
//	//signiture:
//	//func LookupHost(name string) (cname string, addrs []string, err os.Error)
//	addrs, err := net.LookupHost(name)
//	if err != nil {
//		fmt.Println("Error: ", err.Error())
//		os.Exit(2)
//	}
//	for _, s := range addrs {
//		fmt.Println(s)
//	}
//	os.Exit(0)
//}
