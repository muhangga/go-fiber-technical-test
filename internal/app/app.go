package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhangga/config"
	"gorm.io/gorm"
)

type server struct {
	httpServer *fiber.App
	config     config.Config
}

type Server interface {
	RunServer()
}

func InitServer(config config.Config) *server {
	return &server{
		httpServer: fiber.New(),
		config:     config,
	}
}

func (s *server) DB() *gorm.DB {
	return s.config.Database()
}

func (s *server) RunServer() {

	app := s.httpServer

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!!!!!!!!!!!!!")
	})

	app.Listen(":3030")
}
