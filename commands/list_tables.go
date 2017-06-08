package commands

import (
	"fmt"
	"github.com/Persata/mysqldumptablepurger/io"
	"github.com/Persata/mysqldumptablepurger/parser"
)

func ListTables(path string) {
	s, f := io.GetScanner(path)
	defer f.Close()

	fmt.Println(fmt.Sprintf("Scanning for tables in %s...\n", path))

	for s.Scan() {
		match := parser.TableStructureRegex.FindStringSubmatch(s.Text())

		if len(match) > 0 {
			fmt.Println(match[1])
		}
	}

	if s.Err() != nil {
		fmt.Println(fmt.Sprintf("An error occurred scanning the file: %s", s.Err()))
	}

	fmt.Println("\nScan complete!")
}
