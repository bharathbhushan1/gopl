package main

import (
	"flag"
	"fmt"
	"time"
)

var sleepDuration = flag.Duration("sleep", 5*time.Second, "Sleep duration")

func main() {
	flag.Parse()
	fmt.Printf("SLEEPING FOR %v\n", *sleepDuration)
	time.Sleep(*sleepDuration)
	fmt.Println("DONE WITH SLEEP")
}
