package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"user-service/config"
	"user-service/handler"
	"user-service/repository"
	"user-service/service"

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

	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "User service is up!")
	})

	api := r.Group("/auth")
	{

		api.POST("/login", handler.GetUserByUsername)

		api.POST("/register", handler.CreateUser)
	}

	api2 := r.Group("/user")
	{
		api2.GET(":lastname", handler.GetUserByLastname)
	}

	addr := fmt.Sprintf(":%s", cfg.Port)
	r.Run(addr)
}
