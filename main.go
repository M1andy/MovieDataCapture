package main

import (
	"MovieDataCapture/utils"
	"fmt"
	"log"
)

func main() {
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("load config error: %v", err)
	}
	fmt.Printf("%v\n", config)

}
