package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/FontysResIT/config_api/config"
	"github.com/FontysResIT/config_api/repository"
	httpRouter "github.com/FontysResIT/config_api/router/http"
	"github.com/FontysResIT/config_api/util"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	resItLogo := figure.NewFigure("ResIT", "", true)
	resItLogo.Print()
	config.Init(*environment)
	util.InitProducers()
	repository.Init()
	httpRouter.Init()

}
