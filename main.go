package main

import (
	"github.com/spf13/viper"
	_ "go-gin/config"
	"go-gin/controller"
	"go-gin/tasks"
)

func main() {
	if viper.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer controller.Close()
	controller.ServerRun()
}
