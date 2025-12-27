package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"user-service/config"
	"user-service/repository"

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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	migrator := repository.MustGetNewMigrator(MigrationsFS, migrationsDir)

	err = migrator.ApplyMigrationsWithGORM(db)
	if err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	r := gin.Default()

	api := r.Group("/user")
	{
		api.GET("/health", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "User service is up!")
		})
	}
}
