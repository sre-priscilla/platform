package main

import (
	"log"

	"github.com/spf13/viper"

	"platform/cmd/app"
)

func init() {
	viper.SetConfigType("yaml")
}

func main() {
	if err := app.NewRootCommand().Execute(); err != nil {
		log.Fatalln("execute root command:", err)
	}
}
