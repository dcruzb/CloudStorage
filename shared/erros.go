package shared

type RemoteError struct {
	ErrorMessage string
}

func NewRemoteError(errorMessage string) *RemoteError {
	return &RemoteError{errorMessage}
}

func (re RemoteError) Error() string {
	return re.ErrorMessage
}
