package main

import (
	"context"
	"os"

	"grail-participant-registry/internal/app"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// main Initialises the App
func main() {
	setupLogger()

	cfg, err := bootstrap()
	if err != nil {
		logrus.Fatal(err)
	}

	a := app.New(cfg)

	err = a.Run()
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	os.Exit(0)
}

func bootstrap() (app.Config, error) {
	conf := app.Config{}

	viper.Set("Verbose", true)
	viper.Set("LogFile", os.Stdout)
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return conf, errors.Wrap(err, "unable to load .env file")
	}

	conf.Context = context.Background()
	conf.HTTPPort = viper.GetInt("HTTP_PORT")

	return conf, nil
}

func setupLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "level_name",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFile:  "caller_file",
			logrus.FieldKeyFunc:  "caller_func",
		},
	})
	logrus.SetLevel(logrus.DebugLevel)
}
