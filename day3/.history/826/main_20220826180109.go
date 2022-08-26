package main

import "time"

const (
	fmat = "2006-01-02 15:04:05"
)

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		c <- 10
		close(c)
	}()
}
