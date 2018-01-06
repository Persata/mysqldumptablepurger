package commands_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Persata/mysqldumptablepurger/commands"
)

func TestGetConfigFileDataSingle(t *testing.T) {
	configFileData := commands.GetConfigFileData("../test/conf/test_single_table.yml")
	assert.Equal(t, configFileData.OutputFile, "test_dump_cleaned.sql")
	assert.Len(t, configFileData.Tables, 1)
}


func TestGetConfigFileDataTwo(t *testing.T) {
	configFileData := commands.GetConfigFileData("../test/conf/test_two_tables.yml")
	assert.Equal(t, configFileData.OutputFile, "test_dump_cleaned.sql")
	assert.Len(t, configFileData.Tables, 2)
}
