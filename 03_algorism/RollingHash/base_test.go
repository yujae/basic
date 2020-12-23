package RollingHash

import (
	"fmt"
	"testing"
)

func TestRollingHash(t *testing.T) {
	table := NewHashTable()

	table.Add("AAA", "010101")
	table.Add("AAABBBC", "010101")
	table.Add("AABBBC", "0101021")

	table.Remove("AABBBC")
	for i, v := range table.elements {
		fmt.Println(i, v)
	}

	fmt.Println("AABBBC = ", table.Value("AABBBC"))
}
