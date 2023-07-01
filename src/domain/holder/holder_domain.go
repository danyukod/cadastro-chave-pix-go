package holder

type holderDomain struct {
	name     string
	lastName string
}

func (h *holderDomain) Validate() error {
	if len(h.name) < 3 || len(h.name) > 50 {
		return &ErrInvalidHolderName{}
	}
	return nil
}

func (h *holderDomain) GetName() string {
	return h.name
}

func (h *holderDomain) GetLastName() string {
	return h.lastName
}
