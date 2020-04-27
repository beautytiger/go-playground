package main
import (
	//"runtime"
	"sync"
	"fmt"
)
func main() {
	//runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for j := 0; j < 10; j++ {
		go func() {
			fmt.Println("j: ", j)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
