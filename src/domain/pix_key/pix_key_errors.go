package pix_key

type ErrInvalidPixKeyType struct{}

func (e *ErrInvalidPixKeyType) Error() string {
	return "Invalid pix key type"
}
