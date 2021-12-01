package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// We dont want to reload this cache every time a new template is being called
// or loaded in. To do this we need to set up an application wide
// configurations settings.
// The common approach is to create global variables; which isn't the way to go

// Why is this application wide config so usefull?
// -> We can add anything to this struct which will in turn be available to every
// part of the application that has access to this appConfig

// AppConfig holds the application config
// We can set any type we want in this stuct and it will be available to every script
// that import this package
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}