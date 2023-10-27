package main

import (
	"fmt"
	"sync"
)

type MonaLisa struct {
	title       string
	artist      string
	description string
}

var instance *MonaLisa
var lock = &sync.Mutex{}

func PaintMonaLisa() *MonaLisa {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Painting Mona Lisa..")
			instance = &MonaLisa{title: "Mona Lisa", artist: "Leonardo Da Vinci", description: "Once created masterpiece"}
		} else {
			fmt.Println("Mona Lisa has been already painted")
		}
	} else {
		fmt.Println("Mona Lisa has been already painted")
	}
	return instance
}

func main() {
	for i := 0; i < 3; i += 1 {
		go PaintMonaLisa()
	}
	fmt.Scanln()
}
