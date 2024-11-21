package infra

import (
	"e-commerce-app/database"
	"e-commerce-app/handler"
	"e-commerce-app/repository"
	"e-commerce-app/service"
	"e-commerce-app/util"

	"go.uber.org/zap"
)

type Context struct {
	Log     *zap.Logger
	Config  util.Configuration
	Handler handler.AllHandler
}

func NewContext() (Context, error) {
	logger, err := util.LoggerInit()
	if err != nil {
		return Context{}, err
	}

	config, err := util.ReadConfig()
	if err != nil {
		return Context{
			Log: logger,
		}, err
	}

	db, err := database.InitDB(config)
	if err != nil {
		return Context{
			Log: logger,
		}, err
	}

	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	handler := handler.NewAllHandler(service, logger)
	return Context{Log: logger, Config: config, Handler: handler}, nil
}
