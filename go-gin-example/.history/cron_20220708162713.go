package main

import (
	"go-gin-example/models"
	"log"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
}
