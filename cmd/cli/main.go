package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/mribica/tmt/config"
	"github.com/mribica/tmt/timer"
)

func main() {
	fmt.Println("🍅 tmt - Pomodoro tracker")

	fmt.Println(time.Now().Format("Monday 2006-January-1 15:04:05"))

	c, err := config.Load()
	if errors.Is(err, config.ErrParsingConfig) {
		fmt.Printf("Error while parsing config: %v\n", err)
	}

	if errors.Is(err, config.ErrReadingConfig) {
		fmt.Println("Starting with default configuration.")
	}

	t := timer.NewTimer(time.Minute, c)
	t.Start()
}
