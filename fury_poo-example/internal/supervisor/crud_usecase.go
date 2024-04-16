package supervisor

import (
	"fmt"
	"users_example/internal/supervisor/developer"
)

type CRUDUseCase interface {
	Create(name string, team string, seniority string) (*developer.Developer, error)
	Read(id string) (*developer.Developer, error)
	Update(name string, team string, seniority string) error
	Delete(id string) error
}

type useCaseCRUD struct {
	repo developer.Repository
}

func NewCRUDUseCase(repo developer.Repository) *useCaseCRUD {
	return &useCaseCRUD{repo: repo}
}

func (c useCaseCRUD) Create(name string, team string, seniority string) (*developer.Developer, error) {
	s, err := developer.SeniorityFromString(seniority)
	if err != nil {
		return nil, err
	}

	dev := developer.Developer{
		Name:      name,
		Team:      team,
		Seniority: s,
	}
	if err := c.repo.Save(&dev); err != nil {
		return nil, err
	}

	return &dev, nil
}

func (c useCaseCRUD) Read(id string) (*developer.Developer, error) {
	dev, err := c.repo.Get(id)
	if err != nil {
		return nil, err
	}

	if dev == nil {
		return nil, fmt.Errorf("developer:%s not found", id)
	}

	return dev, nil
}

func (c useCaseCRUD) Update(name string, team string, seniority string) error {
	panic("implement me")
}

func (c useCaseCRUD) Delete(id string) error {
	panic("implement me")
}
