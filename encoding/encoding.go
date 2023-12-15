package encoding

import (
	"encoding/json"
	"fmt"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
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

	jsonFile, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		return err
	}

	err = json.Unmarshal(jsonFile, j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации: %s", err.Error())
		return err
	}

	yamlData, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при сериализации: %s", err.Error())
		return err
	}

	d, err := os.Create("data.yaml")
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return err
	}
	defer d.Close()

	_, err = d.Write(yamlData)
	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {

	yamlFile, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s", err.Error())
		return err
	}

	err = yaml.Unmarshal(yamlFile, y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при десериализации: %s", err.Error())
		return err
	}
	jsonData, err := json.MarshalIndent(y.DockerCompose, "", "    ")
	fmt.Println(string(jsonData))

	if err != nil {
		fmt.Printf("ошибка при сериализации: %s", err.Error())
		return err
	}
	f, err := os.Create("data.json")
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s", err.Error())
		return err
	}
	defer f.Close()

	_, err = f.Write(jsonData)

	if err != nil {
		fmt.Printf("ошибка при записи данных в файл: %s", err.Error())
		return err
	}
	return nil
}
