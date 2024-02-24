package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"

	"flag"
)

const (
	log_analyser_version = "1.0"
	default_output_file = "./OUTPUT.TXT"
	default_keyword_file = "./KEYWORDS.TXT"
	default_log_file = "./LOGFILE.TXT"
)

var (
	outputFile string
	keywordFile string
	logFile string

	logLines []string
	keywords []string
)

func showBanner() {
	fmt.Printf("Log Analyser (Version %s).\n", log_analyser_version)
	fmt.Println("(c) Colin Wilcox 2024.")
}

func parseCommandLine() {

	flag.StringVar(&keywordFile, "keywords", default_keyword_file, "Name of keyword file.")
	flag.StringVar(&outputFile, "output", default_output_file, "Name of output results file.")
	flag.StringVar(&logFile, "log", default_log_file, "Name of log file to process.")

	flag.Parse()

	if len(keywordFile) == 0 {
		keywordFile = default_keyword_file
	}

	fmt.Printf("\nLog          : '%s'.\n", logFile)
	fmt.Printf("Output File  : '%s'.\n", outputFile)
	fmt.Printf("Keyword file : '%s'.\n\n", keywordFile)
}

func readLog() error {
		file, err := os.Open(logFile)
    if err != nil {
       return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
   
    for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			logLines = append(logLines, line)
		}
    }

    if err = scanner.Err(); err != nil {
       return err
    }

	return  nil
}

func readKeywordFile() error {

	return nil
}

func writeLineToOutputFile(line string) bool {
	return true
}

func scanLinesForKeywords() int {
	
	matchingLineCount:=0
	for lineNumber, line := range(logLines) {
		for _, word := range(keywords) {
			if strings.Contains(strings.ToLower(line), strings.ToLower(word)) {
				fmt.Printf("%04d:%s\n", lineNumber, line)
				writeLineToOutputFile(line)
				matchingLineCount++
			}
		}
	}

	fmt.Printf("\nMatching line count = %d.\n", matchingLineCount)
	return matchingLineCount
}


func main () {
	showBanner()

	parseCommandLine()

	err := readKeywordFile()
	if err != nil {
		fmt.Printf("*** Error : Problem reading keywords (%v).\n", err)
		os.Exit(-2)
	}

	err = readLog()
	if err != nil {
		fmt.Printf("*** Error : Problem reading logfile (%v).\n", err)
		os.Exit(-1)
	}

	scanLinesForKeywords()
}