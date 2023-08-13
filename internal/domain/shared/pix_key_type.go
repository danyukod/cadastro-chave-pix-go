package shared

import (
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
