package responses

import (
	"time"

	"github.com/jinzhu/copier"
)

type Base struct {
	UnidBase
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

type UnidBase struct {
	Unid string `json:"unid" validate:"required"`
}

func FilterCopy(toValue interface{}, fromValue interface{}) {
	err := copier.Copy(toValue, fromValue)
	if err != nil {
		//sharedinstances.Get().GetLogger().Error("Could not filter copy","error",errors.Wrap(err,0).ErrorStack())
	}
}
