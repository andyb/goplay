//A simple example to try and replicate this Akka example - http://blog.knoldus.com/2013/01/12/akka-futures-in-scala-with-a-simple-example/

package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Started sync")
	startSync := time.Now()

	a := TimeTakingIdentityFunction(1)
	b := TimeTakingIdentityFunction(2)
	c := TimeTakingIdentityFunction(3)
	log.Println(a)
	log.Println(b)
	log.Println(c)
	log.Printf("Sync Finished and took %v", time.Since(startSync))

	log.Println("Started async")
	startAsync := time.Now()

	a1 := TimeTakingIdentityFunctionAsync(1)
	b1 := TimeTakingIdentityFunctionAsync(2)
	c1 := TimeTakingIdentityFunctionAsync(3)

	log.Println(<-a1)
	log.Println(<-b1)
	log.Println(<-c1)
	log.Printf("Async Finished and took %v", time.Since(startAsync))
}

func TimeTakingIdentityFunction(num int) int {
	time.Sleep(3 * time.Second)
	return num
}

func TimeTakingIdentityFunctionAsync(num int) chan int {
	c := make(chan int)
	go func() {
		c <- TimeTakingIdentityFunction(num)
	}()
	return c
}
