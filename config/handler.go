package config

func (cfg *Handler) TemplatesDir() string {
	return cfg.RawTemplatesDir
}

func (cfg *Handler) TemplateName() string {
	return cfg.RawTemplateName
}
