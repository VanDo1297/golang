package dbbasemodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBInternalModel struct {
	ID        uint64    `gorm:"primary_key;AUTO_INCREMENT" json:"-"`
	CreatedAt time.Time `gorm:"type:datetime;not null" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:datetime;not null" json:"updatedAt"`
}

type DBModel struct {
	DBInternalModel
	Unid string `gorm:"column:unid;type:varchar(50);unique;not null" json:"unid"`
}

func (dbModel *DBModel) BeforeCreate(tx *gorm.DB) (err error) {
	dbModel.Unid = uuid.NewString()
	dbModel.CreatedAt = time.Now().UTC()
	dbModel.UpdatedAt = time.Now().UTC()

	return nil
}

func (dbModel *DBInternalModel) BeforeCreate(tx *gorm.DB) (err error) {
	dbModel.CreatedAt = time.Now().UTC()
	dbModel.UpdatedAt = time.Now().UTC()

	return nil
}
