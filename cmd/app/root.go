package app

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewRootCommand() *cobra.Command {
	command := &cobra.Command{
		Use: "app",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			confPath, err := cmd.Flags().GetString("config")
			if err != nil {
				log.Fatalln("get config path:", err)
			}
			confData, err := ioutil.ReadFile(confPath)
			if err != nil {
				log.Fatalln("read config data:", err)
			}
			buffer := bytes.NewBuffer(confData)
			if err := viper.ReadConfig(buffer); err != nil {
				log.Fatalln("read config yaml:", err)
			}
		},
	}

	command.PersistentFlags().StringP(
		"config", "c", "config.yaml", "absolute path of config file")

	command.AddCommand(NewHttpCommand())

	return command
}
