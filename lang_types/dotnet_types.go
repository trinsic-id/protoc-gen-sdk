package lang_types

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

func DotnetMessageType(entity EntityWithParent) string {
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

func DotnetMethodParamType(method pgs.Method) string {
	return DotnetMessageType(method.Input())
}

func DotnetMethodReturnType(method pgs.Method) string {
	return DotnetMethodType(method.Output(), method.ServerStreaming())
}

func DotnetDocComment(method pgs.Method) string {
	commentLines := deleteEmpty(strings.Split(method.SourceCodeInfo().LeadingComments(), "\n"))
	commentLines = append(GetAnnotatedComment(method), commentLines...)
	if len(commentLines) == 0 {
		return ""
	}
	const tmpl = `/// <summary>
    ///{COMMENT_TEXT}
    /// </summary>`

	return strings.ReplaceAll(tmpl, "{COMMENT_TEXT}", strings.Join(commentLines, "\n    /// "))
}

func DotnetMethodType(message pgs.Message, streaming bool) string {
	t := DotnetMessageType(message)
	if streaming {
		return fmt.Sprintf("IAsyncStreamReader<%s>", t)
	} else {
		return fmt.Sprintf("%s", t)
	}
}

func MethodIsStreaming(method pgs.Method) bool {
	return method.ServerStreaming()
}

func DotnetMethodArguments(method pgs.Method) string {
	if SdkNoArguments(method) {
		return ""
	} else {
		return fmt.Sprintf("%s request", DotnetMethodParamType(method))
	}
}

func DotnetDefaultRequestObject(method pgs.Method) string {
	if SdkNoArguments(method) {
		return fmt.Sprintf("var request = new %s();", DotnetMethodParamType(method))
	}
	return ""
}
func DotnetAnnotations(method pgs.Method) string {
	if SdkDeprecated(method) {
		return "[Obsolete(\"This method is deprecated\")]"
	} else if SdkExperimental(method) {
		return "[Obsolete(\"This method is experimental\")]"
	}
	// TODO - Support experimental
	return ""
}

const DotnetServiceTpl = `// BEGIN Code generated by protoc-gen-trinsic. DO NOT EDIT.
// target: {{ .TargetPath }}
{{ range .File.Services }}{{ range .Methods }}{{ if SdkTemplateGenerate . }}{{ if MethodIsStreaming . }}
	{{ DotnetDocComment . }}
    {{ DotnetAnnotations . }}
    public {{ DotnetMethodReturnType . }} {{ .Name.UpperCamelCase }}({{ DotnetMethodArguments .}}) {
        {{ DotnetDefaultRequestObject . }}
        return Client.{{ .Name.UpperCamelCase }}(request, {{ BuildMetadata . false }}).ResponseStream;
    }{{else}}
	{{ DotnetDocComment . }}
    {{ DotnetAnnotations . }}
    public {{ DotnetMethodReturnType . }} {{ .Name.UpperCamelCase }}({{ DotnetMethodArguments .}}) {
        {{ DotnetDefaultRequestObject . }}
        return Client.{{ .Name.UpperCamelCase }}(request, {{ BuildMetadata . false }});
    }
	
	{{ DotnetDocComment . }}
    {{ DotnetAnnotations . }}
    public async Task<{{ DotnetMethodReturnType . }}> {{ .Name.UpperCamelCase }}Async({{ DotnetMethodArguments .}}) {
        {{ DotnetDefaultRequestObject . }}
        return await Client.{{ .Name.UpperCamelCase }}Async(request, {{ BuildMetadata . true }});
    }
{{ end }}{{ end }}{{ end }}{{ end }}
// END Code generated by protoc-gen-trinsic. DO NOT EDIT.`
