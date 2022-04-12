package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Cmd struct {
	Cmd  string
	Args []string
}

type Config struct {
	PomodoroLength   int
	ShortBreakLength int
	LongBreakLength  int
	PomodoroCmd      *Cmd
	ShortBreakCmd    *Cmd
	LongBreakCmd     *Cmd
}

var ErrReadingConfig = errors.New("error while reading config")
var ErrParsingConfig = errors.New("error while parsing config")

func defaultConfig() Config {
	return Config{
		PomodoroLength:   25,
		ShortBreakLength: 5,
		LongBreakLength:  15,
	}
}

func Load() (*Config, error) {
	config := defaultConfig()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return &config, err
	}

	configPath := filepath.Join(homeDir, ".tmt", "config.json")
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return &config, fmt.Errorf("%s:%w", configPath, ErrReadingConfig)
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return &config, fmt.Errorf("%v:%w", err, ErrParsingConfig)
	}

	return &config, nil
}
