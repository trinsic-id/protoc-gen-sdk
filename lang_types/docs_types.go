package lang_types

import (
	"fmt"
	pgs "github.com/lyft/protoc-gen-star/v2"
	"strings"
)

func DocsFileName(method pgs.Method) string {
	return fmt.Sprintf("%s_service_examples", method.Service().Name().LowerSnakeCase().String())
}
func DocsFileNameService(service pgs.Service) string {
	return fmt.Sprintf("%s_service_examples", service.Name().LowerSnakeCase().String())
}
func DocsMethodName(method pgs.Method) string {
	return fmt.Sprintf("%sService_%s", method.Service().Name().LowerCamelCase(), method.Name().UpperCamelCase().String())
}
func DocsMethodTab(method pgs.Method) string {
	// Strip the leading "."
	return method.FullyQualifiedName()[1:]
}
func DocsServiceTab(method pgs.Service) string {
	// Strip the leading "."
	return method.FullyQualifiedName()[1:]
}
func DocsCreateService(service pgs.Service) string {
	return fmt.Sprintf("create%sService", service.Name().UpperCamelCase())
}

func DocsDocComment(method pgs.Method) string {
	commentLines := deleteEmpty(strings.Split(method.SourceCodeInfo().LeadingComments(), "\n"))
	commentLines = append(GetAnnotatedComment(method), commentLines...)
	if len(commentLines) == 0 {
		return ""
	}
	return strings.Join(commentLines, "\n")
}

const DocsSampleTpl = `// BEGIN Code generated by protoc-gen-trinsic. DO NOT EDIT.
// target: {{ .TargetPath }}
{{ range .File.Services }}
## Service Creation

{{"{{"}} proto_sample_create_service("{{DocsServiceTab .}}") {{"}}"}}

{{ range .Methods }}
## {{ .Name }}

{{ DocsDocComment . }}

{{"{{"}} proto_sample_start("{{DocsMethodTab .}}") {{"}}"}}
{{"{{"}} proto_sample_code("{{DocsMethodTab .}}") {{"}}"}}
{{"{{"}} proto_method_tabs("{{DocsMethodTab .}}") {{"}}"}}
{{ end }}{{ end }}
// END Code generated by protoc-gen-trinsic. DO NOT EDIT.`
