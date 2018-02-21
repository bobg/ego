package ego

import "strings"

func Example() {
	scope := NewScope()
	scope.Load(strings.NewReader(exampleSrc))
	scope.Eval(`printFirstNSquares(5)`)
	// Output:
	// 0
	// 1
	// 4
	// 9
	// 16
}

const exampleSrc = `
// Note: package declaration is ignored and can be omitted

import (
	"fmt"
	"math"
)

func mkSquaresChan() <-chan int {
	ch := make(chan int)
	go func() {
		for n := 0; ; n++ {
			ch <- n
		}
	}()
	return ch
}

func printFirstNSquares(n int) {
	var result []int
	ch := mkSquaresChan()
	for i := 0; i < n; i++ {
		fmt.Println(<-ch)
	}
}
`
