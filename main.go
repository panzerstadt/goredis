package main

import "fmt"

func main() {
	c := make(chan int)
	go serve_redis(c)
	go serve_http(c)

	out := <-c
	fmt.Println(out)
}
