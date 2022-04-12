package timer

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/mribica/tmt/config"
)

type Timer struct {
	TimeUnit time.Duration
	Config   *config.Config
}

type handlerFunc func(int)

func NewTimer(timeUnit time.Duration, config *config.Config) Timer {
	return Timer{
		TimeUnit: timeUnit,
		Config:   config,
	}
}

func (t Timer) Start() {
	pomodoroLength := t.Config.PomodoroLength
	shortBreakLength := t.Config.ShortBreakLength
	longBreakLength := t.Config.LongBreakLength

	fmt.Printf("Pomodoro: %d min. | Short break: %d min. | Long break %d min.",
		pomodoroLength,
		shortBreakLength,
		longBreakLength,
	)

	completedPause := make(chan bool)
	completedPomodoro := make(chan bool)
	sigterm := make(chan os.Signal, 1)
	completedCounter := 0

	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)

	ExecuteCommand(t.Config.PomodoroCmd)
	NewTicker(t.TimeUnit, pomodoroLength, TickPomodoro, completedPomodoro)

	for {
		select {
		case <-completedPomodoro:
			completedCounter++
			if completedCounter%4 == 0 {
				ExecuteCommand(t.Config.LongBreakCmd)
				NewTicker(t.TimeUnit, longBreakLength, TickPause, completedPause)
			} else {
				ExecuteCommand(t.Config.ShortBreakCmd)
				NewTicker(t.TimeUnit, shortBreakLength, TickPause, completedPause)
			}
		case <-completedPause:
			ExecuteCommand(t.Config.PomodoroCmd)
			NewTicker(t.TimeUnit, pomodoroLength, TickPomodoro, completedPomodoro)
		case <-sigterm:
			fmt.Printf("\nBye!\n")
			fmt.Println("Completed pomodoros: ", completedCounter)
			os.Exit(0)
		}
	}
}

func NewTicker(period time.Duration, ticks int, tickHandler handlerFunc, completed chan bool) *time.Ticker {
	fmt.Printf("\n\a")

	ticks--
	tickHandler(ticks)
	ticker := time.NewTicker(period)

	go func() {
		for range ticker.C {
			if ticks == 0 {
				ticker.Stop()
				completed <- true
				return
			}
			ticks--
			tickHandler(ticks)
		}
	}()
	return ticker
}

func TickPomodoro(ticks int) {
	bar := strings.Repeat("⚫️", ticks)
	fmt.Printf("\r🍅 %s ", bar)
}

func TickPause(ticks int) {
	bar := strings.Repeat("⚫️", ticks)
	fmt.Printf("\r☕ %s ", bar)
}

func ExecuteCommand(c *config.Cmd) {
	if c != nil {
		go exec.Command(c.Cmd, c.Args...).Run()
	}
}
