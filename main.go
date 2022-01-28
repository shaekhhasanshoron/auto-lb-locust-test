package main

import (
	"github.com/locust-test/api"
	"github.com/locust-test/config"
)

func main(){
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}


