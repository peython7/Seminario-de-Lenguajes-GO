package main

import "fmt"

func pxng(enviar, recibe chan string, mensajes chan string, str string) {
	for i := 0; i < 5; i++ {
		<-recibe
		mensajes <- str
		enviar <- "ok"
	}
}

func main() {
	mensajes := make(chan string)
	ping := make(chan string)
	pong := make(chan string)

	go pxng(pong, ping, mensajes, "PING")
	go pxng(ping, pong, mensajes, "PONG")

	ping <- "start"

	for i := 0; i < 10; i++ {
		fmt.Println(<-mensajes)
	}
}
