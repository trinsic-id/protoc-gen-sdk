package lang_types

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

func TypescriptMessageType(entity EntityWithParent) string {
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

func TypescriptMethodReturnType(method pgs.Method) string {
	return TypescriptMethodType(method.Output(), method.ServerStreaming())
}

func TypescriptMethodParamType(method pgs.Method) string {
	return fmt.Sprintf("proto.%s", MethodParamType(method))
}

func TypescriptDocComment(method pgs.Method) string {
	commentLines := deleteEmpty(strings.Split(method.SourceCodeInfo().LeadingComments(), "\n"))
	isDep, msgDep := SdkDeprecated(method)
	isExp, msgExp := SdkExperimental(method)
	if isDep {
		commentLines = append(commentLines, fmt.Sprintf("@deprecated %s", msgDep))
	}
	if isExp {
		commentLines = append(commentLines, fmt.Sprintf("@deprecated %s", msgExp))
	}
	if len(commentLines) == 0 {
		return ""
	}
	return fmt.Sprintf("/**%s */", strings.Join(commentLines, "\n*"))
}

func TypescriptAsync(method pgs.Method) string {
	if method.ServerStreaming() {
		return ""
	}
	return "async"
}
func TypescriptAwait(method pgs.Method) string {
	if method.ServerStreaming() {
		return ""
	}
	return "await"
}

func TypescriptMethodType(message pgs.Message, streaming bool) string {
	t := TypescriptMessageType(message)
	if streaming {
		return fmt.Sprintf("Promise<AsyncIterable<proto.%s>>", t)
	} else {
		return fmt.Sprintf("Promise<proto.%s>", t)
	}
}

func TypescriptBuildMetadata(method pgs.Method) string {
	s := ""
	if !SdkAnonymous(method) {
		s = fmt.Sprintf("proto.%s.encode(request).finish()", TypescriptMessageType(method.Input()))
	}
	return fmt.Sprintf("await this.buildMetadata(%s)", s)
}

func TypescriptMethodArguments(method pgs.Method) string {
	if SdkNoArguments(method) {
		return ""
	} else {
		return fmt.Sprintf("request: proto.%s", MethodParamType(method))
	}
}

func TypescriptDefaultRequestObject(method pgs.Method) string {
	if SdkNoArguments(method) {
		return fmt.Sprintf("let request = proto.%s.fromPartial({});", MethodParamType(method))
	}
	return ""
}

const TypescriptServiceTpl = `// BEGIN Code generated by protoc-gen-trinsic. DO NOT EDIT.
// target: {{ .TargetPath }}
{{ range .File.Services }}{{ range .Methods }}{{ if SdkTemplateGenerate . }}
  {{ TypescriptDocComment . }}
  public async {{ .Name.LowerCamelCase }}({{ TypescriptMethodArguments . }}): {{ TypescriptMethodReturnType . }} {
    {{ TypescriptDefaultRequestObject . }}
    return this.client.{{ .Name.LowerCamelCase }}(request, {
      metadata: {{ TypescriptBuildMetadata . }}
    });
  }{{ end }}{{ end }}{{ end }}
// END Code generated by protoc-gen-trinsic. DO NOT EDIT.`

const TypescriptDocTpl = `// BEGIN Doc Code generated by protoc-gen-trinsic. DO NOT EDIT.
// target: {{ .TargetPath }}
{{ range .File.Services }}
// {{ DocCreateServiceInjection . }}
// {{ DocMethodInjectionEnd }}
{{ range .Methods }}
// {{ DocMethodInjection . }}
// {{ DocMethodInjectionEnd }}
{{ end }}{{ end }}
// END Doc Code generated by protoc-gen-trinsic. DO NOT EDIT.`
