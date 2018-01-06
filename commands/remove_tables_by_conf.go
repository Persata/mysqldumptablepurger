package commands

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

type RemoveConf struct {
	OutputFile string   `yaml:"outputFile"`
	Tables     []string `yaml:"tables"`
}

func RemoveTablesByConf(inputPath string, configFilePath string) int {
	configFileBuffer, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		panic(err)
	}

	var configFileData RemoveConf

	err = yaml.Unmarshal(configFileBuffer, &configFileData)

	if err != nil {
		panic(err)
	}

	if len(configFileData.OutputFile) == 0 {
		fmt.Println(fmt.Sprintf("No output file found in the given config file - got '%s'.", configFileData.OutputFile))
		return 1
	}

	if len(configFileData.Tables) == 0 {
		fmt.Println("Specified config file does not contain any tables to remove.")
		return 1
	}

	return RemoveTables(inputPath, configFileData.OutputFile, configFileData.Tables)
}
