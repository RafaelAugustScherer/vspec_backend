package main

import (
	"vspec_backend/src/domain"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/catalog/car/:id", func(ctx *gin.Context) {
		carro := domain.Car{
			ID:   1,
			Year: 2024,
			Model: domain.CarModel{
				ID:   1,
				Name: "Gol",
				Brand: domain.CarBrand{
					ID:   1,
					Name: "Volkswagen",
				},
			},
			Edition: "Trendline",
			Engine: domain.CarEngine{
				ID:            1,
				Code:          "EA-211",
				LiterCapacity: 1.6,
			},
		}
		ctx.JSON(200, gin.H{
			"car": carro.GetName(),
		})
	})
	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080
}
