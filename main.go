package main

import (
	"fmt"
	"sync"
)
func SliceToCh (num []int) <-chan int{
	out := make(chan int)
	go func(){
		for _,x := range num{
			out <- x
		}
		close(out)
		}()
		return out
}

func AddNumberToSlice (in <-chan int) <-chan int{
	out := make(chan int)
	app := []int{}
	var mu sync.Mutex
	go func(){
		for x := range in{
			mu.Lock()
			app = append(app, x)
			mu.Unlock()
		}
		app = append(app, app...)
		for _,y := range app{
				out <- y
		}
		close(out)
	}()
	return out
}

func main(){
	num := []int{1, 2, 3, 4, 5}
	var wait sync.WaitGroup

	stage1 := SliceToCh(num)
	stage2 := AddNumberToSlice(stage1)

	wait.Add(2)
	for n := range stage2{
		defer wait.Done()
		fmt.Println(n)
	}

	wait.Wait()	
}
