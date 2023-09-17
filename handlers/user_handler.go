package handlers

import (
	"github.com/SDmrly/go_fiber_crud/apps"
	"github.com/SDmrly/go_fiber_crud/models"
	"github.com/SDmrly/go_fiber_crud/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Handler interface {
	Create(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type handler struct {
	service apps.Service
}

func UserHandler(service apps.Service) Handler {
	return &handler{service: service}
}

func (h *handler) Create(ctx *fiber.Ctx) error {
	newUser := models.CreateUser{}
	err := ctx.BodyParser(&newUser)
	utils.ErrorPanics(err)

	h.service.Create(newUser)

	webStatus := models.WebStatus{
		Code:    200,
		Status:  "OK",
		Message: "Successfully created user!",
		Data:    newUser,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webStatus)
}

func (h *handler) FindAll(ctx *fiber.Ctx) error {
	webStatus := models.WebStatus{
		Code:    200,
		Status:  "OK",
		Message: "Successfully listed users!",
		Data:    h.service.FindAll(),
	}

	return ctx.Status(fiber.StatusCreated).JSON(webStatus)
}

func (h *handler) Update(ctx *fiber.Ctx) error {
	updateUser := models.UpdateUser{}
	err := ctx.BodyParser(&updateUser)
	utils.ErrorPanics(err)

	userId := ctx.Params("userId")
	id, _ := strconv.Atoi(userId)

	updateUser.Id = id

	h.service.Update(updateUser)

	webStatus := models.WebStatus{
		Code:    200,
		Status:  "OK",
		Message: "Successfully updated user!",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webStatus)
}

func (h *handler) ChangePassword(ctx *fiber.Ctx) error {
	userPassword := models.UpdatePassword{}
	err := ctx.BodyParser(&userPassword)
	utils.ErrorPanics(err)

	userId := ctx.Params("userId")
	id, _ := strconv.Atoi(userId)
	userPassword.Id = id

	h.service.ChangePassword(userPassword)

	webStatus := models.WebStatus{
		Code:    200,
		Status:  "OK",
		Message: "Successfully updated password!",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webStatus)
}

func (h *handler) FindByID(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	utils.ErrorPanics(err)

	userData := h.service.FindByID(id)

	webStatus := models.WebStatus{
		Code:    200,
		Status:  "OK",
		Message: "Successfully find user!",
		Data:    userData,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webStatus)
}

func (h *handler) Delete(ctx *fiber.Ctx) error {
	userId := ctx.Params("userId")
	id, err := strconv.Atoi(userId)
	utils.ErrorPanics(err)

	h.service.Delete(id)

	webStatus := models.WebStatus{
		Code:    200,
		Status:  "OK",
		Message: "Successfully deleted user!",
		Data:    nil,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webStatus)
}
