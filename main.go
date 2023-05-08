package main

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo"
	// pacote nao esta sendo usado diretamente, esta sendo utilizado por outro pacote
	// _ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// slice++
var cars []Car

func generateCars() {
	cars = append(cars, Car{Name: "Ferrari", Price: 100})
	cars = append(cars, Car{Name: "BMW", Price: 999})
	cars = append(cars, Car{Name: "Porshe", Price: 400})
}

func (c Car) Andar() {
	fmt.Println("O carro", c.Name, "está andando")
}

// subindo um webserver e chamando a função getCars
func main() {
	generateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":7557"))
}

// GET
func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

// POST
func createCar(c echo.Context) error {
	car := new(Car)
	//faz o Bind do quie eu to recebend no paramentro do post com os campos da struct
	//se bind ser certo, vai dar um append e vai retornar 200 - OK

	if err := c.Bind(car); err != nil {
		return err
	}
	//passando o local de memória aonde eu quero dar o append com *ponteiro
	cars = append(cars, *car)
	//vai executar e guadar o car no banco
	saveCar(*car)
	return c.JSON(200, cars)
}

func saveCar(car Car) error {
	// abrir conexao no banco
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}

	//enquanto nao terminar a execução, o db não é executado
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, price) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	// _ ignora o resultado pois nao vamos usar
	_, err = stmt.Exec(car.Name, car.Price)
	if err != nil {
		return err
	}
	return nil
}

//PUT

//DELETE
