package delivery

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/muhangga/internal/entity/dto"
	"github.com/muhangga/internal/helper"
	"github.com/muhangga/internal/service"
)

var (
	badRequest     = "Bad Request"
	statusNotFound = "Not Found"
)

func (d *activitiesDelivery) GetAllActivities(c *fiber.Ctx) error {
	activities, err := d.activitiesService.GetAllActivities()
	if err != nil {
		res := helper.BuildErrorResponse("Failed to get activities", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := helper.BuildResponse("Success", activities)
	return c.Status(fiber.StatusOK).JSON(res)
}

func (d *activitiesDelivery) GetActivitiesByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	activities, err := d.activitiesService.GetActivitiesByID(int64(id))
	if err != nil {
		res := helper.ValidResponse(statusNotFound, fmt.Sprintf("Activity with ID %d Not Found", id))
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := helper.BuildResponse("Success", activities)
	return c.Status(fiber.StatusOK).JSON(res)
}

func (d *activitiesDelivery) CreateActivities(c *fiber.Ctx) error {
	var activitiesDTO dto.ActivitiesDTO
	if err := c.BodyParser(&activitiesDTO); err != nil {
		res := helper.BuildErrorResponse("Failed to create activities", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	if isEmailExist := d.activitiesService.FindByEmail(activitiesDTO.Email); isEmailExist {
		res := helper.ValidResponse(badRequest, "email already exist")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if activitiesDTO.Title == "" {
		res := helper.ValidResponse(badRequest, "title cannot be null")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if activitiesDTO.Email == "" {
		res := helper.ValidResponse(badRequest, "email cannot be null")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	activities, err := d.activitiesService.CreateActivities(activitiesDTO)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to create activities", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := helper.BuildResponse("Success", activities)
	return c.Status(fiber.StatusOK).JSON(res)
}

func (d *activitiesDelivery) UpdateActivities(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	var activitiesDTO dto.ActivitiesDTO
	if err := c.BodyParser(&activitiesDTO); err != nil {
		res := helper.BuildErrorResponse("Failed to update activities", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	_, err := d.activitiesService.GetActivitiesByID(int64(id))
	if err != nil {
		res := helper.ValidResponse(statusNotFound, fmt.Sprintf("Activity with ID %d Not Found", id))
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	if activitiesDTO.Title == "" {
		res := helper.ValidResponse(badRequest, "title cannot be null")
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	activities, err := d.activitiesService.UpdateActivities(activitiesDTO, id)
	if err != nil {
		res := helper.BuildErrorResponse("Failed to update activities", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	res := helper.BuildResponse("Success", activities)
	return c.Status(fiber.StatusOK).JSON(res)
}

func (d *activitiesDelivery) DeleteActivities(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	err := d.activitiesService.DeleteActivities(id)
	if err != nil {
		res := helper.ValidResponse(statusNotFound, fmt.Sprintf("Activity with ID %d Not Found", id))
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := helper.BuildResponse("Success", helper.EmptyObject{})
	return c.Status(fiber.StatusOK).JSON(res)
}

type activitiesDelivery struct {
	activitiesService service.ActivitiesService
}

func NewActivitiesDelivery(service service.ActivitiesService) *activitiesDelivery {
	return &activitiesDelivery{service}
}
