package main

import "github.com/lesser-P/learnMap/demo"

func main() {
	//group := sync.WaitGroup{}
	//group.Add(1)
	//defer group.Done()
	//i := demo.DeferFunc()
	//i2 := demo.DeferFuncNameBack()
	//fmt.Println(i)
	//fmt.Println(i2)
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go demo.ParentGoroutine(&wg)
	//wg.Wait()
	//fmt.Println("主协程结束")
	//demo.PrintTenWord()
	//group.Wait()
	//bools := make(chan bool)
	//strings := make(chan string)
	//go demo.SelectDemo(bools, strings)
	//for {
	//	go func() {
	//		bools <- true
	//	}()
	//	time.Sleep(2 * time.Second)
	//	go func() {
	//		strings <- "a"
	//		time.Sleep(time.Second)
	//	}()
	//}
	demo.SelectDemo1()
}
