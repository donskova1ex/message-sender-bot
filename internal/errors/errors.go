package errors

import "errors"

var (
	UserNotFound          = errors.New("user not found")
	InvalidCredentials    = errors.New("invalid credentials")
	EmptyUserIdError      = errors.New("user id is empty")
	EmptyPlannedDateError = errors.New("planned date is empty")
	EmptyTextMessageError = errors.New("empty message text")
	EmptyOrNegativeTypeId = errors.New("empty or negative type_id")
	InvalidEmailError     = errors.New("invalid email")
	InvalidPasswordError  = errors.New("invalid password")
	InvalidTokenError     = errors.New("invalid token")
	InvalidSingingMethod  = errors.New("unexpected signing method")
	MessageNotFoundError  = errors.New("message not found")
	FailedToCreateMessage = errors.New("failed to create message")
)
