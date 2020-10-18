package iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	t.Run("character repeater repeats 5 times", func(t *testing.T) {
		got := CharRepeater("A", 5)
		want := "AAAAA"
		if got != want {
			t.Errorf("got %s wanted %s", got, want)
		}
	})
}

func BenchmarkRepeatChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CharRepeater("A", 5)
	}
}

func ExampleCharRepeater() {
	res := CharRepeater("a", 5)
	fmt.Println(res)
	// Output: "aaaaa"
}
