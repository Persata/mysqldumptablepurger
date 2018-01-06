package main

import (
	"os"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/Persata/mysqldumptablepurger/commands"
)

var (
	// App Declaration
	app = kingpin.New("mysqldumptablepurger", "A command line app to remove table data from MySQL dumps.")

	// List Command & Args
	list         = app.Command("list", "List all tables in the given MySQL dump.")
	listDumpFile = list.Arg("file", "MySQL dump file").Required().String()

	// Remove Command & Args
	remove         = app.Command("remove", "Remove the list of provided tables from the given MySQL dump")
	removeDumpFile = remove.Arg("inputFile", "MySQL dump file").Required().String()
	removeNewFile  = remove.Arg("outputFile", "MySQL dump file").Required().String()
	removeTables   = remove.Arg("tables", "Space delimited list of tables to remove").Required().Strings()
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	// List all tables in the given MySQL dump
	case list.FullCommand():
		commands.ListTables(*listDumpFile)

	// Make a copy of the given MySQL dump with the specified tables removed
	case remove.FullCommand():
		commands.RemoveTables(*removeDumpFile, *removeNewFile, *removeTables)

	}
}
