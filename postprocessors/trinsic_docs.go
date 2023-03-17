package postprocessors

import (
	pgs "github.com/lyft/protoc-gen-star"
	"os"
	"regexp"
	"strings"
)

type TrinsicDocsPostProcessor struct {
	pgs.PostProcessor
	inputFile  string
	outputFile string
}

type CustomDocSampleFile struct {
	pgs.CustomTemplateFile

	SampleFile string
}

func (dpp TrinsicDocsPostProcessor) Match(a pgs.Artifact) bool {
	switch a.(type) {
	case CustomDocSampleFile:
		return true
	default:
		return false
	}
}

func (dpp TrinsicDocsPostProcessor) Process(in []byte) ([]byte, error) {
	// Determine which file it is
	templateFileString := string(in)
	newLines := strings.Split(templateFileString, "\n")

	r, err := regexp.Compile("target: ([\\w\\d\\-\\\\\\/\\.\\:]*)")
	if err != nil {
		panic(err)
	}
	matches := r.FindStringSubmatch(templateFileString)
	// Write the generated data to the appropriate file
	err = appendTargetFile(matches[1], newLines)
	if err != nil {
		return nil, err
	}

	return []byte{}, nil
}

func appendTargetFile(targetPath string, appendLines []string) error {
	//fmt.Fprintf(os.Stderr, "Target file: %s\n", targetPath)
	fileBytes, err := os.ReadFile(targetPath)
	if err != nil {
		return err
	}
	fileLines := strings.Split(string(fileBytes), "\n")
	newLines := append(fileLines, appendLines...)

	outputString := strings.Join(newLines, "\n")
	//fmt.Fprintln(os.Stderr, outputString)
	err = os.WriteFile(targetPath, []byte(outputString), 0777)
	if err != nil {
		return err
	}

	return nil
}

func AppendSampleFiles() pgs.PostProcessor { return TrinsicDocsPostProcessor{} }
