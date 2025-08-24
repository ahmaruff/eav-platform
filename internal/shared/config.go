package shared

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	Lifetime string // "24h"
	Name     string // "session_id"
	Secure   bool   // false for dev
}

type Log struct {
	Level      string
	File       string
	MaxSize    int // MB
	MaxAge     int // days
	MaxBackups int // number of old files
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

func getEnvInt(key string, defaultValue int) int {
	if str := os.Getenv(key); str != "" {
		if val, err := strconv.Atoi(str); err == nil {
			return val
		}
	}
	return defaultValue
}

func (c *Config) setupLogConfig() *Config {
	c.Log.Level = os.Getenv("LOG_LEVEL")
	if c.Log.Level == "" {
		c.Log.Level = "info"
	}

	c.Log.File = os.Getenv("LOG_FILE")
	if c.Log.File == "" {
		c.Log.File = "logs/app.log"
	}

	c.Log.MaxSize = getEnvInt("LOG_MAX_SIZE", 100)
	c.Log.MaxAge = getEnvInt("LOG_MAX_AGE", 14)
	c.Log.MaxBackups = getEnvInt("LOG_MAX_BACKUPS", 2)

	return c
}

func (c *Config) setupDatabaseConfig() *Config {
	c.Database.Path = os.Getenv("DB_PATH")
	if c.Database.Path == "" {
		c.Database.Path = "./data.db"
	}

	return c
}

func (c *Config) setupSessionConfig() *Config {
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

func LoadConfig() Config {
	c := Config{}

	c.setupLogConfig()
	c.setupDatabaseConfig()
	c.setupSessionConfig()

	c.Port = os.Getenv("PORT")
	if c.Port == "" {
		c.Port = "8080"
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
