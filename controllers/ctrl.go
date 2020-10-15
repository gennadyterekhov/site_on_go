package controllers

import (
	"io/ioutil"
	"portfolio/helper"
)

type config struct {
	Port       int    `json:"PORT"`
	APIKey     string `json:"api_key"`
	WebhookUrl string `json:"webhook_url"`
	Cyrillic   string `json:"cyrillic"`
	Latin      string `json:"latin"`
}

func GetFile(filename string) string {
	configStr, err := ioutil.ReadFile(filename)
	helper.Check(err)

	return string(configStr)
}

func GetView(viewName string) string {
	return GetFile("views/" + viewName + ".html")
}
