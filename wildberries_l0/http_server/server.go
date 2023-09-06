package http_server

import (
	"github.com/gin-gonic/gin"
	"log"
	"solve/producer"
)

var client *producer.PostgresClient
var cache *producer.MemoryCache

// handler getOrder
func getOrder(c *gin.Context) {
	id := c.Param("id")

	// Try to get order from cache

	orderFromCache, err := cache.Get(id)
	if err == nil {
		c.JSON(200, orderFromCache)
		log.Println("Order get from cache")
		return
	} else {
		log.Println("Error getting order from cache:", err)
	}

	// If order is not found in cache, retrieve it from Postgres
	order, err := client.GetOrderFromPostgres(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	// Save order to cache
	cache.Add(*order)

	log.Println("Order get from DataBase")
	c.JSON(200, order)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/order/:id", getOrder)

	return router
}

func StartHTTPServer(postgresClient *producer.PostgresClient, memCache *producer.MemoryCache) {
	cache = memCache
	client = postgresClient
	router := setupRouter()

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
