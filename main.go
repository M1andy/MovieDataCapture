package main

import (
	"MovieDataCapture/utils"
	"fmt"
)

func main() {
	for _, file := range utils.VideoList {
		if info, err := file.Info(); err == nil {
			fmt.Printf("file name: %s\n", info.Name())
		} else {
			fmt.Printf("%s", err)
		}
	}
}
