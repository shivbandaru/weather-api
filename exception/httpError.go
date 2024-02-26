package exception

type RestError struct {
	ErrStatus int    `json:"status,omitempty"`
	ErrError  string `json:"error,omitempty"`
}

// New Error
func NewError(status int, err string) RestError {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
	}
}
