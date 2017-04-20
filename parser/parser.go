package parser

import "regexp"

var TableStructureRegex = regexp.MustCompile("-- Table structure for table `(.*)`")

var TableStructureRegexReplace = "^-- Dumping data for table `(%s)`"

var EndSkippingPrefix = "-- Table structure for table"