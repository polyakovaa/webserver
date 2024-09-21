package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// base API server instance description
type API struct {
	//unexported field!!
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// API constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// start http server/configure loggers, router,database connection ...
func (api *API) Start() error {
	//trying tp configure logger
	if err := api.configureLoggerField(); err != nil {
		return err
	}
	api.logger.Info("starting api server at port: ", api.config.BindAddr)

	//Конфигурируем маршрутизатор
	api.configureRouterField()
	//на этапе валидноо завершения стартуем сервер
	return http.ListenAndServe(api.config.BindAddr, api.router)
}
