package generators

import (
	"embed"
	"github.com/go-openapi/analysis"
	"github.com/go-swagger/go-swagger/generator"
	"path"
)

type ClientGenerator struct {
}

func NewClientGenerator() Generator {
	return &ClientGenerator{}
}

func (c *ClientGenerator) Generate(options GenerateOpts) {
	for _, tag := range options.tags {
		tagsNames := GetTagsNames(options.tags)
		options.template = c.parseTemplates(options.template)
		opts := c.setOpts(tagsNames, options.dir, options.file)
		err := opts.EnsureDefaults()
		if err != nil {
			panic(err)
		}
		err = generator.GenerateClient(tag.Name, nil, nil, opts)
		if err != nil {
			panic(err)
		}
	}
}

func (c *ClientGenerator) setOpts(tagsNames []string, dir string, file string) *generator.GenOpts {
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
		Tags:              tagsNames,
		Name:              tagsNames[0],
		FlagStrategy:      "go-flags",
		CompatibilityMode: "modern",
		ExistingModels:    "",
		Copyright:         "",
		Sections: generator.SectionOpts{
			Application: []generator.TemplateOpts{
				{
					Name:       "service",
					Source:     "swagger-templates/templates/client/service.gotmpl",
					Target:     "./pkg/client/services",
					FileName:   "{{ snakize (pascalize .Name) }}_service.go",
					SkipExists: false,
					SkipFormat: false,
				},
				{
					Name:       "controllers",
					Source:     "swagger-templates/templates/client/controller.gotmpl",
					Target:     "./pkg/client/controllers",
					FileName:   "{{ snakize (pascalize .Name) }}_controllers.go",
					SkipExists: false,
					SkipFormat: false,
				},
				{
					Name:       "requests",
					Source:     "swagger-templates/templates/client/requests.gotmpl",
					Target:     "./pkg/client/requests",
					FileName:   "{{ snakize (pascalize .Name) }}_requests.go",
					SkipExists: false,
					SkipFormat: false,
				},
				{
					Name:       "repositories",
					Source:     "swagger-templates/templates/client/repositories.gotmpl",
					Target:     "./pkg/client/repositories",
					FileName:   "{{ snakize (pascalize .Name) }}_repositories.go",
					SkipExists: false,
					SkipFormat: false,
				},
			},
			Operations: []generator.TemplateOpts{
				{
					Name:       "parameters",
					Source:     "swagger-templates/templates/client/params.gotmpl",
					Target:     "./pkg/models/parameters",
					FileName:   "{{ snakize (pascalize .Name) }}_parameters.go",
					SkipExists: false,
					SkipFormat: false,
				},
			},
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

func (c *ClientGenerator) parseTemplates(templates *embed.FS) *embed.FS {
	assets, err := templates.ReadDir("swagger-templates/templates/client")
	if err != nil {
		panic(err)
	}

	for _, asset := range assets {
		data, err := templates.ReadFile(path.Join("swagger-templates/templates/client", asset.Name()))
		if err != nil {
			panic(err)
		}

		err = generator.AddFile(path.Join("swagger-templates/templates/client", asset.Name()), string(data))
		if err != nil {
			panic(err)
		}
	}

	return templates
}
