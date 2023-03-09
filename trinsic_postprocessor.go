package main

import (
	"fmt"
	pgs "github.com/lyft/protoc-gen-star"
	"os"
	"path"
	"regexp"
	"strings"
)

type TrinsicPostProcessor struct {
	pgs.PostProcessor
	inputFile  string
	outputFile string
}

func (tpp TrinsicPostProcessor) Match(a pgs.Artifact) bool {
	switch a.(type) {
	case pgs.CustomTemplateFile:
		return true
	default:
		return false
	}
}

func (tpp TrinsicPostProcessor) Process(in []byte) ([]byte, error) {
	// Determine which file it is
	templateFileString := string(in)
	templateLines := strings.Split(templateFileString, "\n")

	r, err := regexp.Compile("target: ([\\w\\d\\-\\\\\\/\\.\\:]*)")
	if err != nil {
		panic(err)
	}
	matches := r.FindStringSubmatch(templateFileString)
	//fmt.Fprintln(os.Stderr, matches)
	//fmt.Fprintln(os.Stderr, templateLines)
	// Write the generated data to the appropriate file
	err = tpp.updateTargetFile(matches[1], templateLines)
	if err != nil {
		return nil, err
	}

	return []byte{}, nil
}

func (tpp TrinsicPostProcessor) updateTargetFile(targetPath string, templateLines []string) error {
	//fmt.Fprintf(os.Stderr, "Target file: %s\n", targetPath)
	fileBytes, err := os.ReadFile(targetPath)
	if err != nil {
		return err
	}

	firstLine := strings.TrimSpace(templateLines[0])
	lastLine := strings.TrimSpace(templateLines[len(templateLines)-1])

	//fmt.Fprintln(os.Stderr, firstLine, lastLine)

	fileLines := strings.Split(string(fileBytes), "\n")
	newLines := make([]string, 0)

	for idx, line := range fileLines {
		testLine := strings.TrimSpace(line)
		//fmt.Fprintln(os.Stderr, testLine+" :: "+firstLine)
		if strings.HasPrefix(testLine, firstLine) {
			//fmt.Fprintln(os.Stderr, fileLines[0:idx])
			newLines = append(newLines, fileLines[0:idx]...)
			newLines = append(newLines, templateLines...)
		} else if strings.HasPrefix(testLine, lastLine) {
			//fmt.Fprintln(os.Stderr, fileLines[idx+1:])
			newLines = append(newLines, fileLines[idx+1:]...)
		}
	}

	// TODO - Refactor this behavior?
	if path.Ext(targetPath) == ".md" {
		fmt.Fprintf(os.Stderr, "Correcting ``` escapement in file:%s\n", targetPath)
		// Handle the ``` issue
		for idx, line := range newLines {
			newLines[idx] = strings.ReplaceAll(line, "'''", "```")
		}
	}

	// If we didn't find anything, don't write anything
	if len(newLines) == 0 {
		fmt.Fprintln(os.Stderr, "No target comments found for file: "+targetPath)
		return nil
	}

	outputString := strings.Join(newLines, "\n")
	//fmt.Fprintln(os.Stderr, outputString)
	err = os.WriteFile(targetPath, []byte(outputString), 0777)
	if err != nil {
		return err
	}

	return nil
}

func applyTemplateFiles() pgs.PostProcessor { return TrinsicPostProcessor{} }
