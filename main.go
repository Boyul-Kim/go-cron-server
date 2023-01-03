package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var cron_ex = cronInit("cronEX")
var apiKey string
var server string

// "EVERY_TWO_MINUTES"
// "EVERY_FIVE_MINUTES"
// "EVERY_30_MINUTES"
// "EVERY_DAY_TEN_PM"
// "EVERY_HOUR"
// "EVERY_MINUTE"
// "EVERY_MIDNIGHT",
// "EVERY_MONTH_AT_MIDNIGHT"

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	fmt.Println("STARTING CRON SERVER")

	apiKey = goDotEnvVariable("API_KEY")
	server = goDotEnvVariable("SERVER")

	ch := make(chan cron_scheduler)
	go cron_ex.startCron("EVERY_TWO_MINUTES", ch)
	fmt.Println("CRON INITIALIZED")

	router := setupRouter()
	fmt.Println("ROUTER", router)
	router.router.Run(":8080")
}
