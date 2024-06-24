package main

import (
	"gin-fleamarket/controllers"
	"gin-fleamarket/models"
	"gin-fleamarket/repositries"
	"gin-fleamarket/servicies"

	"github.com/gin-gonic/gin"
)

func main() {
	items := []models.Item{
		{ID: 1, Name: "商品１", Price: 1000, Description: "説明１", SoldOut: false},
		{ID: 2, Name: "商品2", Price: 2000, Description: "説明2", SoldOut: true},
		{ID: 3, Name: "商品3", Price: 3000, Description: "説明3", SoldOut: false},
	}

	itemRepository := repositries.NewItemMemoryRepositry(items)
	itemService := servicies.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.POST("/items", itemController.Create)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)
	r.Run("localhost:8080") // 0.0.0.0:8080 でサーバーを立てます。
}
