package cfg

import (
	"log/slog"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

var logLevelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
}

// AppConfig from config yml
type AppConfig struct {
	Project *Project `yaml:"project"`
	Log     *Log     `yaml:"log"`
	Server  *Server  `yaml:"server"`
	Theme   *Theme   `yaml:"theme"`
}

// Project contains the project information
type Project struct {
	Name        string `yaml:"name"`
	Version     string `yaml:"version"`
	Debug       bool   `yaml:"debug"`
	Environment string `yaml:"environment"`
}

// Log - log parameters
type Log struct {
	RawLevelCode    string `yaml:"level"`
	RawLevel        slog.Level
	RawOutputFormat string `yaml:"output_format"`
	RawWriteToFile  bool   `yaml:"write_to_file"`
	RawPath         string `yaml:"path"`
}

// Server contains server configuration
type Server struct {
	RawHost              string        `yaml:"host"`
	RawPort              int           `yaml:"port"`
	RawReadHeaderTimeout time.Duration `yaml:"read_header_timeout"`
	RawWriteTimeout      time.Duration `yaml:"write_timeout"`
	RawIdleTimeout       time.Duration `yaml:"idle_timeout"`
}

// Theme contains theme configuration
type Theme struct {
	RawThemesDir string `yaml:"themes_dir"`
	RawTheme     string `yaml:"theme"`
}

// Load config from file
func Load(filename string) (*AppConfig, error) {
	res := &AppConfig{}
	data, err := os.ReadFile(filename) // #nosec G304
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, res); err != nil {
		return nil, err
	}

	level, ok := logLevelMap[res.Log.RawLevelCode]
	if ok {
		res.Log.RawLevel = level
	} else {
		res.Log.RawLevel = slog.LevelInfo
	}

	return res, nil
}
