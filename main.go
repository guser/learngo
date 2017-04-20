package main

import (
	"fmt"
)

// Stringer is a test interface
type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

func (f *fakeString) String() string {
	return f.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}

func main() {
	f1 := &fakeString{"this is a fake string"}
	printString(f1)

	f2 := "this is a real string"
	printString(f2)
}
