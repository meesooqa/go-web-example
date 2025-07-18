package cfg

import "github.com/meesooqa/cfg"

// AppConfig from config yml
type AppConfig struct {
	cfg.AppConfig
	Log    *Log    `yaml:"log"`
	Server *Server `yaml:"server"`
	Theme  *Theme  `yaml:"theme"`
}

// Load config from file
func Load(filename string) (*AppConfig, error) {
	res := &AppConfig{}
	err := cfg.LoadInto(filename, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
