package commands_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/Persata/mysqldumptablepurger/commands"
)

func TestListTables(t *testing.T) {
	tables := commands.GetTablesList("../test/data/test_dump.sql")
	assert.Len(t, tables, 2)
	assert.Contains(t, tables, "test_table_1")
	assert.Contains(t, tables, "test_table_2")
}
