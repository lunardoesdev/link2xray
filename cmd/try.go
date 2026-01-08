package main

import (
	"fmt"
	"os"

	"github.com/lunardoesdev/link2xray"
)

const exampleLink = "ss://Y2hhY2hhMjAtaWV0Zi1wb2x5MTMwNTpjdklJODVUclc2bjBPR3lmcEhWUzF1@45.87.175.187:8080#%3E%3E%40FreakConfig%3A%3AXX"

func main() {
	name, config, err := link2xray.SharedLinkToXrayConfig(exampleLink)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Configration found: %+v\n", name)
	fmt.Printf("%+v\n", config.OutboundConfigs[0])
}
