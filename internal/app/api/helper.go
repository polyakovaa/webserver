package api

import (
	"net/http"

	_ "github.com/gorilla/mux"
	"github.com/polyakovaa/standartserver/storage"
	"github.com/sirupsen/logrus"
)

// конфигурируем API инстанс( конкретнее логгер)
func (a *API) configureLoggerField() error {
	log_level, err := logrus.ParseLevel(a.config.LoggerLevel)
	if err != nil {
		return err
	}
	a.logger.SetLevel(log_level)
	return nil
}

// конфигурируем маршрутизатор
func (a *API) configureRouterField() {
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! I'm rest api"))
	})
}

// конфигурируем хранилище
func (a *API) configureStorageField() error {
	storage := storage.New(a.config.Storage)
	if err := storage.Open(); err != nil {
		return err
	}
	a.storage = storage
	return nil
}
