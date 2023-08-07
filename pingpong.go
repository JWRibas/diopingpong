package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) {
	for {
		c <- "PING"
		time.Sleep(time.Second * 1)
		c <- "PONG"
		time.Sleep(time.Second * 1)
	}
}

func imprimir(c chan string) {
	for msg := range c {
		fmt.Print(msg, " ")
	}
}

func main() {
	var c = make(chan string)

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
