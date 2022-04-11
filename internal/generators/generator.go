package generators

import (
	"embed"
	"github.com/go-swagger/go-swagger/generator"
	"github.com/wh3r3areyou/goswagg/internal/tag"
)

type (
	Generator interface {
		Generate
		Parser
		Setter
	}

	// Generate templates in dir
	Generate interface {
		Generate(templates *embed.FS, tag tag.Tag, dir string, file string)
	}

	// Parser Parse templates in templates dir
	Parser interface {
		ParseTemplates(templates *embed.FS) *embed.FS
	}

	// Setter set opts to generate
	Setter interface {
		SetOpts(tag tag.Tag, dir string, file string) *generator.GenOpts
	}
)
