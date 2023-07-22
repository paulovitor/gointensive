package main

import (
	"fmt"
	"time"
)

func processando() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// T1
func main() {
	// go processando() // T2
	// go processando() // T3
	// processando()    // T1

	canal := make(chan int) // cria um canal de comunicação

	go func() {
		// canal <- 1 // T2 - enche o canal <- valor
		for i := 0; i < 10; i++ {
			canal <- i
			fmt.Println("Jogou no canal", i)
		}
	}()

	// fmt.Println(<-canal) // esvazia o <-canal

	// time.Sleep(time.Second * 2)

	// for x := range canal {
	// 	fmt.Println(x)
	// 	fmt.Println("Recebeu do canal", x)
	// 	time.Sleep(time.Second)
	// }

	go worker(canal, 1)
	worker(canal, 2)
}

func worker(canal chan int, workerId int) { // recebe valores em um canal
	for { // loop infinito
		fmt.Println("Recebeu do canal", <-canal, "no worker", workerId)
		time.Sleep(time.Second)
	}
}
