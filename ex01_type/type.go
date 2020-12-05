package main

import (
	"fmt"
)

/*
uint8		부호 없는(unsigned) 8비트, 1바이트 정수	0 ~ 255
uint16		부호 없는 16비트, 2바이트 정수	0 ~ 65535
uint32		부호 없는 32비트, 4바이트 정수	0 ~ 4294967295
uint64		부호 없는 64비트, 8바이트 정수	0 ~ 18446744073709551615
int8		부호 있는(signed) 8비트, 1바이트 정수	-128 ~ 127
int16		부호 있는 16비트, 2바이트 정수	-32768 ~ 32767
int32		부호 있는 32비트, 4바이트 정수	-2147483648 ~ 2147483647
int64		부호 있는 64비트, 8바이트 정수	-9223372036854775808 ~ 9223372036854775807
uint		32비트 시스템에서는 uint32, 64비트 시스템에서는 uint64
int			32비트 시스템에서는 int32, 64비트 시스템에서는 int64
uintptr		uint와 크기가 동일하며 포인터를 저장할 때 사용
float32		IEEE-754 32비트 단정밀도 부동소수점, 7자리 정밀도 보장
float64		IEEE-754 64비트 배정밀도 부동소수점, 15자리 정밀도 보장
complex64	float32 크기의 실수부와 허수부로 된 복소수
complex128	float64 크기의 실수부와 허수부로 된 복소수
byte		uint8과 크기가 동일, 바이트 단위로 저장할 때 사용.
			byte에는 보통 16진수, 문자 값으로 저장합니다.
			실무에서는 바이너리 파일에서 데이터를 읽고 쓰거나, 데이터를 암호화, 복호화할 때 주로 사용
rune		int32와 크기가 동일, 유니코드 문자 코드(Code point)를 저장할 때 사용
*/

func main() {
	fmt.Println("------- 정수 ------")
	var int1 int = 32
	var int2 int = -15
	var int3 int = 0723       // 8진수로 저장
	var int4 int = 0x32fa2c75 // 16진수로 저장
	fmt.Println(int1)
	fmt.Println(int2)
	fmt.Println(int3)
	fmt.Println(int4)
	fmt.Println("------- 실수 ------")
	// 소수점 사용
	var float1 float32 = 0.1
	var float2 float32 = .35
	var float3 float32 = 132.73287
	// 지수 표기법
	var float4 float32 = 1e7
	var float5 float64 = .12345e+2
	var float6 float64 = 5.32521e-10
	fmt.Println(float1)
	fmt.Println(float2)
	fmt.Println(float3)
	fmt.Println(float4)
	fmt.Println(float5)
	fmt.Println(float6)

	fmt.Println("------- 실수의 주의사항 ------")
	// 실수에서 주의 사항
	var a float64 = 10.0
	for i := 0; i < 10; i++ {
		a = a - 0.1
	}
	fmt.Println(a)        // 9.000000000000004: 반올림 오차 발생
	fmt.Println(a == 9.0) // false: 실수는 ==로 비교하면 안됨

	fmt.Println("------ byte -------")
	var byte1 byte = 10   // 10진수 저장
	var byte2 byte = 0x32 // 16진수 저장
	var byte3 byte = 'a'  // 문자 저장
	fmt.Println(byte1)
	fmt.Println(byte2)
	fmt.Println(byte3)

	fmt.Println("------ rune -------")
	var rune1 rune = '한'
	var rune2 rune = '\ud55c'     // 한
	var rune3 rune = '\U0000d55c' // 한
	fmt.Println(string(rune1))
	fmt.Println(string(rune2))
	fmt.Println(string(rune3))
}
