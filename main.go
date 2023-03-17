package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	docs "github.com/go-albums-service/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}


var album_list = []album{
	{ID:"001",Title: "Blue Tarain", Artist: "John Cennin",Price: 109.56},
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
// @Router /albums [get]
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK,album_list)
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8000")
}