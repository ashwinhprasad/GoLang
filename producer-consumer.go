package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var elements [100]int

func main(){
	
	c := make(chan int,100)
	wg.Add(2)
	go producer(c)
	go consumer(c)
	wg.Wait()
	fmt.Printf("%v\n",elements)
}


// producer
func producer(c chan int) {
	for i:=1;i<=100;i++ {
		fmt.Printf("Producing element: %v\n",i)
		time.Sleep(3*time.Second)
		elements[i-1] = i
		c<-1
		fmt.Printf("Produced element: %v\n",i)
		
	}
	wg.Done()
}

// consumer
func consumer(c chan int) {
	var sleepTime time.Duration = 1
	for i:=1;i<=100;i++{

		// if (i%25==0 && sleepTime > 1){sleepTime = sleepTime -1}

		fmt.Printf("Yet to recieve element %v\n",i)
		<-c
		time.Sleep(sleepTime*time.Second)
		elements[i-1] = 0
		fmt.Printf("Recieved element %v\n",i)
		displayElements(elements)
	}
	wg.Done()
}

func displayElements(elements [100]int) {
	fmt.Println("\n\n####################")
	fmt.Println(elements)
	fmt.Printf("####################\n\n\n")
}