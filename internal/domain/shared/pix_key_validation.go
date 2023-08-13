package shared

import (
	"github.com/klassmann/cpfcnpj"
	"regexp"
)

type PixKeyValidation interface {
	PixKeyValidate(pixKey string) bool
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

const MAX_EMAIL_LENGTH = 77
const MIN_EMAIL_LENGTH = 3

func ValidateEmail(email string) bool {
	pattern := regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z0-9.\\-]+\\.[a-z]{2,4}$")

	return pattern.MatchString(email) && len(email) <= MAX_EMAIL_LENGTH && len(email) >= MIN_EMAIL_LENGTH
}

func ValidateRandom(random string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9]{36}$")

	return pattern.MatchString(random)
}
