package main

import (
	"log"

	"github.com/MakMoinee/go-mith/pkg/goserve"
	"github.com/MakMoinee/gojailms/cmd/webapp/config"
	"github.com/MakMoinee/gojailms/cmd/webapp/routes"
	"github.com/MakMoinee/gojailms/internal/common"
)

func main() {

	config.Set()
	httpService := goserve.NewService(common.SERVER_PORT)
	httpService.EnableProfiling(common.SERVER_ENABLE_PROFILING)
	routes.Set(httpService)
	log.Println("Server Starting in Port ", common.SERVER_PORT)
	if err := httpService.Start(); err != nil {
		panic(err)
	}
}
