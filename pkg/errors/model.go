package errors

type BusinessProcessError struct {
	message string
	status  int
}

func NewBusinessProcessError(message string, status int) *BusinessProcessError {
	return &BusinessProcessError{
		message: message,
		status:  status,
	}
}

func (bpe *BusinessProcessError) Error() string {
	return bpe.message
}
