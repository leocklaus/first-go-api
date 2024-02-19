package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/cars", handleGetCars)

	v1.GET("/cars/:id", handleGetCarById)

	v1.POST("/cars", handleSaveCar)

	v1.DELETE("/cars/:id", handleDeleteCar)

	router.Run()

}

func handleGetCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func handleGetCarById(c *gin.Context) {
	id := c.Param("id")

	for _, car := range cars {
		if car.ID == id {
			c.IndentedJSON(http.StatusOK, car)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Car not found"})
}

func handleSaveCar(c *gin.Context) {
	var newCar Car

	if err := c.BindJSON(&newCar); err != nil {
		return
	}

	cars = append(cars, newCar)

	c.IndentedJSON(http.StatusCreated, newCar)
}

func handleDeleteCar(c *gin.Context) {
	id := c.Param("id")

	for i, car := range cars {
		if (car.ID) == id {
			cars = deleteElement(cars, i)
			c.Status(http.StatusNoContent)
			return
		}

	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message:": "Car Not Found"})

}

func deleteElement(slice []Car, index int) []Car {
	return append(slice[:index], slice[index+1:]...)
}

type Car struct {
	ID    string  `json:"id"`
	Brand string  `json:"brand"`
	Model string  `json:"model"`
	Price float32 `json:"price"`
}

var cars = []Car{
	{ID: "1", Brand: "VW", Model: "Gol", Price: 32.000},
	{ID: "2", Brand: "Ford", Model: "Fiesta", Price: 39.000},
	{ID: "3", Brand: "GM", Model: "Onix", Price: 45.000},
}
