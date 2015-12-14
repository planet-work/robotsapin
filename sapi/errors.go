package sapi

type RbacForbiddenError struct {
	msg string
}

func (e *RbacForbiddenError) Error() string { return e.msg }

type RbacUnknownMethodError struct {
	msg string
}

func (e *RbacUnknownMethodError) Error() string { return e.msg }

type ResourceNotFoundError struct {
	msg string
}

func (e *ResourceNotFoundError) Error() string { return e.msg }

type ResourceForbiddenError struct {
	msg string
}

func (e *ResourceForbiddenError) Error() string { return e.msg }

type ResourceDuplicateError struct {
	msg string
}

func (e *ResourceDuplicateError) Error() string { return e.msg }

type ResourceValidationError struct {
	msg string
}

func (e *ResourceValidationError) Error() string { return e.msg }

type AuthFailedError struct {
	msg string
}

func (e *AuthFailedError) Error() string { return e.msg }

type UnexpectedError struct {
	msg string
}

func (e *UnexpectedError) Error() string { return e.msg }

type QueryFailedError struct {
	msg string
}

func (e *QueryFailedError) Error() string { return e.msg }
