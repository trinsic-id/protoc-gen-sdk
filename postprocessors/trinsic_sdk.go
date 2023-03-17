package postprocessors

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"regexp"
	"strings"
)

type TrinsicSdkPostProcessor struct {
	pgs.PostProcessor
}

func (tpp TrinsicSdkPostProcessor) Match(a pgs.Artifact) bool {
	switch a.(type) {
	case pgs.CustomTemplateFile:
		ctf := a.(pgs.CustomTemplateFile)
		switch ctf.Data.(type) {
		case *TrinsicSdkTemplateArtifact:
			return true
		default:
			return false
		}
	default:
		return false
	}
}

type CustomSdkTemplateFile struct {
	pgs.Artifact
	pgs.TemplateArtifact
	TargetFile string
}

func (tpp TrinsicSdkPostProcessor) Process(in []byte) ([]byte, error) {
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
	err = UpdateTargetFile(matches[1], templateLines)
	if err != nil {
		return nil, err
	}

	return []byte{}, nil
}

func ApplyTemplateFiles() pgs.PostProcessor { return TrinsicSdkPostProcessor{} }
