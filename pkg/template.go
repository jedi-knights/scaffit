package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func applyLineTemplate(line string, data map[string]string) string {
	for key, value := range data {
		line = strings.ReplaceAll(line, key, value)
	}

	return line
}

func applyTemplate(lines []string, data map[string]string) []string {
	var (
		result []string
	)

	for _, line := range lines {
		result = append(result, applyLineTemplate(line, data))
	}

	return result
}

// https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go

func applyTemplateFile(srcFile string, dstFile string, data map[string]string) error {
	var (
		err  error
		tmpl *template.Template
	)

	// Open and read the template inputFile
	inputFile, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// OPen the output inputFile for writing.
	outputFile, err := os.Create(dstFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Create bufio.Scanners for both input and output files.
	inputScanner := bufio.NewScanner(inputFile)
	outputWriter := bufio.NewWriter(outputFile)

	// Iterate over each line in the inputFile
	for inputScanner.Scan() {
		line := inputScanner.Text()

		// Process the line
		// Todo: treat the line as a templatefam

		_, err := outputWriter.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to outputFile: ", err)
			return err
		}
	}

	// Check for any errors that may have occurred during scanning.
	if err = inputScanner.Err(); err != nil {
		fmt.Println("Error scanning inputFile: ", err)
	}
	if err = outputWriter.Flush(); err != nil {
		fmt.Println("Error flushing outputFile: ", err)
	}

	// Create a new template and parse the template inputFile
	template, err := template.New("template").Parse(templateFile)

	if tmpl, err = template.ParseFiles(srcFile); err != nil {
		return err
	}

	if err = tmpl.ExecuteTemplate(dstFile, srcFile, data); err != nil {
		return err
	}

	return nil
}
