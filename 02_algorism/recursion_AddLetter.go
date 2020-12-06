package algorism

import "fmt"

func AddLetter(value string) {
	if len(value) == 3 {
		fmt.Println("Completed Value : ", value)
		return
	}
	for i := 0; i < 26; i++ {
		letter := string(97 + i)
		valueTemp := value + letter
		AddLetter(valueTemp)
	}
}
