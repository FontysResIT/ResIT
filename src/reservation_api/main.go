package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/FontysResIT/ResIT/config"
	"github.com/FontysResIT/ResIT/repository"
	httpRouter "github.com/FontysResIT/ResIT/router/http"
	"github.com/FontysResIT/ResIT/util"
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
