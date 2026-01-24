package errors

import "errors"

var (
	UserNotFound            = errors.New("user not found")
	InvalidCredentials      = errors.New("invalid credentials")
	EmptyUserIdError = errors.New("user id is empty")
	EmptyPlannedDateError = errors.New("planned date is empty")
	EmptyTextMessageError = errors.New("empty message text")
	EmtyOrNegativeTypeId = errors.New("empty or negative type_id")
	InvalidEmailError = errors.New("invalid email")
	InvalidPasswordError = errors.New("invalid password")
	InvalidTokenError = errors.New("invalid token")
	InvalidSingingMethod = errors.New("unexpected signing method")
)