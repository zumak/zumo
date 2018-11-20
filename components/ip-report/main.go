package main

import (
	"net/http"

	"github.com/codingconcepts/env"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Bind       string `env:"BIND" required:"true"`
	Endpoint   string `env:"ENDPOINT" required:"true"`
	CORSDomain string `env:"CORS_DOMAIN"`
	DBPath     string `env:"DBPATH" default:"ip-report.db"`
}

type IPInfo struct {
	Location string
	IP       string
}

func main() {
	config := &Config{}
	if err := env.Set(config); err != nil {
		logrus.Fatal(err)
	}

	logrus.Debugf("config: %+v", config)

	db, err := gorm.Open("sqlite3", config.DBPath)
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&IPInfo{}).Error; err != nil {
		panic(err)
	}

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		list := []IPInfo{}
		err := db.Find(&list).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, list)
	})
	app.POST("/", func(c *gin.Context) {
		info := &IPInfo{}
		err := c.ShouldBind(info)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		err = db.Assign(info).FirstOrCreate(info).Error
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.Status(http.StatusCreated)
	})
	app.POST("/hooks/*path", func(c *gin.Context) {
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
