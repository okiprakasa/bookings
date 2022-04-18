package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// AppConfig holds the application config
type AppConfig struct {
	InProduction  bool //False means the app is in development mode
	TemplateCache map[string]*template.Template
	err           error
	Session       *scs.SessionManager
}
