package config

import "log/slog"

func (cfg *Log) Path() string {
	return cfg.RawPath
}

func (cfg *Log) OutputFormat() string {
	return cfg.RawOutputFormat
}

func (cfg *Log) Level() slog.Level {
	return cfg.RawLevel
}

func (cfg *Log) IsWriteToFile() bool {
	return cfg.RawWriteToFile
}
