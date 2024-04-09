package demo

import (
	"fmt"
	"sync"
	"time"
)

func ParentGoroutine(wg *sync.WaitGroup) {
	fmt.Println("父协程开始")
	go children(wg)
	fmt.Println("父协程结束")
}

func children(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("子协程结束")
	wg.Done()
}

// 开启10个协程分别打印1-10，按照顺序输出
func PrintTenWord() {
	wg := sync.WaitGroup{}
	chans := make([]chan bool, 10)
	// 初始化10个chan
	for i, _ := range chans {
		chans[i] = make(chan bool)
	}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(ch chan bool, i int) {
			defer wg.Done()
			// 等待信号传入
			<-ch
			fmt.Println(i + 1)
			if i < 9 {
				// 给下一个传
				chans[i+1] <- true
			}
		}(chans[i], i)
	}
	// 存入信号，开始
	chans[0] <- true
}

func SelectDemo(ch1 chan bool, ch2 chan string) {
	for true {
		select {
		case <-ch1:
			fmt.Println("is bool")
		case <-ch2:
			fmt.Println("is string")
		}
	}
}
func SelectDemo1() {
	bools := make(chan bool)
	times := time.After(time.Second * 1)
	select {
	case data := <-bools:
		fmt.Println(data)
	case <-times:
		fmt.Println("超时")
	}
}
