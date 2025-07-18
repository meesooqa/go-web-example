package cfg

// Theme contains theme configuration
type Theme struct {
	RawThemesDir string `yaml:"themes_dir"`
	RawTheme     string `yaml:"theme"`
	RawExtDir    string `yaml:"ext_dir"`
}

func (cfg *Theme) ThemesDir() string {
	return cfg.RawThemesDir
}

func (cfg *Theme) Theme() string {
	return cfg.RawTheme
}

func (cfg *Theme) ExtDir() string {
	return cfg.RawExtDir
}
