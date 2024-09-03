// ref: https://pkg.go.dev/context
package main

import (
	"context"
	"fmt"
)

func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("got the stop channel", n)
					return // returning not to leak the goroutine
				case dst <- n:
					fmt.Println(n, "qqq")
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		// fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// // https://leetcode.com/problems/two-sum/

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("got the stop channel")
// 				return
// 			default:
// 				fmt.Println("still working")
// 				time.Sleep(1 * time.Second)
// 			}
// 		}
// 	}()

// 	time.Sleep(5 * time.Second)
// 	fmt.Println("stop the gorutine")
// 	cancel()
// 	time.Sleep(5 * time.Second)
// }
