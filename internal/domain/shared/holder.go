package shared

type HolderDomainInterface interface {
	Validate() error
	GetName() string
	GetLastName() string
}

func NewHolderDomain(name string, lastName string) (HolderDomainInterface, error) {
	holder := holderDomain{
		name:     name,
		lastName: lastName,
	}
	if err := holder.Validate(); err != nil {
		return nil, err
	}
	return &holder, nil
}

type holderDomain struct {
	name     string
	lastName string
}

func (h *holderDomain) Validate() error {
	var businessErrors BusinessErrors
	if len(h.name) < 3 || len(h.name) > 50 {
		businessErrors = AddError(businessErrors, *NewBusinessError("Account Holder Name", "O nome do titular esta invalido.", "holderName"))
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
