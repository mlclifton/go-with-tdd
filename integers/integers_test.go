package integers

import (
	"fmt"
	"testing"
)

func TestIntegers(t *testing.T) {
	t.Run("adding two numbers", func(t *testing.T) {
		got := Add(2, 2)
		want := 4
		if got != want {
			t.Errorf("wanted %d got %d", got, want)
		}
	})

	t.Run("subtracting one int from another", func(t *testing.T) {
		got := Subtract(5, 2)
		want := 3
		if got != want {
			t.Errorf("wanted %d got %d", got, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
