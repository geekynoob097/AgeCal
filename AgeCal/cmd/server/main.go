package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"

	"AgeCal/db/sqlc"
	"AgeCal/internal/handler"
	"AgeCal/internal/logger"
	"AgeCal/internal/middleware"
	"AgeCal/internal/repository"
	"AgeCal/internal/routes"
	"AgeCal/internal/service"
	"AgeCal/internal/validator"
)

func main() {
	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if logger.Log != nil {
			_ = logger.Log.Sync()
		}
	}()
	db, err := sql.Open("postgres", "postgres://postgres:Family@localhost:5432/userdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	validator.Init()
	queries := sqlc.New(db)
	repo := repository.NewUserRepository(queries)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	app := fiber.New()
	app.Use(middleware.RequestID())
	app.Use(middleware.ZapLogger())
	routes.Register(app, handler)

	log.Fatal(app.Listen(":8080"))
}
