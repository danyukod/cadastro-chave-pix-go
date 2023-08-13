package value_object

type BusinessErrors []BusinessError

type BusinessError struct {
	field   string
	param   string
	message string
}

func NewBusinessError(field, message, param string) *BusinessError {
	return &BusinessError{field: field, message: message, param: param}
}

func (b *BusinessError) Tag() string {
	return "business"
}

func (b *BusinessError) Field() string {
	return b.field
}

func (b *BusinessError) Param() string {
	return b.field
}

func (b *BusinessError) Error() string {
	return b.message
}

func (bs BusinessErrors) Error() string {
	var result string
	for _, e := range bs {
		result += e.Error() + "\n"
	}
	return result
}

func (bs BusinessErrors) HasErrors() bool {
	return bs.Len() > 0
}

func (bs BusinessErrors) Len() int {
	return len(bs)
}

func AddError(es BusinessErrors, e BusinessError) BusinessErrors {
	return append(es, e)
}

func AppendErrors(es BusinessErrors, newEs BusinessErrors) BusinessErrors {
	return append(es, newEs...)
}

func CreatePixKeyAlreadyExistsError(pixKey string) *BusinessError {
	return NewBusinessError("PixKey", "Chave pix ja cadastrada.", pixKey)
}
