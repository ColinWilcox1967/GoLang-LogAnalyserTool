package main

import (
	"fmt"

	"flag"
)

const (
	log_analyser_version = "1.0"
	default_output_file = "./OUTPUT.TXT"
	default_keyword_file = "./KEYWORDS.TXT"
)

var (
	outputFile string
	keywordFile string
)

func showBanner() {
	fmt.Printf("Log Analyser (Version %s).\n", log_analyser_version)
	fmt.Println("(c) Colin Wilcox 2024.")
}

func parseCommandLine() {

	flag.StringVar(&keywordFile, "keywords", default_keyword_file, "Name of keyword file.")
	flag.StringVar(&outputFile, "output", default_output_file, "Name of output results file.")
	flag.Parse()

	if len(keywordFile) == 0 {
		keywordFile = default_keyword_file
	}

	fmt.Printf("Output File : '%s'.\n", outputFile)
	fmt.Printf("Keytword file : '%s'.\n", keywordFile)
}

func main {
	showBanner()

	parseCommandLine()
}