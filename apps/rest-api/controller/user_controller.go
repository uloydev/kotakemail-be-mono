package controller

import (
	"restapi/model"
	"restapi/service"

	"github.com/gofiber/fiber/v2"
	appcontext "kotakemail.id/pkg/context"
)

type UserController struct {
	userService service.UserService
	appCtx      *appcontext.AppContext
}

func NewUserController(userService service.UserService, appCtx *appcontext.AppContext) *UserController {
	return &UserController{
		userService: userService,
		appCtx:      appCtx,
	}
}

// Index godoc
// @Summary List all users
// @Description get users
// @Tags users
// @Produce json
// @Success 200 {array} model.UserResponse
// @Router /users [get]
func (c *UserController) Index(ctx *fiber.Ctx) error {
	users, err := c.userService.ListUser(c.appCtx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(users)
}

// Show godoc
// @Summary Show a user
// @Description get user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} model.UserResponse
// @Router /users/{id} [get]
func (c *UserController) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.userService.GetUser(c.appCtx, id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(user)
}

// Create godoc
// @Summary Create a user
// @Description create a user
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserRequest true "User Request"
// @Success 201 {object} model.UserResponse
// @Router /users [post]
func (c *UserController) Create(ctx *fiber.Ctx) error {
	var userRequest model.UserRequest
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user, err := c.userService.CreateUser(c.appCtx, userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(user)
}

// Update godoc
// @Summary Update a user
// @Description update a user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body model.UserRequest true "User Request"
// @Success 200 {object} model.UserResponse
// @Router /users/{id} [put]
func (c *UserController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var userRequest model.UserRequest
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	user, err := c.userService.UpdateUser(c.appCtx, id, userRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(user)
}

// Delete godoc
// @Summary Delete a user
// @Description delete a user
// @Tags users
// @Param id path string true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func (c *UserController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.userService.DeleteUser(c.appCtx, id); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
