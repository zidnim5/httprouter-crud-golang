package exception

type NotFound struct {
	error string
}

func NewNotFoundError(err string) NotFound {
	return NotFound{error: err}
}
