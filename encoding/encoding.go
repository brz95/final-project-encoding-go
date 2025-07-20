package encoding

import (
	"encoding/json"
	"fmt"
	"github.com/hive-bootcamp/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"os"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	file, err := os.ReadFile("jsonInput.json")
	if err != nil {
		panic(err)
	}
	var fJson models.DockerCompose
	err = json.Unmarshal(file, &fJson)
	if err != nil {
		panic(err)
	}
	out, err := yaml.Marshal(fJson)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	file, err := os.ReadFile("yamlInput.yml")
	if err != nil {
		panic(err)
	}

	var fYaml models.DockerCompose
	err = yaml.Unmarshal(file, &fYaml)
	if err != nil {
		panic(err)
	}

	data, err := json.MarshalIndent(fYaml, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
	return nil
}
