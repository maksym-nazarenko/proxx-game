package game

type outOfBoundsErr struct {
	msg string
}

func NewOutOfBoundsError(msg string) *outOfBoundsErr {
	return &outOfBoundsErr{msg}
}

func (e outOfBoundsErr) Error() string {
	return e.msg
}
