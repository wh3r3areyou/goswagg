package generators

import (
	"embed"
	"github.com/go-openapi/analysis"
	"github.com/go-swagger/go-swagger/generator"
	"path"
)

type ModelGenerator struct {
}

func NewModelsGenerator() Generator {
	return &ModelGenerator{}
}

func (m *ModelGenerator) Generate(options GenerateOpts) {
	tagsNames := GetTagsNames(options.tags)
	options.template = m.parseTemplates(options.template)
	opts := m.setOpts(tagsNames, options.dir, options.file)
	err := opts.EnsureDefaults()
	if err != nil {
		panic(err)
	}
	err = generator.GenerateModels([]string{}, opts)
	if err != nil {
		panic(err)
	}
}

func (m *ModelGenerator) setOpts(tagsNames []string, dir string, file string) *generator.GenOpts {
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
		Name:              "models",
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

func (m *ModelGenerator) parseTemplates(templates *embed.FS) *embed.FS {
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
