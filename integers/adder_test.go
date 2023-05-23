package integers

import (
	"fmt"
	"testing"
)



func Add(x, y int) int {
	return x + y
}

func TestAdder(t *testing.T) {
	sum := Add(2,2)
	expected := 4

	assertCorrectMessage(t, sum, expected)
}

func assertCorrectMessage(t testing.TB, sum , expected int) {
	t.Helper()

	if sum != expected {
		t.Errorf("sum %v expected %v", sum, expected)
	}
}

func ExampleAdd() {
	sum := Add(1,5)
	fmt.Println(sum)
}