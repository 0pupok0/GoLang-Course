package main

import (
	"fmt"
	"lizaCourse/dataStructs"
)

func main() {
	l := dataStructs.New(2)
	l.Println()

	l.Add(1)
	l.Add(2)
	l.Println()

	err := l.Pop()
	if err != nil {
		fmt.Println(err)
		return
	}
	l.Println()

	val, err := l.At(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)
	val, err = l.At(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)

	fmt.Println(l.Size())

	err = l.DeleteFrom(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	l.Println()

	err = l.UpdateAt(0, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	l.Println()

	l2 := dataStructs.NewFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	l2.Println()
}
