package infra

import (
	"e-commerce-app/database"
	"e-commerce-app/handler"
	"e-commerce-app/repository"
	"e-commerce-app/service"
	"e-commerce-app/util"
)

type Context struct {
	Config  util.Configuration
	Handler handler.AllHandler
}

func NewContext() (Context, error) {
	config, err := util.ReadConfig()
	if err != nil {
		return Context{}, err
	}

	logger, err := util.LoggerInit(config)
	if err != nil {
		return Context{}, err
	}

	db, err := database.InitDB(config)
	if err != nil {
		return Context{}, err
	}

	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	handler := handler.NewAllHandler(service, logger)
	return Context{Config: config, Handler: handler}, nil
}
