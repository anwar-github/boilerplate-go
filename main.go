package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// instance router Gin
	r := gin.Default()

	// handler dan rute
	r.GET("/", func(c *gin.Context) {

	})

	asset := r.Group("asset")
	{
		asset.GET("/", handleAsset)
		asset.GET("/detail", handleAssetDetail)
	}

	trx := r.Group("asset")
	{
		trx.POST("/", handleTransaction)
	}

	// run server
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

func handleAsset(c *gin.Context) {
	c.JSON(200, gin.H{
		"x": 1,
		"d": 2,
	})
}

func handleAssetDetail(c *gin.Context) {
	c.JSON(200, gin.H{
		"x": 1,
		"d": 2,
	})
}

func handleTransaction(c *gin.Context) {
	c.JSON(200, gin.H{
		"x": 1,
		"d": 2,
	})
}
