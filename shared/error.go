package shared

import "fmt"

// NOTE: errorをフィールドに持つ場合はUnwrap()を実装する

type (
	CustomError interface {
		Error() string
		Is(error) bool
	}

	ValidationError struct {
		Field string
		Err   string
	}
	NotFoundError struct {
		Resource string
	}
	UnauthorizedError struct{}
	ConflictError     struct {
		Resource string
	}
	ArgumentError struct {
		Field string
		Err   string
	}
)

func NewValidationError(field, err string) CustomError {
	return &ValidationError{Field: field, Err: err}
}
func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field %s: %s", e.Field, e.Err)
}
func (e *ValidationError) Is(target error) bool {
	_, ok := target.(*ValidationError)
	return ok
}

func NewNotFoundError(resource string) CustomError {
	return &NotFoundError{Resource: resource}
}
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Resource)
}
func (e *NotFoundError) Is(target error) bool {
	_, ok := target.(*NotFoundError)
	return ok
}

func NewUnauthorizedError() CustomError {
	return &UnauthorizedError{}
}
func (e *UnauthorizedError) Error() string {
	return "Unauthorized"
}
func (e *UnauthorizedError) Is(target error) bool {
	_, ok := target.(*UnauthorizedError)
	return ok
}

func NewConflictError(resource string) CustomError {
	return &ConflictError{Resource: resource}
}
func (e *ConflictError) Error() string {
	return fmt.Sprintf("%s already exists", e.Resource)
}
func (e *ConflictError) Is(target error) bool {
	_, ok := target.(*ConflictError)
	return ok
}

func NewArgumentError(field, err string) CustomError {
	return &ArgumentError{Field: field, Err: err}
}
func (e *ArgumentError) Error() string {
	return fmt.Sprintf("Error on field %s: %s", e.Field, e.Err)
}
func (e *ArgumentError) Is(target error) bool {
	_, ok := target.(*ArgumentError)
	return ok
}
