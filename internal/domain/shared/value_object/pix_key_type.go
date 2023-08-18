package value_object

import (
	"github.com/klassmann/cpfcnpj"
	"regexp"
	"strings"
)

const (
	CPF                   = "CPF"
	CNPJ                  = "CNPJ"
	Phone                 = "PHONE"
	Email                 = "EMAIL"
	Random                = "RANDOM"
	Undefined             = "UNDEFINED"
	MaxEmailLength        = 77
	MinEmailLength        = 3
	RandomRegexValidation = "^[a-zA-Z0-9]{36}$"
	PhoneRegexValidation  = "\\+((\\d{11,14}))"
	EmailRegexValidation  = "^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$"
)

type PixKeyType interface {
	Validate(pixKey string) bool
	GetType() string
}

func PixKeyTypeFromText(text string) PixKeyType {
	switch strings.ToUpper(text) {
	case CPF:
		return CPFType{
			pixKeyType: CPF,
		}
	case CNPJ:
		return CNPJType{
			pixKeyType: CNPJ,
		}
	case Phone:
		return PhoneType{
			pixKeyType: Phone,
		}
	case Email:
		return EmailType{
			pixKeyType: Email,
		}
	case Random:
		return RandomType{
			pixKeyType: Random,
		}
	default:
		return UndefinedType{
			pixKeyType: text,
		}
	}
}

type CPFType struct {
	pixKeyType string
}

func (c CPFType) Validate(pixKey string) bool {
	return cpfcnpj.ValidateCPF(pixKey)
}

func (c CPFType) GetType() string {
	return c.pixKeyType
}

type CNPJType struct {
	pixKeyType string
}

func (c CNPJType) Validate(pixKey string) bool {
	return cpfcnpj.ValidateCNPJ(pixKey)
}

func (c CNPJType) GetType() string {
	return c.pixKeyType
}

type PhoneType struct {
	pixKeyType string
}

func (c PhoneType) Validate(pixKey string) bool {
	pattern := regexp.MustCompile(PhoneRegexValidation)

	return pattern.MatchString(pixKey)
}

func (c PhoneType) GetType() string {
	return c.pixKeyType
}

type EmailType struct {
	pixKeyType string
}

func (c EmailType) Validate(pixKey string) bool {
	pattern := regexp.MustCompile(EmailRegexValidation)

	return pattern.MatchString(pixKey) && len(pixKey) <= MaxEmailLength && len(pixKey) >= MinEmailLength
}

func (c EmailType) GetType() string {
	return c.pixKeyType
}

type RandomType struct {
	pixKeyType string
}

func (c RandomType) Validate(pixKey string) bool {
	pattern := regexp.MustCompile(RandomRegexValidation)

	return pattern.MatchString(pixKey)
}

func (c RandomType) GetType() string {
	return c.pixKeyType
}

type UndefinedType struct {
	pixKeyType string
}

func (d UndefinedType) Validate(_ string) bool {
	return false
}

func (d UndefinedType) GetType() string {
	return d.pixKeyType
}
