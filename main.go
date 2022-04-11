package main

import (
	"embed"
	_ "embed"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wh3r3areyou/goswagg/internal/generators/client"
	"github.com/wh3r3areyou/goswagg/internal/generators/models"
	"github.com/wh3r3areyou/goswagg/internal/generators/server"
	"github.com/wh3r3areyou/goswagg/internal/tag"
	"github.com/wh3r3areyou/goswagg/internal/tools"
	"log"
	"path/filepath"
)

//go:embed swagger-templates/*
var templates embed.FS

func main() {
	var (
		pathFile string
	)

	var commandGen = &cobra.Command{
		Use:   "generate",
		Short: "generate app from swagger.yaml",
		Long:  ``,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			makeApp(pathFile)
		},
	}

	commandGen.Flags().StringVarP(&pathFile, "file", "f", "", "path to swagger.yaml")

	err := commandGen.Execute()

	if err != nil {
		log.Fatal(err)
	}

}

// getFileAnDir get dir and filename swagger.yaml
func getFileAndDir(filePath string) (string, string) {
	dir, err := filepath.Abs(filepath.Dir(filePath))

	if err != nil {
		log.Fatal("not find file")
	}

	file := filepath.Base(filePath)

	return dir, file
}

// getSwagger Get configuration swagger from dir
func getSwagger(dir string, file string) {
	extension := filepath.Ext(file)
	nameFile := file[0 : len(file)-len(extension)]
	viper.AddConfigPath(dir)
	viper.SetConfigName(nameFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("errors with swagger docs")
		log.Fatal(err)
	}
}

// makeApp Generate application from swagger.yaml
func makeApp(pathFile string) {
	dir, file := getFileAndDir(pathFile)

	getSwagger(dir, file)

	tags := tag.GetTags()

	models.NewModelsGenerator().Generate(&templates, tag.Tag{}, dir, file)

	for _, tag := range tags {
		client.NewClientGenerator().Generate(&templates, tag, dir, file)
	}
	server.NewServerGenerator().Generate(&templates, tags[0], dir, file)

	tools.MakeTools(dir)
}
