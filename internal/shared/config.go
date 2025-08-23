package shared

import (
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	Port     string
	Database struct {
		Path string
	}
	Log struct {
		Level string
	}
}

func Load() Config {
	c := Config{}

	c.Port = os.Getenv("PORT")
	if c.Port == "" {
		c.Port = "8080"
	}

	c.Database.Path = os.Getenv("DB_PATH")
	if c.Database.Path == "" {
		c.Database.Path = "./data.db"
	}

	c.Log.Level = os.Getenv("LOG_LEVEL")
	if c.Log.Level == "" {
		c.Log.Level = "info"
	}

	return c
}

func (c *Config) GetLogLevel() slog.Level {
	switch strings.ToLower(c.Log.Level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
