package integers

import (
	"fmt"
	"testing"
)

func TestAddr(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}

	//assertCorrectMessage(t, expected, sum)
}

/* func assertCorrectMessage(t testing.TB, expected, got int) {
	t.Helper()
	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
} */

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
