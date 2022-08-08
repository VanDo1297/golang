package daogorm

import (
	"demo/internal/daos/dao"
	"demo/internal/daos/interfaces"
	"demo/internal/database"
	"demo/internal/logging"

	"github.com/go-errors/errors"
	"gorm.io/gorm"
)

type GormDAOImpl struct {
	DBConnector     database.IDatabaseConnector
	tx              *gorm.DB
	isTransactional bool
	Logger          logging.Logger
}

func (dao *GormDAOImpl) WithTransaction(tx *gorm.DB) {
	dao.isTransactional = true
	dao.tx = tx
}

func (dao *GormDAOImpl) NewTransactionalDAO() (dao.IDao, error) {
	t := &GormDAOImpl{
		DBConnector: dao.DBConnector,
		Logger:      dao.Logger,
	}
	err := t.Transactional()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (dao *GormDAOImpl) Transactional() error {
	dao.isTransactional = true
	return dao.beginTransactionIfNeeded()
}

func (dao *GormDAOImpl) IsTransactional() bool {
	return dao.isTransactional
}

func (dao *GormDAOImpl) beginTransactionIfNeeded() error {
	if dao.isTransactional && dao.tx == nil {
		txTemp := dao.DBConnector.GetConnection().Begin()
		if txTemp.Error != nil {
			return errors.Wrap(txTemp.Error, 0)
		}
		dao.tx = txTemp
	}
	return nil
}

func (dao *GormDAOImpl) RollbackCheck() {
	if dao.isTransactional && dao.tx != nil {
		dao.tx.Rollback()
		dao.tx = nil
	}
}

func (dao *GormDAOImpl) Commit() error {
	if dao.isTransactional {
		if dao.tx != nil {
			result := dao.tx.Commit()
			if result.Error != nil {
				dao.tx = nil
				return errors.Wrap(result.Error, 0)
			}
			dao.tx = nil
		} else {
			return errors.New("No transaction available")
		}
	}
	return nil
}

func (dao *GormDAOImpl) GetDemoDAO() interfaces.DemoDAOInterface {
	if dao.isTransactional {
		err := dao.beginTransactionIfNeeded()
		if err != nil {
			dao.Logger.Error("could not begin transaction", "error", errors.Wrap(err, 0))
			return nil
		}
		return &DemoDAOGormImpl{DB: dao.DBConnector.GetConnection(), Transactional: true}
	}
	return &DemoDAOGormImpl{DB: dao.DBConnector.GetConnection(), Transactional: true}
}
