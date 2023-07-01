package holder

type ErrInvalidHolderName struct{}

func (e *ErrInvalidHolderName) Error() string {
	return "the holder name is invalid"
}
