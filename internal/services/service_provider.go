package services

import (
	"demo/internal/daos/daogorm"
	"demo/internal/database"
	"demo/internal/logging"
	"demo/internal/models/dbmodels"
)

type DefaultServiceProvider struct {
	DBConnector database.IDatabaseConnector
	Logger      logging.Logger
}

type ServiceProvider interface {
	NewDemoService(demo *dbmodels.Demo) DemoService
}

func (provider DefaultServiceProvider) NewDemoService(demo *dbmodels.Demo) DemoService {
	return &DefaultDemoService{
		DefaultBaseService: DefaultBaseService{dao: &daogorm.GormDAOImpl{
			DBConnector: provider.DBConnector,
		}},
		demo: demo,
	}
}
