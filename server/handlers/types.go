package handlers

type Config interface {
	TemplatesDir() string
	TemplateName() string
}
