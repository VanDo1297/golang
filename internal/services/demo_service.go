package services

import (
	"demo/internal/models/dbmodels"
)

type DefaultDemoService struct {
	DefaultBaseService
	demo *dbmodels.Demo
}

type DemoService interface {
	BaseService
	GetDemo() ([]*dbmodels.Demo, error)
}

func (service DefaultDemoService) GetDemo() ([]*dbmodels.Demo, error) {
	demos, err := service.dao.GetDemoDAO().GetDemo()
	if err != nil {
		return nil, err
	}

	return demos, nil
}
