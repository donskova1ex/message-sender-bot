package errors

import "errors"

var EmptyUserIdError = errors.New("user id is empty")
var EmptyPlannedDateError = errors.New("planned date is empty")
var EmptyTextMessageError = errors.New("empty message text")
var EmtyOrNegativeTypeId = errors.New("empty or negative type_id")
