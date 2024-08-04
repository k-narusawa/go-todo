package middleware

import (
	"errors"
	"go-todo/domain"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GlobalErrorHandler(err error, c echo.Context) {
	for errors.Unwrap(err) != nil {
		err = errors.Unwrap(err)
	}

	switch err {
	case domain.ErrUserNotFound:
		log.Warn(err)
		domainError := domain.ToDomainError(err)
		c.JSON(domain.StatusMap[domainError.Code], domainError)
	case domain.ErrToDoNotFound:
		log.Warn(err)
		domainError := domain.ToDomainError(err)
		c.JSON(domain.StatusMap[domainError.Code], domainError)
	default:
		log.Error(err)
		c.JSON(500, domain.ToDomainError(err))
	}

	c.Echo().DefaultHTTPErrorHandler(err, c)
}
