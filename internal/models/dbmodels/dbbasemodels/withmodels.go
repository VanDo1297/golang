package dbbasemodels

type WithUserModel struct {
	UserID  	uint64 				`gorm:"column:user_id;not null" json:"-"`
}

type WithUniqueUserIDModel struct {
	UserID  	uint64 				`gorm:"column:user_id;not null;unique" json:"-"`
}

type WithUniqueTeamID struct {
	TeamID  	uint64 				`gorm:"column:team_id;not null;unique" json:"-"`
}

type WithTeamID struct {
	TeamID  	uint64 				`gorm:"column:team_id;not null" json:"-"`
}

type WithStripeIDModel struct {
	StripeID    string    			`gorm:"column:stripe_id;type:text;not null" json:"-"`
}
