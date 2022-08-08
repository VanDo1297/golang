package dbmodels

import "demo/internal/models/dbmodels/dbbasemodels"

type Demo struct {
	dbbasemodels.DBModel
	Message string `gorm:"type:text"`
}
