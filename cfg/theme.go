package cfg

func (cfg *Theme) ThemesDir() string {
	return cfg.RawThemesDir
}

func (cfg *Theme) Theme() string {
	return cfg.RawTheme
}
