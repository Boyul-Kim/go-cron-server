package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

type cron_scheduler struct {
	cron     *gocron.Scheduler
	cronType string
}

func cronInit(cronType string) *cron_scheduler {
	s := gocron.NewScheduler(time.UTC)

	cronScheduler := cron_scheduler{
		cron:     s,
		cronType: cronType,
	}

	return &cronScheduler
}

func cronFunc(cronType string, ch chan string) *http.Response {
	var responseBody *bytes.Buffer

	fmt.Println("CRON FUNC TRIGGERED ", cronType)
	postBody, _ := json.Marshal(map[string]string{
		"type": cronType,
	})
	responseBody = bytes.NewBuffer(postBody)

	if responseBody == nil {
		fmt.Println("NOTHING TRIGGERED")
		return nil
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", server+"cron/update", responseBody)
	req.Header.Add("apiKey", apiKey)
	resp, err := client.Do(req)
	// resp, err := http.Post(server+"cron/update", "application/json", responseBody)

	if err != nil {
		fmt.Println("An Error Occured trying to run cron func", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("An response body error", err)
	}
	sb := string(body)
	ch <- sb
	close(ch)
	return resp
}

func (c *cron_scheduler) startCron(time string, ch chan cron_scheduler) {

	switch time {
	case "EVERY_30_SECS":
		c.cron.Every(30).Seconds().Do(func() {
			ch := make(chan string)
			go cronFunc(c.cronType, ch)
			fmt.Println("CHANNEL RESULT", <-ch)
		})
		c.cron.StartAsync()
	case "EVERY_TWO_MINUTES":
		c.cron.Every(2).Minutes().Do(func() {
			ch := make(chan string)
			go cronFunc(c.cronType, ch)
			fmt.Println("CHANNEL RESULT", <-ch)
		})
		c.cron.StartAsync()
	case "EVERY_FIVE_MINUTES":
		c.cron.Every(5).Minutes().Do(func() {
			ch := make(chan string)
			go cronFunc(c.cronType, ch)
			fmt.Println("CHANNEL RESULT", <-ch)
		})
		c.cron.StartAsync()
	case "EVERY_MONTH_AT_MIDNIGHT":
		c.cron.Every(1).MonthLastDay().At("23:59").Do(func() {
			ch := make(chan string)
			go cronFunc(c.cronType, ch)
			fmt.Println("CHANNEL RESULT", <-ch)
		})
	default:
		fmt.Println("NO TYPES FOUND")
	}

}

func (c *cron_scheduler) stopCron() {
	fmt.Println("CANCELLING")
	c.cron.Stop()
}

func (c *cron_scheduler) adjustCronTime(time string, ch chan cron_scheduler) {
	fmt.Println("CHANGING TIME", time)
	c.cron.Clear()
	c.startCron(time, ch)
	fmt.Println("TIME CHANGED")
	close(ch)
}
