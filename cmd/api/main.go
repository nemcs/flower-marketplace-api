package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/nemcs/flower-marketplace-api/internal/repository"
	"github.com/nemcs/flower-marketplace-api/internal/service"
	clientHttp "github.com/nemcs/flower-marketplace-api/internal/transport/http"
	courierHttp "github.com/nemcs/flower-marketplace-api/internal/transport/http"
	shopHttp "github.com/nemcs/flower-marketplace-api/internal/transport/http"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("Starting API server on :8080")
	r.Run(":8080")

	db, err := sql.Open("postgres", "postgres://flower:secret@localhost:5432/flowerdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	shopRepo := repository.NewShopRepoPostgres(db)
	shopSvc := service.NewShopService(shopRepo)
	shopHttp.NewShopHandler(r, shopSvc)

	clientRepo := repository.NewClientRepoPostgres(db)
	clientSvc := service.NewClientService(clientRepo)
	clientHttp.NewClientHandler(r, clientSvc)

	courierRepo := repository.NewCourierRepoPostgres(db)
	courierSvc := service.NewCourierService(courierRepo)
	courierHttp.NewCourierHandler(r, courierSvc)
}
