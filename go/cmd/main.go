package main

import (
	"fmt"
	"log"

	"github.com/DmitrijP/frontendmasters-projector/pkg/config"
)

func main() {
	opts, err := config.GetOpts()
	if err != nil {
		log.Fatalf("Unable to get options: %v", err)
	}

	fmt.Printf("opts %+v", opts)

}
