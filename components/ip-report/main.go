package main

import (
	"fmt"
	"net/http"

	"github.com/codingconcepts/env"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Bind     string `env:"BIND"`
	Endpoint string `env:"ENDPOINT"`
}

func main() {
	fmt.Println("vim-go")
	config := &Config{}
	if err := env.Set(&config); err != nil {
		logrus.Fatal(err)
	}

	logrus.Debugf("config: %+v", config)

	app := gin.Default()

	app.POST("/*", func(c *gin.Context) {
		msg := &struct {
			Text string
		}{}
		err := c.ShouldBind(msg)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		logrus.Info(msg.Text)
		c.Status(http.StatusOK)
	})

	if err := app.Run(config.Bind); err != nil {
		panic(err)
	}
}
