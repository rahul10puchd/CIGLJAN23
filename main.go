package main

// var wait sync.WaitGroup

// func main() {
// 	/*
// 		A WaitGroup waits for a collection of goroutines to finish.
// 		The main goroutine calls Add to set the number of goroutines to wait for.
// 		Then each of the goroutines runs and calls Done when finished.
// 		At the same time, Wait can be used to block until all goroutines have finished.
// 	*/

// 	wait.Add(3) //2 go routines
// 	for i := 0; i < 3; i++ {
// 		go hello(i)
// 	}
// 	//time.Sleep(30 * time.Second)
// 	wait.Wait()
// 	fmt.Println("Done with main function")
// }
// func hello(i int) {
// 	fmt.Println("Hello World -->" + fmt.Sprint(i))

// 	wait.Done()
// }

// shell - > bash, zsh,

// type Counter struct { //Lock, Unlock, Trylock
// 	sync.Mutex //all feature of Mutex, counter will get that:.composition
// 	Value      int
// }

/*
A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
A Mutex must not be copied after first use.
*/
// func main() {
// 	var wait sync.WaitGroup
// 	wait.Add(10)
// 	counter := Counter{}
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			counter.Lock()
// 			defer counter.Unlock()
// 			defer wait.Add(-1)
// 			counter.Value++

// 		}()
// 	}

// 	wait.Wait()
// 	fmt.Println(counter.Value)
// }
