package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "test-value") //write
	ctx = context.WithValue(ctx, "key", "test1-value1")
	fmt.Println(ctx.Value("key")) //read

}

//gin engine -> middleware(populate gin context) -> handler(read gin context)

// func addValue(ctx context.Context) context.Context {
// 	return context.WithValue(ctx, "key", "test-value")
// }

// func readValue(ctx context.Context) {
// 	val := ctx.Value("key")
// 	fmt.Println(val)
// }
// func slowOperationWithTimeout(ctx context.Context) (Result, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
// 	defer cancel() // releases resources if slowOperation completes before timeout elapses
// 	return slowOperation(ctx)
// }

func doSomething(ctx context.Context) {
	fmt.Println("doSomething")
}
func mainmmm() {
	ctx := context.TODO()
	/*
		TODO returns a non-nil, empty Context.
		Code should use context.TODO when it's unclear
		 which Context to use or it is not yet available
		  (because the surrounding function has not yet been extended to accept a
			Context parameter).
	*/
	context.Background()
	/*
		Background returns a non-nil, empty Context.
		It is never canceled, has no values, and has no deadline.
		 It is typically used by the main function, initialization, and tests,
		 and as the top-level Context for incoming requests.
	*/
	doSomething(ctx)
	var number Number
	number = 10
	var ns NumberSystem
	ns = number
	ns.Print()
	ns.Square()
	number.Print()
	number.Square()
}

type Number int

func (n Number) Print() {
	fmt.Println(n)
}
func (n Number) Square() Number {
	// result := int(n) * int(n)
	return n * n
}

type NumberSystem interface {
	Print()
	Square() Number
}
