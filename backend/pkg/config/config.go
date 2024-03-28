package config

import (
	"errors"
	"os"
	"strconv"
	"sync"

	"github.com/rs/zerolog/log"
)

type Config struct {
	URL   string
	Port  uint64
	Proxy bool
}

var (
	config = &Config{
		URL:   "https://localhost",
		Port:  80,
		Proxy: false,
	}
	mutex = &sync.RWMutex{}
)

func init() {
	err := loadFromEnv()
	if err != nil {
		log.Fatal().Err(err).Msg("[config-init-1] initialization failed")
	}
}

func loadFromEnv() error {
	URL, ok := os.LookupEnv("AP_URL")
	if ok {
		config.URL = URL
	}
	port, err := parseUintEnv("AP_PORT")
	if err != nil {
		return err
	}
	if port != 0 {
		config.Port = port
	}
	proxy, err := parseBoolEnv("AP_PROXY")
	if err == nil {
		config.Proxy = proxy
	}
	return nil
}

func parseUintEnv(envName string) (uint64, error) {
	valueStr, ok := os.LookupEnv(envName)
	if !ok {
		return 0, nil
	}
	value, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func parseBoolEnv(envName string) (bool, error) {
	valueStr, ok := os.LookupEnv(envName)
	if !ok {
		return false, errors.New("not set")
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return false, err
	}
	return value, nil
}

func Get() *Config {
	mutex.RLock()
	defer mutex.RUnlock()
	return config
}
