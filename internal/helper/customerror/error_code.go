package customerror

import toolboxError "github.com/rahmat412/go-toolbox/error"

type ErrorCode string

var (
	ErrorInternalServer = toolboxError.NewErrorWithCode(toolboxError.ErrUnexpected, "Something Went Wrong")

	ErrorBirthDateFormat = toolboxError.NewErrorWithCode(toolboxError.ErrClient, "Birth date format is not valid")

	ErrorUserNotFound = toolboxError.NewErrorWithCode(toolboxError.ErrNotFound, "User not found")
)
