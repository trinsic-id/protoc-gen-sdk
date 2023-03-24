package postprocessors

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"regexp"
	"strings"
)

type TrinsicDocsPostProcessor struct {
	pgs.PostProcessor
}

func (dpp TrinsicDocsPostProcessor) Match(a pgs.Artifact) bool {
	switch a.(type) {
	case pgs.CustomTemplateFile:
		ctf := a.(pgs.CustomTemplateFile)
		switch ctf.Data.(type) {
		case *TrinsicDocTemplateArtifact:
			return true
		default:
			return false
		}
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
	err = UpdateTargetFile(matches[1], newLines)
	if err != nil {
		return nil, err
	}

	return []byte{}, nil
}

func AppendSampleFiles() pgs.PostProcessor { return TrinsicDocsPostProcessor{} }
