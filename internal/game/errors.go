package game

type outOfBoundsErr struct {
	msg string
}

// NewOutOfBoundsError instantiates OutOfBounds error
func NewOutOfBoundsError(msg string) *outOfBoundsErr {
	return &outOfBoundsErr{msg}
}

func (e outOfBoundsErr) Error() string {
	return e.msg
}
