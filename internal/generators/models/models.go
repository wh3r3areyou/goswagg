package models

import (
	"embed"
	"github.com/go-openapi/analysis"
	"github.com/go-swagger/go-swagger/generator"
	"github.com/wh3r3areyou/goswagg/internal/generators"
	"github.com/wh3r3areyou/goswagg/internal/tag"
	"path"
)

type Generator struct {
}

func NewModelsGenerator() generators.Generator {
	return &Generator{}
}

func (m *Generator) Generate(templates *embed.FS, tag tag.Tag, dir string, file string) {
	templates = m.ParseTemplates(templates)
	opts := m.SetOpts(tag, dir, file)
	err := opts.EnsureDefaults()
	if err != nil {
		panic(err)
	}
	err = generator.GenerateModels([]string{}, opts)
	if err != nil {
		panic(err)
	}
}

func (m *Generator) SetOpts(tag tag.Tag, dir string, file string) *generator.GenOpts {
	return &generator.GenOpts{
		Spec:              file,
		Target:            dir,
		APIPackage:        "operations",
		ModelPackage:      "models",
		ServerPackage:     "restapi",
		ClientPackage:     "client",
		Principal:         "",
		DefaultScheme:     "http",
		IncludeModel:      true,
		IncludeValidator:  true,
		IncludeHandler:    true,
		IncludeParameters: true,
		IncludeResponses:  true,
		IncludeURLBuilder: true,
		IncludeMain:       true,
		IncludeSupport:    true,
		ValidateSpec:      true,
		FlattenOpts: &analysis.FlattenOpts{
			Minimal:      true,
			Verbose:      true,
			RemoveUnused: false,
			Expand:       false,
		},
		ExcludeSpec:       false,
		TemplateDir:       "",
		DumpData:          false,
		Models:            nil,
		Operations:        nil,
		Tags:              []string{tag.Name},
		Name:              tag.Name,
		FlagStrategy:      "go-flags",
		CompatibilityMode: "modern",
		ExistingModels:    "",
		Copyright:         "",
		Sections: generator.SectionOpts{
			Models: []generator.TemplateOpts{
				{
					Name:       "models",
					Source:     "swagger-templates/templates/models/mod.gotmpl",
					Target:     "./pkg/models",
					FileName:   "{{ (snakize (pascalize .Name)) }}.go",
					SkipExists: true,
					SkipFormat: false,
				},
			},
		},
	}
}

func (m *Generator) ParseTemplates(templates *embed.FS) *embed.FS {
	assets, err := templates.ReadDir("swagger-templates/templates/models")
	if err != nil {
		panic(err)
	}

	for _, asset := range assets {
		data, err := templates.ReadFile(path.Join("swagger-templates/templates/models", asset.Name()))
		if err != nil {
			panic(err)
		}

		err = generator.AddFile(path.Join("swagger-templates/templates/models", asset.Name()), string(data))
		if err != nil {
			panic(err)
		}
	}

	return templates
}
