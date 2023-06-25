package enum

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/entity/validation"
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
	return [...]string{"Undefined", "CPF", "CNPJ", "Phone", "Email", "Random"}[t-1]
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
		return validation.ValidateCPF(pixKey)
	case CNPJ:
		return validation.ValidateCNPJ(pixKey)
	case PHONE:
		return validation.ValidatePhone(pixKey)
	case EMAIL:
		return validation.ValidateEmail(pixKey)
	case RANDOM:
		return validation.ValidateRandom(pixKey)
	default:
		return false
	}
}
