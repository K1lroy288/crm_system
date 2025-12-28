package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"visit-service/config"
	"visit-service/repository"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "User service is up!")
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	r.Run(addr)
}
