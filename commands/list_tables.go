package commands

import (
	"fmt"
	"github.com/Persata/mysqldumptablepurger/iowrapper"
	"github.com/Persata/mysqldumptablepurger/parser"
)

func ListTables(path string) int {
	fmt.Println(fmt.Sprintf("Scanning for tables in %s...\n", path))

	tables := GetTablesList(path)

	for _, table := range tables {
		fmt.Println(table)
	}

	fmt.Println("\nScan complete!")

	return 0
}

func GetTablesList(path string) []string {
	s, f := iowrapper.GetScanner(path)
	defer f.Close()

	var tables []string

	for s.Scan() {
		match := parser.TableStructureRegex.FindStringSubmatch(s.Text())

		if len(match) > 0 {
			tables = append(tables, match[1])
		}
	}

	if s.Err() != nil {
		fmt.Println(fmt.Sprintf("An error occurred scanning the file: %s", s.Err()))
	}

	return tables
}
