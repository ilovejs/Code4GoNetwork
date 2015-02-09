package main

//http://www.zhihu.com/question/22528008

//import (
//	"log"
//	"runtime"
//	"time"
//)

//func main() {
//	runtime.GOMAXPROCS(1)
//	for i := 0; i < 10; i++ {
//		go func() {
//			//runtime.LockOSThread()
//			for {
//				log.Println("a")
//				time.Sleep(time.Second)
//			}
//		}()
//	}
//	time.Sleep(time.Hour)
//}
