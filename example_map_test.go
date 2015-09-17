package stringcache

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	a := &Map{}
	var b Getter
	b = a

	if b.Get("Bawb") != "Bawb" {
		t.Errorf("'%s' != %s", b.Get("Bawb"), "Bawb")
	}

	s := "Hello, "
	s += "World"
	s += "!"

	b.Get(s)

	l := a.Len()
	hw := "Hello, World!"
	b.Get(hw)
	if l != a.Len() {
		t.Errorf("Same string should not have grown cache ? expected %d, got %s", l, a.Len())
	}
}

func ExampleNewMap() {
	a := NewMap(2)
	a.Get("1")
	a.Get("2")
	fmt.Println(a.Len())
	a.Get("2")
	fmt.Println(a.Len())
	a.Get("3")
	fmt.Println(a.Len())

	// Output:
	// 2
	// 2
	// 1
}
