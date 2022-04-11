package tools

import (
	"log"
	"os"
)

type ToolFile struct {
	nameFile string
	bodyFile string
}

// MakeTools Make tools for APP
func MakeTools(dir string) {
	files := []ToolFile{
		{
			bodyFile: DockerCompose,
			nameFile: "docker-compose.yml",
		},
		{
			bodyFile: DockerIgnore,
			nameFile: ".dockerignore",
		},
		{
			bodyFile: Env,
			nameFile: ".env.example",
		},
		{
			bodyFile: Makefile,
			nameFile: "Makefile",
		},
		{
			bodyFile: GitIgnore,
			nameFile: ".gitignore",
		},
		{
			bodyFile: Config,
			nameFile: "/configs/config.yml",
		},
		{
			bodyFile: DockerDev,
			nameFile: "/build/dev/Dockerfile",
		},
		{
			bodyFile: DockerProd,
			nameFile: "/build/prod/Dockerfile",
		},
	}

	makeDockerDirectories(dir)
	makeConfigDir(dir)
	makeFiles(dir, files)
}

// makeConfigDir make configs dir in main directory
func makeConfigDir(dir string) {
	err := os.MkdirAll(dir+"/configs", 0755)

	if err != nil {
		log.Println("can`t make dir configs")
		log.Fatal(err)
	}
}

// Make build dir and prod/dev dirs
func makeDockerDirectories(dir string) {
	err := os.MkdirAll(dir+"/build", 0755)

	if err != nil {
		log.Println("can`t make dir configs")
		log.Fatal(err)
	}

	err = os.MkdirAll(dir+"/build/dev", 0755)

	if err != nil {
		log.Println("can`t make dir configs")
		log.Fatal(err)
	}

	err = os.MkdirAll(dir+"/build/prod", 0755)

	if err != nil {
		log.Println("can`t make dir configs")
		log.Fatal(err)
	}
}

// makeFiles make tools
func makeFiles(dir string, toolFiles []ToolFile) {
	for _, file := range toolFiles {
		f, err := os.Create(dir + "/" + file.nameFile)

		log.Println("Make tool file: " + file.nameFile + " " + dir + "/" + file.nameFile)

		if err != nil {
			log.Fatal(err)
		}

		_, err = f.WriteString(file.bodyFile)

		if err != nil {
			log.Fatal(err)
		}

		err = f.Close()
		if err != nil {
			log.Println(err)
			continue
		}
	}

}
