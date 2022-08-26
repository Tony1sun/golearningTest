package main

import "time"

const (
	fmat = "2006-01-02 15:04:05"
)

func main() {
	c := make(chan int)
	BadSleep(1 * time.Second)

}
