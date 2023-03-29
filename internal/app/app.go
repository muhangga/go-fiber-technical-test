package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhangga/config"
	"github.com/muhangga/internal/delivery"
	"github.com/muhangga/internal/helper"
	"github.com/muhangga/internal/repository"
	"github.com/muhangga/internal/service"
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

	app.Use(helper.HandleCors())

	activityRepository := repository.NewActivitiesRepository(s.DB())
	activityService := service.NewActivitiesService(activityRepository)
	activityDelivery := delivery.NewActivitiesDelivery(activityService)

	todoRepository := repository.NewTodoRepository(s.DB())
	todoService := service.NewTodoService(todoRepository)
	todoDelivery := delivery.NewTodoDelivery(todoService)

	activity_groups := app.Group("/activity-groups")
	activity_groups.Post("/", activityDelivery.CreateActivities)
	activity_groups.Get("/", activityDelivery.GetAllActivities)
	activity_groups.Get("/:id", activityDelivery.GetActivitiesByID)
	activity_groups.Patch("/:id", activityDelivery.UpdateActivities)
	activity_groups.Delete("/:id", activityDelivery.DeleteActivities)

	todo := app.Group("/todo-items")
	todo.Post("/", todoDelivery.CreateTodo)
	todo.Get("/", todoDelivery.FindAll)
	todo.Get("/:id", todoDelivery.GetTodoById)
	todo.Patch("/:id", todoDelivery.UpdateTodo)
	todo.Delete("/:id", todoDelivery.DeleteTodo)

	app.Listen(":3030")
}
