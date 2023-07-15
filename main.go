package main

import (
	"fmt"
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
	go func(){
		for x := range in{
			app = append(app, x)
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

	stage1 := SliceToCh(num)
	stage2 := AddNumberToSlice(stage1)
	for n := range stage2{
		fmt.Println(n)
	}
}