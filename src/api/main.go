package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RealSnowKid/ResIT/config"
	"github.com/RealSnowKid/ResIT/repository"
	httpRouter "github.com/RealSnowKid/ResIT/router/http"
)


func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	repository.Init()
	httpRouter.Init()
	// defer repository.DB.Close()
}
