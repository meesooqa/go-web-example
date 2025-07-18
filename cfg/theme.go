package cfg

// Theme contains theme configuration
type Theme struct {
	RawDir    string `yaml:"dir"`
	RawName   string `yaml:"name"`
	RawExtDir string `yaml:"ext_dir"`
}

func (cfg *Theme) Dir() string {
	return cfg.RawDir
}

func (cfg *Theme) Name() string {
	return cfg.RawName
}

func (cfg *Theme) ExtDir() string {
	return cfg.RawExtDir
}
