package lang_types

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

func PythonMessageType(entity EntityWithParent) string {
	names := make([]string, 0)
	outer := entity
	ok := true
	for ok {
		name := outer.Name().String()
		names = append([]string{strings.Title(name)}, names...)
		outer, ok = outer.Parent().(pgs.Message)
	}
	return fmt.Sprintf("%s", strings.Join(names, "."))
}

func PythonMethodReturnType(method pgs.Method) string {
	t := PythonMessageType(method.Output())
	if method.ServerStreaming() {
		return fmt.Sprintf("AsyncIterator[%s]", t)
	}
	return t
}

func PythonDocComment(method pgs.Method) string {
	return fmt.Sprintf("\"\"\"\n%s\n\"\"\"", method.SourceCodeInfo().LeadingComments())
}

const PythonServiceTpl = `# BEGIN Code generated by protoc-gen-trinsic. DO NOT EDIT.
# target: {{ .TargetPath }}
{{ range .File.Services }}{{ range .Methods }}{{ if SdkTemplateGenerate . }}
    async def {{ .Name.LowerSnakeCase }}(self, *, request: {{ MethodParamType . }}) -> {{ PythonMethodReturnType . }}:
        {{ PythonDocComment . }}
        return await self.client.{{ .Name.LowerSnakeCase }}(request)
    {{ end }}{{ end }}
{{ end }}
# END Code generated by protoc-gen-trinsic. DO NOT EDIT.`
