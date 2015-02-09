package main

//import (
//	"fmt"
//	"net"
//	"os"
//)

//func main() {
//	if len(os.Args) != 2 {
//		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n E.g. xx.exe 255.255.255.1", os.Args[0])
//		os.Exit(1)
//	}
//	dotAddr := os.Args[1]
//	//return IP type
//	addr := net.ParseIP(dotAddr)

//	if addr == nil {
//		fmt.Println("Invalid address")
//		os.Exit(1)
//	}

//	//Class method signiture ?
//	//func (ip IP) DefaultMask() IPMask
//	mask := addr.DefaultMask()
//	network := addr.Mask(mask)
//	ones, bits := mask.Size()
//	fmt.Println(" Address is ", addr.String(),
//		"\n Default mask length is ", bits,
//		"\n Leading ones count is ", ones,
//		"\n Mask is (hex) ", mask.String(),
//		"\n Network is ", network.String())
//	os.Exit(0)
//}
