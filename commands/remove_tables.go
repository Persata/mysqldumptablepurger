package commands

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"github.com/persata/mysqldumptablepurger/io"
	"github.com/persata/mysqldumptablepurger/parser"
	"regexp"
)

func RemoveTables(inputPath string, outputPath string, tables []string) {
	fmt.Println("The following tables will be removed: \n")

	for _, table := range tables {
		fmt.Println(fmt.Sprintf("- %s", table))
	}

	if _, err := os.Stat(outputPath); err == nil {
		fmt.Println(fmt.Sprintf("\nThe specified output file '%s' already exists, and will be overwritten!", outputPath))
	}

	fmt.Println("\nContinue? Y/n")

	reader := bufio.NewReader(os.Stdin)
	c, _ := reader.ReadString('\n')

	c = strings.TrimSpace(c)

	if strings.ToLower(c) != "y" {
		fmt.Println("Exiting, no action performed")
		return
	} else {
		fmt.Println("Processing...\n")
	}

	s, fr := io.GetScanner(inputPath)
	defer fr.Close()

	w, fw := io.GetWriter(outputPath)
	defer fw.Close()

	replaceRegex := regexp.MustCompile(fmt.Sprintf(parser.TableStructureRegexReplace, strings.Join(tables, "|")))

	skipping := false

	matchCount := 0

	for s.Scan() {
		currentLine := s.Text()

		match := replaceRegex.FindStringSubmatch(currentLine)

		if len(match) > 0 {
			matchCount += 1
			skipping = true
			fmt.Println(fmt.Sprintf("Found section for `%s`, skipping data", match[1]))

			w.WriteString(fmt.Sprintf("-- Section for `%s` removed by mysqldumptablepurger", match[1]))
			w.WriteString("\n")
		}

		if !skipping {
			w.WriteString(currentLine)
			w.WriteString("\n")
		}

		if skipping {
			if strings.HasPrefix(currentLine, parser.EndSkippingPrefix) {
				skipping = false
			}
		}
	}

	w.Flush()

	fmt.Println(fmt.Sprintf("\nExiting, %d tables removed", matchCount))
}
