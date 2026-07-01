package lexorank

import (
	"fmt"
	"testing"
)

func TestLexoRank(t *testing.T) {
	// Create a new generator with default settings
	generator := NewGenerator()

	// Generate an initial key
	key1, _ := generator.Between("", "")
	fmt.Println("Initial key:", key1)

	// Generate a key after key1
	key2, _ := generator.Next(key1)
	fmt.Println("Next key:", key2)

	// Generate a key before key1
	key0, _ := generator.Prev(key1)
	fmt.Println("Previous key:", key0)

	// Generate a key between key0 and key1
	keyMiddle, _ := generator.Between(key0, key1)
	fmt.Println("Middle key:", keyMiddle)
}
