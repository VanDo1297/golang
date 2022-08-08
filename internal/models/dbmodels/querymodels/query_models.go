package querymodels

import "time"

type UserSubscription struct {
	Manager string
	ManagerSubscriptionID string

	Unid string
	Name string
	Amount float64
	Active bool
	Canceled bool
	CanceledEndsAt *time.Time
	NextCheck time.Time
}
