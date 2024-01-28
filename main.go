package main

import (
	"log"

	"github.com/Geovanny0401/bookmarks/cmd"
	"github.com/Geovanny0401/bookmarks/internal/config"
)

func main() {
	cfg, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	app := cmd.NewApp(cfg)
	app.Run()
}
