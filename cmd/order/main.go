package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // foi preciso add _ porque não tem uma referência direta ao import
	"github.com/paulovitor/gointensive/internal/entity"
	"github.com/paulovitor/gointensive/internal/infra/database"
	"github.com/paulovitor/gointensive/internal/usecase"
)

type Car struct { // similar a uma classe
	Model string
	Color string
}

var x string // declara somente

// método: função referente a struct
func (c Car) Start() {
	println(c.Model + " has been started")
}

// função comum
func soma(x, y int) int { // quando os 2 parâmetros são iguais não precisa repetir o tipo
	return x + y
}

// * ponteiro - aponta para o endereço onde determinado valor está na memória
func (c *Car) ChangeColor(color string) {
	c.Color = color // duplicando o valor de c.Color na memória - cópia do color original
	println("New color: " + c.Color)
}

func main() {
	order, err := entity.NewOrder("1", 10, -1)
	if err != nil {
		println(err.Error())
		// println(order.ID) // nil pointer dereference
	} else {
		println(order.ID)
	}

	a := 10
	// b := a // copiou o valor de a, mas criou o seu próprio espaço na memória
	b := &a // b está apontando para o valor na memória de a
	*b = 20 // esta pedindo que seja alterado o valor na memória quando usa-se o ponteiro *

	println(&a) // & é o endereço da variável a na memória
	println(b)
	println(a)

	car := Car{ // := está declarando e atribuindo a variável
		Model: "Ferrari",
		Color: "Red",
	}

	car.Model = "Fiat" // altera o valor da variável

	println(car.Model)

	println("Hello, World!")

	car.Start()

	car.ChangeColor("Blue")
	println(car.Color)

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close() // espera tudo rodar e depois executa o close
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	input := usecase.OrderInput{
		ID:    "1234",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
