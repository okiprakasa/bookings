package config

import "html/template"

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool //False means the app is in development mode
	TemplateCache map[string]*template.Template
	err           error
}
