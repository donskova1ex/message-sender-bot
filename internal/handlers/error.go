package handlers

import (
	"errors"

	custom_errors "message-sender-bot/internal/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
)

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *ApiError) Send(c *fiber.Ctx) error {
	return c.Status(e.Status).JSON(e)
}

func HandleError(err error) *ApiError {
	switch {
	case errors.Is(err, custom_errors.InvalidTokenError),
		errors.Is(err, custom_errors.InvalidSingingMethod):
		return &ApiError{
			Status:  fiber.StatusUnauthorized,
			Message: "invalid or expired session",
		}

	case errors.Is(err, custom_errors.InvalidCredentials),
		errors.Is(err, custom_errors.UserNotFound):
		return &ApiError{
			Status:  fiber.StatusUnauthorized,
			Message: "invalid credentials",
		}

	case errors.Is(err, custom_errors.InvalidEmailError):
		return &ApiError{
			Status:  fiber.StatusBadRequest,
			Message: "invalid email format",
		}

	case errors.Is(err, custom_errors.InvalidPasswordError):
		return &ApiError{
			Status:  fiber.StatusBadRequest,
			Message: "invalid password",
		}

	case errors.Is(err, custom_errors.EmptyPlannedDateError):
		return &ApiError{
			Status:  fiber.StatusBadRequest,
			Message: "planned date is empty",
		}

	case errors.Is(err, custom_errors.EmptyTextMessageError):
		return &ApiError{
			Status:  fiber.StatusBadRequest,
			Message: "empty message text",
		}

	case errors.Is(err, custom_errors.EmptyOrNegativeTypeId):
		return &ApiError{
			Status:  fiber.StatusBadRequest,
			Message: "empty or negative type_id",
		}

	case errors.Is(err, custom_errors.EmptyUserIdError):
		return &ApiError{
			Status:  fiber.StatusBadRequest,
			Message: "user id is empty",
		}
	case errors.Is(err, custom_errors.MessageNotFoundError):
		return &ApiError{
			Status:  fiber.StatusNotFound,
			Message: "message not found",
		}
	case errors.Is(err, custom_errors.FailedToCreateMessage):
		return &ApiError{
			Status:  fiber.StatusBadRequest,
			Message: "failed to create message",
		}

	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return &ApiError{
				Status:  fiber.StatusConflict,
				Message: "email already exists",
			}
		case "23503":
			return &ApiError{
				Status:  fiber.StatusBadRequest,
				Message: "specified type does not exist",
			}
		case "42P01":
			return &ApiError{
				Status:  fiber.StatusInternalServerError,
				Message: "something went wrong",
			}

		default:
			return &ApiError{
				Status:  fiber.StatusInternalServerError,
				Message: "database operation failed",
			}
		}
	}
	return &ApiError{
		Status:  fiber.StatusInternalServerError,
		Message: "internal server error",
	}

}
