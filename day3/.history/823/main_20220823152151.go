package main

func main() {
	c := make(chan int)
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
	}
}
