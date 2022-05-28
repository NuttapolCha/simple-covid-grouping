package custom_error

type UserError struct {
	Message    string
	StatusCode int
}

func (err *UserError) Error() string {
	return err.Message
}
