package holder

import (
	businesserrors "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/shared/errors"
)

type holderDomain struct {
	name     string
	lastName string
}

func (h *holderDomain) Validate() error {
	var businessErrors businesserrors.BusinessErrors
	if len(h.name) < 3 || len(h.name) > 50 {
		businessErrors = businesserrors.AddError(businessErrors, businesserrors.NewBusinessError("Account Holder Name", "O nome do titular esta invalido.", "holderName"))
	}
	if businessErrors.Len() > 0 {
		return businessErrors
	}
	return nil
}

func (h *holderDomain) GetName() string {
	return h.name
}

func (h *holderDomain) GetLastName() string {
	return h.lastName
}
