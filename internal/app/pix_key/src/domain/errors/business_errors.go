package errors

type BusinessError struct {
	Field   string
	Message string
}

type BusinessErrors []BusinessError

func NewBusinessError(field, message string) *BusinessError {
	return &BusinessError{Field: field, Message: message}
}

func (es BusinessErrors) Error() string {
	var result string
	for _, e := range es {
		result += e.Error() + "\n"
	}
	return result
}

func (es BusinessErrors) Len() int {
	return len(es)
}

func AddError(es BusinessErrors, e BusinessError) BusinessErrors {
	return append(es, e)
}

func AppendErrors(es BusinessErrors, newEs BusinessErrors) BusinessErrors {
	return append(es, newEs...)
}

func (be BusinessErrors) HasErrors() bool {
	return be.Len() > 0
}

func (e *BusinessError) Error() string {
	return e.Message
}
