package shared

import (
	"log/slog"
	"os"
	"strings"
	"time"
)

type Session struct {
	Lifetime string // "24h"
	Name     string // "session_id"
	Secure   bool   // false for dev
}

type Log struct {
	Level string
}

type Database struct {
	Path string
}

type Config struct {
	Port     string
	Database Database
	Log      Log
	Session  Session
}

func LoadConfig() Config {
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

	c.Session.Lifetime = os.Getenv("SESSION_LIFETIME")
	if c.Session.Lifetime == "" {
		c.Session.Lifetime = "24h"
	}

	c.Session.Name = os.Getenv("SESSION_NAME")
	if c.Session.Name == "" {
		c.Session.Name = "session_id"
	}

	secureStr := os.Getenv("SESSION_SECURE")
	if secureStr == "" {
		c.Session.Secure = false
	} else {
		c.Session.Secure = secureStr == "true"
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

func (c Config) GetSessionLifetime() time.Duration {
	d, err := time.ParseDuration(c.Session.Lifetime)
	if err != nil {
		return 24 * time.Hour
	}
	return d
}
