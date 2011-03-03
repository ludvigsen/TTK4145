package main
import (
	"fmt"
	"time"
)

var ch = make(chan int)
var srvInt int = 0

func server(){
	for ;true; {
		temp := <- ch //wait for int to add to srvInt
		srvInt+=temp 
		fmt.Printf("Int has changed, it is now: %d\n", srvInt)
	}
}

func client(){
	for ;true; {
		ch <- 3 //Put the int to add in channel
	    time.Sleep(1000000000)
	}
}

func main(){
	go server()
	for i := 0;i<4;i++{
		go client()
	}
	never := make (chan int)		
	<-never
}
