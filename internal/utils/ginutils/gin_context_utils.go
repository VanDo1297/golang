package ginutils

import (
	"github.com/gin-gonic/gin"

	"demo/internal/models/dbmodels"
)

const (
	DemoKey = "demo"
)

func AddDemo(c *gin.Context, user *dbmodels.Demo) {
	c.Keys[DemoKey] = user
}

func GetDemo(c *gin.Context) *dbmodels.Demo {
	val, found := c.Keys[DemoKey]
	if found {
		return val.(*dbmodels.Demo)
	}
	return nil
}
