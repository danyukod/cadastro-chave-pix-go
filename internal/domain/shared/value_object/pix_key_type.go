package value_object

import (
	"github.com/klassmann/cpfcnpj"
	"regexp"
	"strings"
)

type PixKeyType int

const (
	UNDEFINED_PIX_KEY PixKeyType = iota
	CPF
	CNPJ
	PHONE
	EMAIL
	RANDOM
	MAX_EMAIL_LENGTH = 77
	MIN_EMAIL_LENGTH = 3
)

func (t PixKeyType) String() string {
	return [...]string{"Undefined", "CPF", "CNPJ", "Phone", "Email", "Random"}[t]
}

func (t PixKeyType) EnumIndex() int {
	return int(t)
}

func (t *PixKeyType) UnmarshalText(text []byte) error {
	*t = PixKeyTypeFromText(string(text))
	return nil
}

func PixKeyTypeFromText(text string) PixKeyType {
	switch strings.ToLower(text) {
	case "cpf":
		return CPF
	case "cnpj":
		return CNPJ
	case "phone":
		return PHONE
	case "email":
		return EMAIL
	case "random":
		return RANDOM
	default:
		return UNDEFINED_PIX_KEY
	}
}

func (t PixKeyType) PixKeyValidate(pixKey string) bool {
	switch t {
	case CPF:
		return ValidateCPF(pixKey)
	case CNPJ:
		return ValidateCNPJ(pixKey)
	case PHONE:
		return ValidatePhone(pixKey)
	case EMAIL:
		return ValidateEmail(pixKey)
	case RANDOM:
		return ValidateRandom(pixKey)
	default:
		return false
	}
}

func ValidateCPF(cpf string) bool {
	return cpfcnpj.ValidateCPF(cpf)
}

func ValidateCNPJ(cnpj string) bool {
	return cpfcnpj.ValidateCNPJ(cnpj)
}

func ValidatePhone(phone string) bool {
	pattern := regexp.MustCompile("\\+((\\d{11,14}))")

	return pattern.MatchString(phone)
}

func ValidateEmail(email string) bool {
	pattern := regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$")

	return pattern.MatchString(email) && len(email) <= MAX_EMAIL_LENGTH && len(email) >= MIN_EMAIL_LENGTH
}

func ValidateRandom(random string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9]{36}$")

	return pattern.MatchString(random)
}
