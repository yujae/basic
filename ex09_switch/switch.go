package main

import "fmt"

func main() {
	i := 3

	switch i { // 값을 판단할 변수 설정
	case 0: // 각 조건에 일치하는
		fmt.Println(0) // 코드를 실행합니다.
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3: // 3과 변수의 값이 일치하므로
		fmt.Println(3) // 이 부분을 실행하고 이후 실행을 중단
		fallthrough
	case 4:
		fmt.Println(4)
	default: // 모든 case에 해당하지 않을 때 실행
		fmt.Println(-1)
	}

	fmt.Println("----------------------")
	k := 8
	switch { // case에 조건식을 지정했으므로 판단할 변수는 생략
	case k >= 5 && k < 10: // i가 5보다 크거나 같으면서 10보다 작을 때
		fmt.Println("5 이상, 10 미만")
	case k >= 0 && k < 5: // i가 0보다 크거나 같으면서 5보다 작을 때
		fmt.Println("0 이상, 5 미만")
	}
}
