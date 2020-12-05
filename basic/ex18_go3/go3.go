package main

import (
	"fmt"
	"runtime"
)

func main() {
	var str = "Hello Goorm!"
	done := make(chan bool)

	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정
	fmt.Println(runtime.GOMAXPROCS(0))   // 설정 값 출력

	go func() {
		println("inside go")
		for i := 0; i < 10; i++ {
			fmt.Println(str, i)
		}

		done <- true //채널에 true를 송신함. 🤯버퍼 여러개의 경우 다 차면 대기함
	}()

	println("inside main")
	<-done //수신함으로써 대기를 끝냄
	println("finish")
}
