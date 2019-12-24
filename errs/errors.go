package errs

type Error struct {
	cause error
	msg string
}

func (e Error) Error() string {
	return e.msg
}

func (e Error) Cause() string {
	return e.cause.Error()
}

func DuplicityCartError(cause error) Error {
	return Error{cause, "Cart already exists"}
}

func CartDoesNotExist(cause error) Error {
	return Error{cause, "Cart does not exist"}
}

func DuplicityItemError(cause error) Error {
	return Error{cause, "Item already exists"}
}

func ItemDoesNotExist(cause error) Error {
	return Error{cause, "Item does not exist"}
}

func CartOrItemDoesNotExist(cause error) Error {
	return Error{cause, "Cart does not exist or Item does not exist"}
}