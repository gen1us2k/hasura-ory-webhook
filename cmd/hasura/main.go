package main

import (
	"log"
	"webhook/config"
	"webhook/hook"
)

func main() {
	c, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(hook.NewHook(c).Start())
}
