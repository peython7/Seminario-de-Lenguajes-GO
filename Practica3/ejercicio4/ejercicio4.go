package main

import "fmt"

func pxng(enviar, recibe chan string, mensajes chan string, str string) {
	for i := 0; i < 5; i++ {
		<-recibe //Espera permiso antes de imprimir
		mensajes <- str //Envía el mensaje "PING" o "PONG"
		enviar <- "ok" //Da el permiso al otro
	}
}

func main() {
	mensajes := make(chan string) // canal de datos "PING"/"PONG"
	ping := make(chan string)  // canal para señal de permiso PING
	pong := make(chan string)  // canal para señal de permiso PONG

	//Creo las goroutines
	go pxng(pong, ping, mensajes, "PING")
	go pxng(ping, pong, mensajes, "PONG")

	ping <- "start" //Inicializa: da el primer permiso a PING

	for i := 0; i < 10; i++ {
		fmt.Println(<-mensajes)
	}
}
