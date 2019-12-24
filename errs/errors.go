package errs

type DuplicityCartError struct {
	cause error
	msg string
}

func NewDuplicityCartError(cause error) DuplicityCartError {
	return DuplicityCartError{cause, "Cart already exists"}
}

func (e DuplicityCartError) Error() string {
	return e.msg
}

func (e DuplicityCartError) Cause() string {
	return e.cause.Error()
}