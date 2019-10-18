package main

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func run() {
	log.Println("Starting...")

	c := cron.New()
	// Add your own cron jobs:
	// c.AddFunc("* * * * * *", func() {
	//     log.Println("Run models.CleanAllTag...")
	//     models.CleanAllExample()
	// })

	c.Start()

	t1 := time.NewTimer(time.Second * 10) // trigger sending messages in 10 secs interval
	// blocking for...select loop, will block main thread
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
