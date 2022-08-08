package resources

import (
	"demo/internal/daos/dao"
	"demo/internal/daos/daogorm"
	"demo/internal/database"

	"gorm.io/gorm"

	"demo/internal/logging"
	"demo/internal/services"
)

type DefaultResources struct {
	logger          logging.Logger
	dbConnector     database.IDatabaseConnector
	serviceProvider services.ServiceProvider
}

type IResources interface {
	GetDB() database.IDatabaseConnector
	GetLogger() logging.Logger
	GetDAO() dao.IDao
	GetTransactionalDAO() dao.IDao
}

func NewDefaultResources() (IResources, error) {
	si := &DefaultResources{}

	si.dbConnector = &database.ServerDBConnector{
		Logger: si.logger,
	}

	si.serviceProvider = &services.DefaultServiceProvider{
		DBConnector: si.dbConnector,
		Logger:      si.logger,
	}

	return si, nil
}

func (si DefaultResources) GetDB() database.IDatabaseConnector {
	return si.dbConnector
}

func (si DefaultResources) GetLogger() logging.Logger {
	return si.logger
}

func (si DefaultResources) GetDAO() dao.IDao {
	return &daogorm.GormDAOImpl{
		DBConnector: si.GetDB(),
	}
}

func (si DefaultResources) GetTransactionalDAO() dao.IDao {
	t := &daogorm.GormDAOImpl{
		DBConnector: si.GetDB(),
		Logger:      si.logger,
	}
	t.Transactional()
	return t
}

func (si DefaultResources) GetDAOWithTransaction(tx *gorm.DB) dao.IDao {
	t := &daogorm.GormDAOImpl{
		DBConnector: si.GetDB(),
		Logger:      si.logger,
	}
	t.WithTransaction(tx)
	return t
}
