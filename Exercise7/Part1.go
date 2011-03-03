package main
import (
	   "fmt"
)

/* Hangs if this number gets bigger */
const NGR = 10000

type Point struct {
	 x, y int
}
func test(){
	 var strings = [4]string{"Dette", "er", "en", "test!"}	 
	 for i := 0;i<4;i++{
	 	 fmt.Printf("%s ", strings[i])
	 }
	 fmt.Printf("\n")
}

func test1(i chan int){
	 fmt.Printf("Goroutine: %d\n",<-i)
}

func main() {
	 ci := make(chan int, NGR)	 
	 for i:=0;i<NGR;i++ {
	 	 go test()
	 	 ci<-i
	 	 go test1(ci)
	 }
	 const s = "Hello"
	 world := "World"
	 var p Point
	 p.x = 12
	 p.y = 13
	 fmt.Printf("%s %s\n",s,world) 
	 fmt.Printf("X: %d, Y: %d\n",p.x,p.y)
//	 fmt.Printf("%v\n",strings)
}