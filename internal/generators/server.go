package generators

import (
	"embed"
	"github.com/go-openapi/analysis"
	"github.com/go-swagger/go-swagger/generator"
	"path"
)

type ServerGenerator struct {
}

func NewServerGenerator() Generator {
	return &ServerGenerator{}
}

func (s *ServerGenerator) Generate(options GenerateOpts) {
	tagsNames := GetTagsNames(options.tags)
	options.template = s.parseTemplates(options.template)
	opts := s.setOpts(tagsNames, options.dir, options.file)
	err := opts.EnsureDefaults()
	if err != nil {
		panic(err)
	}
	err = generator.GenerateServer("", nil, nil, opts)
	if err != nil {
		panic(err)
	}
}

func (s *ServerGenerator) setOpts(tagsNames []string, dir string, file string) *generator.GenOpts {
	return &generator.GenOpts{
		Spec:              file,
		Target:            dir,
		APIPackage:        "operations",
		ModelPackage:      "models",
		ServerPackage:     "restapi",
		ClientPackage:     "client",
		Principal:         "",
		DefaultScheme:     "http",
		IncludeModel:      false,
		IncludeValidator:  true,
		IncludeHandler:    false,
		IncludeParameters: false,
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
		Name:              "api",
		FlagStrategy:      "go-flags",
		CompatibilityMode: "modern",
		ExistingModels:    "",
		Copyright:         "",
		Sections: generator.SectionOpts{
			Application: []generator.TemplateOpts{
				{
					Name:       "service",
					Source:     "swagger-templates/templates/server/app.gotmpl",
					Target:     "./internal/app",
					FileName:   "app.go",
					SkipExists: false,
					SkipFormat: false,
				},
				{
					Name:       "main",
					Source:     "swagger-templates/templates/server/main.gotmpl",
					Target:     "./cmd/server",
					FileName:   "main.go",
					SkipExists: false,
					SkipFormat: false,
				},
				{
					Name:       "config",
					Source:     "swagger-templates/templates/server/config.gotmpl",
					Target:     "./internal/config",
					FileName:   "config.go",
					SkipExists: false,
					SkipFormat: false,
				},
				{
					Name:       "response",
					Source:     "swagger-templates/templates/server/response.gotmpl",
					Target:     "./internal/response",
					FileName:   "response.go",
					SkipExists: false,
					SkipFormat: false,
				},
				{
					Name:       "repositories",
					Source:     "swagger-templates/templates/server/repositories.gotmpl",
					Target:     "./internal/repositories",
					FileName:   "repositories.go",
					SkipExists: false,
					SkipFormat: false,
				},
				{
					Name:       "controllers",
					Source:     "swagger-templates/templates/server/controllers.gotmpl",
					Target:     "./internal/controllers",
					FileName:   "controllers.go",
					SkipExists: false,
					SkipFormat: false,
				},
				{
					Name:       "handler",
					Source:     "swagger-templates/templates/server/handler.gotmpl",
					Target:     "./internal/handler",
					FileName:   "handler.go",
					SkipExists: false,
					SkipFormat: false,
				},
				{
					Name:       "requests",
					Source:     "swagger-templates/templates/server/requests.gotmpl",
					Target:     "./internal/requests",
					FileName:   "requests.go",
					SkipExists: false,
					SkipFormat: false,
				},
			},
			Operations: []generator.TemplateOpts{},
			Models:     []generator.TemplateOpts{},
		},
	}
}

func (s *ServerGenerator) parseTemplates(templates *embed.FS) *embed.FS {
	assets, err := templates.ReadDir("swagger-templates/templates/server")
	if err != nil {
		panic(err)
	}

	for _, asset := range assets {
		data, err := templates.ReadFile(path.Join("swagger-templates/templates/server", asset.Name()))
		if err != nil {
			panic(err)
		}

		err = generator.AddFile(path.Join("swagger-templates/templates/server", asset.Name()), string(data))
		if err != nil {
			panic(err)
		}
	}

	return templates
}
