package main

import (
	"fmt"
	"message-service/kafka"
)

func main() {
	fmt.Println("Hello, World!")

	go kafka.StartKafka()
	fmt.Println("Running: StartKafka!")

}
