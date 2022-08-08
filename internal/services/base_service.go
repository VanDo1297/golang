package services

import (
	"demo/internal/daos/dao"
	"github.com/go-errors/errors"
)

type BaseService interface {
	WithTransactionalDAO(dao dao.IDao) error
}

type DefaultBaseService struct {
	dao           dao.IDao
}

func (baseService *DefaultBaseService) WithTransactionalDAO(dao dao.IDao) error{
	if !dao.IsTransactional(){
		return errors.New("dao needs to be transactional")
	}
	baseService.dao = dao
	return nil
}


