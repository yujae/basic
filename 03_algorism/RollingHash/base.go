package RollingHash

/*
	Hash 특징
	1. 어떤 입력을 넣었을 때
	  출력은 입력에 대응하는 항상 일정한 값이 나온다
	2. 출력 값은 범위가 정해져 있다.
	ex) 나머지연상 %12 의 경우, 값의 범위는 0~11
	   13 % 12 = 1
	   24 % 12 = 0
	   15 % 12 = 3
	   25 % 12 = 1
	3. 역방향으로는 해석이 거의 불가능
	ex) x -> Hash(%12) -> 1 ... 13, 25 등 출력은 무한대가 될 수 있지만, 반대로는 해석 불가
	4. Hash 충돌 : 여러 입력값이 하나의 동일한 출력값을 갖는 것
	ex) 13 % 12 = 1, 25 % 12 = 1
	5. Hash 쓰임 : 공개키 암호화 방식, CheckSum, 블록체인 등

	Rolling Hash
	주어진 문자열 : a, b, c, d, e (각 자리는 s0, s1, s2, s3, s4)
	주어진 문자열의 Hash : Hi = (Hi-1 * A + si) % B
		H0 = (s0) %B
		H1 = (H0 * A + s1) % B
		H2 = (H1 * A + s2) % B
		H3 = (H2 * A + s3) % B
		H4 = (H3 * A + s4) % B
		abcde = H4
	참고 : 문자(s0~4)는 Ascii 0~255이므로
	    A는 그보다 큰 256, B는 소수로 하는 것이 좋음 (ex. 3571)
	    (소수로 나머지 연산을 할 경우 값의 분포가 넓게 퍼짐)
*/

const A = 256
const B = 3571

type HashTable struct {
	elements [B][]KeyValue
}

type KeyValue struct {
	key   string
	value string
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func Hash(s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}
	return h
}

func (t *HashTable) Add(key string, value string) {
	h := Hash(key)
	element := KeyValue{key: key, value: value}
	t.elements[h] = append(t.elements[h], element)
}

func (t *HashTable) Remove(key string) {
	h := Hash(key)
	for i, v := range t.elements[h] {
		if key == v.key {
			t.elements[h] = append(t.elements[h][:i], t.elements[h][i+1:]...)
			return
		}
	}
}

func (t *HashTable) Value(key string) string {
	h := Hash(key)
	for _, v := range t.elements[h] {
		if key == v.key {
			return v.value
		}
	}
	return ""
}
