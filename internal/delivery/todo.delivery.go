package delivery

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/muhangga/internal/entity/dto"
	"github.com/muhangga/internal/helper"
	"github.com/muhangga/internal/service"
)

type todoDelivery struct {
	todoService service.TodoService
}

func NewTodoDelivery(service service.TodoService) *todoDelivery {
	return &todoDelivery{todoService: service}
}

func (d *todoDelivery) FindAll(c *fiber.Ctx) error {
	id := c.Query("activity_group_id")

	convID, _ := strconv.Atoi(id)

	if id == "" {
		data, err := d.todoService.FindAll()
		if err != nil {
			res := helper.BuildErrorResponse("Failed to get all todo", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(res)
		}
		res := helper.BuildResponse("Success", data)
		return c.Status(fiber.StatusOK).JSON(res)
	}

	data, err := d.todoService.FindAllByActivityGroupID(int64(convID))
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get all todo", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	res := helper.BuildResponse("Success", data)
	return c.Status(fiber.StatusOK).JSON(res)
}

func (d *todoDelivery) GetTodoById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	todo, err := d.todoService.FindById(int64(id))
	if err != nil {
		res := helper.ValidResponse(statusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id))
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := helper.BuildResponse("Success", todo)
	return c.Status(fiber.StatusOK).JSON(res)
}

func (d *todoDelivery) CreateTodo(c *fiber.Ctx) error {
	var todos dto.TodoDTO

	if err := c.BodyParser(&todos); err != nil {
		res := helper.BuildErrorResponse("Failed to create todo", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	if todos.Title == "" {
		res := helper.ValidResponse(badRequest, "title cannot be null")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	todo, err := d.todoService.Create(todos)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to create todo", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := helper.BuildResponse("Success", todo)
	return c.Status(fiber.StatusOK).JSON(res)
}

func (d *todoDelivery) UpdateTodo(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var todos dto.TodoDTO
	if err := c.BodyParser(&todos); err != nil {
		res := helper.BuildErrorResponse("Failed to update todo", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	_, err := d.todoService.FindById(int64(id))
	if err != nil {
		res := helper.ValidResponse(statusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id))
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	if todos.Title == "" {
		res := helper.ValidResponse(badRequest, "title cannot be null")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}
	if todos.Priority == "" {
		res := helper.ValidResponse(badRequest, "priority cannot be null")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	todo, err := d.todoService.Update(id, todos)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to update todo", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := helper.BuildResponse("Success", todo)
	return c.Status(fiber.StatusOK).JSON(res)
}

func (d *todoDelivery) DeleteTodo(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	err := d.todoService.Delete(int64(id))
	if err != nil {
		res := helper.ValidResponse(statusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id))
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := helper.BuildResponse("Success", helper.EmptyObject{})
	return c.Status(fiber.StatusOK).JSON(res)
}
