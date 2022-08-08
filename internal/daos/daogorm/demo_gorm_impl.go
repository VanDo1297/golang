package daogorm

import (
	"demo/internal/models/dbmodels"

	"github.com/go-errors/errors"
	"gorm.io/gorm"
)

type DemoDAOGormImpl struct {
	DB            *gorm.DB
	Transactional bool
}

func (demoDAO DemoDAOGormImpl) CreateDemo(demo *dbmodels.Demo) error {
	res := demoDAO.DB.Create(demo)
	if res.Error != nil {
		return errors.Wrap(res.Error, 0)
	}
	return nil
}

func (demoDAO DemoDAOGormImpl) getDemo() ([]*dbmodels.Demo, error) {
	var demos []*dbmodels.Demo
	res := demoDAO.DB.Raw("select * from demo").Scan(&demos)
	if res.Error != nil {
		return nil, errors.Wrap(res.Error, 0)
	}
	return demos, nil
}

func (demoDAO DemoDAOGormImpl) GetDemo() ([]*dbmodels.Demo, error) {
	return demoDAO.getDemo()
}
