package spacecraft

type ValidationErr struct {
	Message string
}

func (e *ValidationErr) Error() string {
	return e.Message
}
