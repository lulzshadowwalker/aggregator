package config

// questions
// 1. am i using correct file mode in os.write ?
// 2.  do I need to have a mutex or not ?

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"sync"
)

var Instance *Config

type Config struct {
	Database struct {
		URL string `json:"url"`
	} `json:"database"`
	Username string `json:"username"`

	mu sync.RWMutex `json:"-"`
}

func (c *Config) Read() error {
	bytes, err := os.ReadFile(filepath())
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	return json.Unmarshal(bytes, c)
}

func (c *Config) Write() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	bytes, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath(), bytes, 0600)
}

func init() {
	Instance = &Config{}
	if err := Instance.Read(); err != nil {
		panic(fmt.Errorf("failed to read configuration file: %w", err))
	}
}

func filepath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("failed to read user home directory %w", err))
	}

	return path.Join(dir, ".aggregator.json")
}
