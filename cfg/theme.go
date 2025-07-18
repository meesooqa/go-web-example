package cfg

// Theme contains theme configuration
type Theme struct {
	RawThemesDir string `yaml:"themes_dir"`
	RawTheme     string `yaml:"theme"`
}

func (cfg *Theme) ThemesDir() string {
	return cfg.RawThemesDir
}

func (cfg *Theme) Theme() string {
	return cfg.RawTheme
}

func (cfg *Theme) ExtDir() string {
	// TODO config ext_dir
	return "ext"
}
