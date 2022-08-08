package interfaces

import "demo/internal/models/dbmodels"

type DemoDAOInterface interface {
	CreateDemo(demo *dbmodels.Demo) error
	GetDemo() ([]*dbmodels.Demo, error)
}
