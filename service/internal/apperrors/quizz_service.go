package apperrors

import "errors"

var (
	ErrorSingleQuizzHaveMoreResult = errors.New("QuizzService.ErrorSingleQuizzHaveMoreResult")
	ErrorQuizzIdInvalid            = errors.New("QuizzService.ErrorQuizzIdInvalid")
	ErrorQuizzEntityIdInvalid      = errors.New("QuizzService.ErrorQuizzEntityIdInvalid")
	ErrorQuizzEntityTypeInvalid    = errors.New("QuizzService.ErrorQuizzEntityTypeInvalid")
)
