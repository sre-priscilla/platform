package app

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	v1 "platform/internal/http/v1"
)

func RunHttpServer(cmd *cobra.Command, args []string) {
	port := viper.GetUint32("http.port")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1group := e.Group("/api/v1")
	v1.RegisterUserGroup(v1group.Group("/users"))

	address := fmt.Sprintf("0.0.0.0:%d", port)
	e.Logger.Fatal(e.Start(address))
}

func NewHttpCommand() *cobra.Command {
	return &cobra.Command{Use: "http", Run: RunHttpServer}
}
