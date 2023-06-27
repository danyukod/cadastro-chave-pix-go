package domain

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
