package ui

// GracefulExitError is used to signal a graceful shutdown
type GracefulExitError struct {
	msg string
}

func (e GracefulExitError) Error() string {
	return e.msg
}
