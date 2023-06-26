package domain

import "errors"

type Holder struct {
	Name     string
	LastName string
}

func NewHolder(name string, lastName string) (*Holder, error) {
	holder := Holder{
		Name:     name,
		LastName: lastName,
	}
	if err := holder.Validate(); err != nil {
		return nil, err
	}
	return &holder, nil
}

func (h *Holder) Validate() error {
	if len(h.Name) < 3 || len(h.Name) > 50 {
		return errors.New("the holder name is invalid")
	}
	return nil
}
