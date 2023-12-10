package app

import "app/internal/config"

type App struct {
}

func NewApp(logger *logging.Logger, config *config.Config) (App, error) {
	app := App{}
	return app, nil

}
