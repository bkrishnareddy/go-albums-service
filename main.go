package main

import (
	"github.com/gin-gonic/gin"
	docs "github.com/go-albums-service/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var album_list = []album{
	{ID: "001", Title: "Blue Tarain", Artist: "John Cennin", Price: 109.56},
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /records/albums [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, album_list)
}

// @title Album Service
func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		records := v1.Group("/records")
		{
			records.GET("/albums", getAlbums)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:8000")
}
