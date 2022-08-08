package dao

import (
	"demo/internal/daos/interfaces"

	"gorm.io/gorm"
)

type IDao interface {
	GetDemoDAO() interfaces.DemoDAOInterface

	WithTransaction(tx *gorm.DB)
	NewTransactionalDAO() (IDao, error)
	Transactional() error
	IsTransactional() bool
	RollbackCheck()
	Commit() error
}
