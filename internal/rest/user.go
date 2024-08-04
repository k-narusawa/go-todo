package rest

import (
	"go-todo/domain/value"
	"go-todo/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(e *echo.Echo, userService user.Service) {
	handler := &UserHandler{
		userService: userService,
	}

	e.POST("/users", handler.Register)
	e.GET("/users/:user_id", handler.Get)
	e.GET("/users", handler.GetAll)
	e.PUT("/users/:user_id", handler.Change)
}

func (h *UserHandler) Register(c echo.Context) error {
	in := new(user.RegisterUserInput)
	if err := c.Bind(&in); err != nil {
		return err
	}

	user, err := h.userService.Register(*in)
	if err != nil {
		return err
	}

	return c.JSON(201, user)
}

func (h *UserHandler) Get(c echo.Context) error {
	userID := c.Param("user_id")
	in := user.GetUserInput{
		UserID: value.OfUserID(userID),
	}
	user, err := h.userService.Get(in)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func (h *UserHandler) GetAll(c echo.Context) error {
	users, err := h.userService.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(200, users)
}

func (h *UserHandler) Change(c echo.Context) error {
	in := user.ChangeUserInput{}
	userId := c.Param("user_id")
	if err := c.Bind(&in); err != nil {
		return err
	}
	in.UserID = value.OfUserID(userId)

	user, err := h.userService.Change(in)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}
