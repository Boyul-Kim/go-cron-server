package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type gin_router struct {
	router *gin.Engine
}

type CronEdit struct {
	CronTime string `json:"crontime"`
	CronType string `json:"crontype"`
	Key      string `json:"key"`
}

type CronStop struct {
	CronTime string `json:"crontime"`
	CronType string `json:"crontype"`
	Key      string `json:"key"`
}

func setupRouter() *gin_router {
	// Disable Console Color
	// gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.PUT("/cron/time/edit", func(c *gin.Context) {

		var bindData CronEdit

		if err := c.BindJSON(&bindData); err != nil {
			fmt.Println("ERROR")
		}

		if bindData.Key != apiKey {
			fmt.Println("ERROR: Wrong key")
			return
		}

		switch bindData.CronType {
		case "cronEX":
			fmt.Println("TWITTER EDIT TRIGGERED")
			ch := make(chan cron_scheduler)
			go cron_ex.adjustCronTime(bindData.CronTime, ch)
		default:
			fmt.Println("NO CRON TYPES FOUND")
		}

		c.JSON(200, gin.H{
			"message": "edit completed",
		})
	})

	r.POST("/cron/stop", func(c *gin.Context) {

		var bindData CronStop

		if err := c.BindJSON(&bindData); err != nil {
			fmt.Println("ERROR")
		}

		if bindData.Key != apiKey {
			fmt.Println("ERROR: Wrong key")
			return
		}

		switch bindData.CronType {
		case "cronEX":
			go cron_ex.stopCron()
		default:
			fmt.Println("NO CRON TYPES FOUND")
		}
		c.JSON(200, gin.H{
			"message": "selected cron service stopped",
		})
	})

	ginRouter := gin_router{
		router: r,
	}

	return &ginRouter
}
