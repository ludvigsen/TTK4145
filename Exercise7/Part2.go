package main
import (
	"fmt"
//	"sync"	
)

var ch channel
var srvInt int = 0

func server(newint channel int){
	srvInt <- newint
	fmt.Printf("Int has changed, it is now: %d", srvInt)
}

func client(ch channel int){
	while(true){
		var temp int <- ch
		ch <- temp+3
	    sleep(1)
	}
}

func main(){
	go server(channel)
	go client(channel)
	go client(channel)	
}
