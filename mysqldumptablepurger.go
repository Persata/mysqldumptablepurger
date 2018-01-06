package main

import (
	"github.com/Persata/mysqldumptablepurger/commands"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	// App Declaration
	app = kingpin.New("mysqldumptablepurger", "A command line app to remove table data from MySQL dumps.")

	// List Command & Args
	list         = app.Command("list", "List all tables in the given MySQL dump.")
	listDumpFile = list.Arg("file", "MySQL dump file").Required().String()

	// Remove Command & Args
	remove              = app.Command("remove", "Remove the list of provided tables from the given MySQL dump")
	removeDumpFile      = remove.Arg("inputFile", "MySQL dump file").Required().String()
	removeNewFile       = remove.Arg("outputFile", "MySQL dump file").Required().String()
	removeTables        = remove.Arg("tables", "Space delimited list of tables to remove").Required().Strings()

	// Remove by Config - Command & Args
	removeByConf           = app.Command("remove-by-conf", "Remove tables and output using a predefined YAML config file")
	removeByConfDumpFile   = removeByConf.Arg("inputFile", "MySQL dump file").Required().String()
	removeByConfConfigFile = removeByConf.Arg("configFile", "YAML configuration file").Required().String()
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case list.FullCommand():
		os.Exit(commands.ListTables(*listDumpFile))

	case remove.FullCommand():
		os.Exit(commands.RemoveTables(*removeDumpFile, *removeNewFile, *removeTables))

	case removeByConf.FullCommand():
		os.Exit(commands.RemoveTablesByConf(*removeByConfDumpFile, *removeByConfConfigFile))

	}

}
