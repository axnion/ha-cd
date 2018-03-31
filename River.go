package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jasonlvhit/gocron"
)

var lock bool

func main() {
	lock = false
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		deploy()
	})

	go startUpdateScheduler()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func startUpdateScheduler() {
	gocron.Every(10).Seconds().Do(update)
	<-gocron.Start()
}

func deploy() {
	if lock == false {
		lock = true
		fmt.Println("Deployment started...")
		time.Sleep(5 * time.Second)
		fmt.Println("Deployment end...")
		lock = false
	} else {
		fmt.Println("Can't perform deployment now")
	}
}

func update() {
	if lock == false {
		lock = false
		lock = true
		fmt.Println("Update started...")
		time.Sleep(10 * time.Second)
		fmt.Println("Update ended...")
		lock = false
	} else {
		fmt.Println("Can't perform update now")
	}
}
