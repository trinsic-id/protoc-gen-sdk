package lang_types

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

func SwiftMessageType(entity EntityWithParent) string {
	names := strings.Split(entity.FullyQualifiedName(), ".")
	// Remove empty leading name
	if names[0] == "" {
		names = names[1:]
	}
	// Go to upper case
	for i, name := range names {
		names[i] = pgs.Name(name).UpperCamelCase().String()
	}
	return fmt.Sprintf("%s", strings.Join(names, "_"))
}

func SwiftMethodReturnType(method pgs.Method) string {
	return SwiftMethodType(method.Output(), method.ServerStreaming())
}

func SwiftDocComment(method pgs.Method) string {
	commentLines := deleteEmpty(strings.Split(method.SourceCodeInfo().LeadingComments(), "\n"))
	commentLines = append(GetAnnotatedComment(method), commentLines...)
	if len(commentLines) == 0 {
		return ""
	}
	return fmt.Sprintf(" /// %s", strings.Join(commentLines, "\n///"))
}

func SwiftAsync(method pgs.Method) string {
	if method.ServerStreaming() {
		return ""
	}
	return "async"
}
func SwiftAwait(method pgs.Method) string {
	if method.ServerStreaming() {
		return ""
	}
	return "await"
}

func SwiftMethodParamType(message pgs.Method) string {
	return SwiftMethodType(message.Input(), false)
}

func SwiftMethodType(message pgs.Message, streaming bool) string {
	t := SwiftMessageType(message)
	if streaming {
		return fmt.Sprintf("TODO<%s>>", t)
	} else {
		return t
	}
}

func SwiftBuildMetadata(method pgs.Method) string {
	s := "nil"
	if !SdkAnonymous(method) {
		s = "request"
	}
	return fmt.Sprintf("try buildMetadata(%s)", s)
}

func SwiftAnnotations(method pgs.Method) string {
	isDep, msgDep := SdkDeprecated(method)
	isExp, msgExp := SdkExperimental(method)
	if isDep {
		return fmt.Sprintf("@available(*, deprecated, message: \"%s\")", msgDep)
	}
	if isExp {
		return fmt.Sprintf("@available(*, deprecated, message: \"%s\")", msgExp)
	}
	return ""
}

const SwiftServiceTpl = `// BEGIN Code generated by protoc-gen-trinsic. DO NOT EDIT.
// target: {{ .TargetPath }}
{{ range .File.Services }}{{ range .Methods }}{{ if SdkTemplateGenerate . }}
    {{ SwiftAnnotations . }}
    public func {{ .Name.LowerCamelCase }}(request: {{ SwiftMethodParamType . }}) throws -> {{ SwiftMethodReturnType . }} {
        return try client!.{{ .Name.UpperCamelCase }}(request, callOptions: {{ SwiftBuildMetadata . }})
            .response
            .wait()
    }{{ end }}{{ end }}{{ end }}
// END Code generated by protoc-gen-trinsic. DO NOT EDIT.`
