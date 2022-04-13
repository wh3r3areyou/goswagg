package generators

import (
	"embed"
	"github.com/go-swagger/go-swagger/generator"
	"github.com/wh3r3areyou/goswagg/internal/tag"
	"github.com/wh3r3areyou/goswagg/internal/tools"
	"sync"
)

type (
	Generator interface {
		Generate
		Parser
		Setter
	}

	// GenerateOpts Options for generate
	GenerateOpts struct {
		template *embed.FS
		tags     []tag.Tag
		dir      string
		file     string
	}

	// Generate templates in dir
	Generate interface {
		Generate(options GenerateOpts)
	}

	// Parser Parse templates in templates dir
	Parser interface {
		parseTemplates(templates *embed.FS) *embed.FS
	}

	// Setter set opts to generate
	Setter interface {
		setOpts(tagsNames []string, dir string, file string) *generator.GenOpts
	}
)

func GetTagsNames(tags []tag.Tag) []string {
	tagsNames := make([]string, 0, len(tags))
	for _, tag := range tags {
		tagsNames = append(tagsNames, tag.Name)
	}
	return tagsNames
}

// makeWG Generate work-pool for generate templates
func makeWG() *sync.WaitGroup {
	const countWorkers = 4
	var wg sync.WaitGroup
	wg.Add(countWorkers)
	return &wg
}

// Run Generate APP!
func Run(templates *embed.FS, tags []tag.Tag, dir string, file string) {
	wg := makeWG()
	mg := NewModelsGenerator()
	sg := NewServerGenerator()
	cg := NewClientGenerator()

	opts := GenerateOpts{
		file:     file,
		dir:      dir,
		template: templates,
		tags:     tags,
	}

	go func() {
		mg.Generate(opts)
		wg.Done()
	}()

	go func() {
		sg.Generate(opts)
		wg.Done()
	}()

	go func() {
		cg.Generate(opts)
		wg.Done()
	}()

	go func() {
		tools.MakeTools(opts.dir)
		wg.Done()
	}()

	wg.Wait()
}
