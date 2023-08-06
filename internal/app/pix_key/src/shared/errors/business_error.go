package errors

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type BusinessErrors []validator.FieldError

type BusinessError struct {
	BusinessField string
	BusinessParam string
	Message       string
}

func NewBusinessError(field, message, param string) *BusinessError {
	return &BusinessError{BusinessField: field, Message: message, BusinessParam: param}
}

func (b *BusinessError) ActualTag() string {
	//TODO implement me
	panic("implement me")
}

func (b *BusinessError) Namespace() string {
	//TODO implement me
	panic("implement me")
}

func (b *BusinessError) StructNamespace() string {
	//TODO implement me
	panic("implement me")
}

func (b *BusinessError) StructField() string {
	//TODO implement me
	panic("implement me")
}

func (b *BusinessError) Value() interface{} {
	//TODO implement me
	panic("implement me")
}

func (b *BusinessError) Kind() reflect.Kind {
	//TODO implement me
	panic("implement me")
}

func (b *BusinessError) Type() reflect.Type {
	//TODO implement me
	panic("implement me")
}

func (b *BusinessError) Translate(ut ut.Translator) string {
	//TODO implement me
	panic("implement me")
}

func (b *BusinessError) Tag() string {
	return "business"
}

func (b *BusinessError) Field() string {
	return b.BusinessField
}

func (b *BusinessError) Param() string {
	return b.BusinessParam
}

func (b *BusinessError) Error() string {
	return b.Message
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

func AddError(es BusinessErrors, e *BusinessError) BusinessErrors {
	return append(es, e)
}

func AppendErrors(es BusinessErrors, newEs BusinessErrors) BusinessErrors {
	return append(es, newEs...)
}
