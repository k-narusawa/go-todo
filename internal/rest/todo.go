package rest

import (
	"go-todo/domain/value"
	"go-todo/usecase/todo"

	"go-todo/internal/middleware"

	"github.com/labstack/echo/v4"
)

type ToDoHandler struct {
	todoService todo.Service
}

func NewToDoHandler(e *echo.Echo, todoService todo.Service) {
	handler := &ToDoHandler{
		todoService: todoService,
	}

	users := e.Group("/users/:user_id")

	users.POST("/todos", handler.Create)
	users.GET("/todos", handler.FindByUserId)
	users.PUT("/todos/:todo_id", handler.ChangeToDoDone)
}

func (h *ToDoHandler) Create(c echo.Context) error {
	in := todo.CreateToDoInput{}
	if err := c.Bind(&in); err != nil {
		return middleware.HandleError(c, err)
	}
	in.UserID = value.OfUserID(c.Param("user_id"))

	todo, err := h.todoService.Create(in)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return c.JSON(200, todo)
}

func (h *ToDoHandler) FindByUserId(c echo.Context) error {
	userId := value.OfUserID(c.Param("user_id"))

	todos, err := h.todoService.FindByUserId(userId)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return c.JSON(200, todos)
}

func (h *ToDoHandler) ChangeToDoDone(c echo.Context) error {
	userId := value.OfUserID(c.Param("user_id"))
	todoId := value.OfToDoID(c.Param("todo_id"))

	in := todo.ChangeStatusInput{}
	if err := c.Bind(&in); err != nil {
		return middleware.HandleError(c, err)
	}
	in.UserID = userId
	in.ToDoID = todoId

	todo, err := h.todoService.ChangeToDoDone(in)
	if err != nil {
		return middleware.HandleError(c, err)
	}

	return c.JSON(200, todo)
}
