package app

import (
	"avito-test/internal/config"
	"avito-test/internal/handlers"
	"avito-test/internal/repository/database"
	pullrequest "avito-test/internal/repository/database/pull_request"
	"avito-test/internal/repository/database/team"
	"avito-test/internal/repository/database/user"
	"avito-test/internal/service"
	cLog "avito-test/log"
	"log"
	"os"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	log.Println("START APP")
	cfg := config.MustLoadConfig()

	logger := cLog.SetupLogger()
	logger.Info("logger init")

	db, err := database.InitDB(*cfg.Database)
	if err != nil {
		logger.Error("Cannot connect to db")
		os.Exit(1)
	}
	logger.Info("Connect to db")

	teamRepoSql := team.NewTeamRepoSQL(db, logger)
	logger.Info("Create user repo")

	userRepoSQL := user.NewUserRepoSQL(db, logger)
	logger.Info("Create user repo")

	pullRequestRepoSQL := pullrequest.NewPullRequestRepoSQL(db, logger)

	service := service.NewService(teamRepoSql, userRepoSQL, pullRequestRepoSQL, logger)
	logger.Info("Create service")

	handler := handlers.NewHandler(service)
	logger.Info("Create handler")

	srv := NewServer(*cfg, handler.InitRoutes())
	logger.Info("server data", "host", cfg.Server.Host, "port", cfg.Server.Port)

	err = srv.Run()
	if err != nil {
		logger.Error("Cannot start server")
		os.Exit(1)
	}
}
