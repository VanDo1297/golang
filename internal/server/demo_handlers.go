package server

import (
	"demo/internal/errors"
	"demo/internal/models/dbmodels"
	"demo/internal/models/responses"
	"demo/internal/utils/ginutils"

	"github.com/gin-gonic/gin"
)

func (server Server) GetDemo(c *gin.Context) (*responses.DemoResponses, error) {
	user := ginutils.GetDemo(c)
	userRes, err := server.DemoResponses(user)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return userRes, nil
}

func (server Server) DemoResponses(demo *dbmodels.Demo) (*responses.DemoResponses, error) {

	demos, err := server.Resources.GetDAO().GetDemoDAO().GetDemo()
	if err != nil {
		return nil, errors.Wrap(err)
	}

	var resDemos []*responses.Demo

	for _, demo := range demos {
		resDemo := new(responses.Demo)
		responses.FilterCopy(resDemos, demo)
		resDemos = append(resDemos, resDemo)
	}

	res := new(responses.DemoResponses)
	res.Data.Demos = resDemos
	return res, nil
}
