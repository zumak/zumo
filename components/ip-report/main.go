package main

import (
	"fmt"

	"github.com/codingconcepts/env"
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

	// run grpc server
}
