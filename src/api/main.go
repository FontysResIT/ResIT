package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RealSnowKid/ResIT/config"
	"github.com/RealSnowKid/ResIT/repository"
	httpRouter "github.com/RealSnowKid/ResIT/router/http"
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
	fmt.Println("")
	config.Init(*environment)
	repository.Init()
	httpRouter.Init()
}
