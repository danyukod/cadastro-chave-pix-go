package pix_key

type ErrInvalidPixKeyType struct{}

type ErrInvalidPixKey struct{}

type ErrPixKeyAlreadyExists struct{}

func (e *ErrInvalidPixKeyType) Error() string {
	return "Invalid pix key type"
}

func (e *ErrInvalidPixKey) Error() string {
	return "Invalid pix key"
}

func (e *ErrPixKeyAlreadyExists) Error() string {
	return "Pix key already exists"
}
