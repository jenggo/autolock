package main

import (
	"os"
	"strings"

	"autolock/app"
	"autolock/vars"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	vars.FileConf = strings.Join([]string{home, vars.FileConf}, "/")
	if err := cleanenv.ReadConfig(vars.FileConf, &vars.Config); err != nil {
		panic(err)
	}

	if err := app.Run(vars.Config.Path.Image); err != nil {
		panic(err)
	}
}
