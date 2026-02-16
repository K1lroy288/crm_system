package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"visit-service/config"
	"visit-service/handler"
	"visit-service/repository"
	"visit-service/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const migrationsDir = "migrations"

//go:embed migrations/*.sql
var MigrationsFS embed.FS

func main() {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	migrator := repository.MustGetNewMigrator(MigrationsFS, migrationsDir)

	err := migrator.ApplyMigrationsWithGORM(dsn)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	repo := repository.NewVisitRepository(db)
	service := service.NewVisitService(repo)
	handler := handler.NewVisitHandler(service)

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "User service is up!")
	})

	api := r.Group("/visit")
	{
		api.GET("/visits", handler.GetVisits)

		api.POST("/visits", handler.CreateVisit)

		api.DELETE("/visits/:id", handler.DeleteVisit)
	}

	addr := fmt.Sprintf(":%s", cfg.Port)
	r.Run(addr)
}
